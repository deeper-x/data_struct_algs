// base64.go
// Description: Implements Base64 encoding and decoding as specified in the RFC4648 standard.
// Time Complexity: O(n) - The encoding and decoding processes iterate through the input linearly.
// Space Complexity: O(n) - The output size is proportional to the input size.
// References:
//   - RFC4648 Base64 Encoding: https://datatracker.ietf.org/doc/html/rfc4648#section-4
//   - Wikipedia Base64 Overview: https://en.wikipedia.org/wiki/Base64
// See also: base64_test.go for test cases and verification.

package conversion

import (
	"strings" // Provides an efficient way to build strings without unnecessary memory allocations.
)

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Base64Encode converts a byte slice into a Base64-encoded string.
//
// The encoding follows the RFC4648 standard, where every 3 bytes of input
// are converted into 4 characters using a 6-bit mapping table. If the input
// length is not a multiple of 3, padding is added using '='.
//
// Example:
//
//	Input:  "hello"
//	Output: "aGVsbG8="
func Base64Encode(input []byte) string {
	var sb strings.Builder // Efficient string concatenation

	// Calculate padding required if the input length is not a multiple of 3.
	var padding string
	for i := len(input) % 3; i > 0 && i < 3; i++ {
		var zeroByte byte
		input = append(input, zeroByte) // Append zero bytes to align the length
		padding += "="                  // Add '=' padding to match the expected output length
	}

	// Process input 3 bytes at a time, converting them into 4 Base64 characters
	for i := 0; i < len(input); i += 3 {
		// Extract 3 bytes and split them into four 6-bit groups
		// Each byte contributes to multiple output characters
		group := [4]byte{
			input[i] >> 2,                        // First 6 bits
			(input[i]<<4)&0x3F + input[i+1]>>4,   // Next 6 bits (spanning two bytes)
			(input[i+1]<<2)&0x3F + input[i+2]>>6, // Next 6 bits (spanning two bytes)
			input[i+2] & 0x3F,                    // Last 6 bits
		}

		// Convert each 6-bit group to a Base64 character
		for _, b := range group {
			sb.WriteByte(Alphabet[b])
		}
	}

	encoded := sb.String()

	// Apply '=' padding if necessary
	encoded = encoded[:len(encoded)-len(padding)] + padding[:]

	return encoded
}

// Base64Decode converts a Base64-encoded string back into a byte slice.
//
// This function processes 4-character chunks from the input string, converting
// them back into 3 original bytes. Padding ('=') characters at the end of the input
// are ignored to restore the correct output length.
//
// Example:
//
//	Input:  "aGVsbG8="
//	Output: "hello"
func Base64Decode(input string) []byte {
	padding := strings.Count(input, "=") // Count padding characters, which affect output size
	var decoded []byte

	// Process input in chunks of 4 Base64 characters at a time
	for i := 0; i < len(input); i += 4 {
		// Convert each Base64 character back to its corresponding 6-bit value
		byteInput := [4]byte{
			byte(strings.IndexByte(Alphabet, input[i])),
			byte(strings.IndexByte(Alphabet, input[i+1])),
			byte(strings.IndexByte(Alphabet, input[i+2])),
			byte(strings.IndexByte(Alphabet, input[i+3])),
		}

		// Reassemble original bytes from 6-bit groups
		group := [3]byte{
			byteInput[0]<<2 + byteInput[1]>>4, // First byte
			byteInput[1]<<4 + byteInput[2]>>2, // Second byte
			byteInput[2]<<6 + byteInput[3],    // Third byte
		}

		decoded = append(decoded, group[:]...)
	}

	// Remove extra bytes that were added due to padding
	return decoded[:len(decoded)-padding]
}
