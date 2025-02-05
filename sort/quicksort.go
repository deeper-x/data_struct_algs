package sort

import (
	"math/rand"

	"github.com/deeper-x/data_struct_algs/constraints"
)

// Quicksort sorts a slice using the quicksort algorithm.
//
// The algorithm follows these steps:
//  1. Selects a pivot element randomly.
//  2. Partitions the array into elements smaller and greater than the pivot.
//  3. Recursively applies the same process to the subarrays.
//
// Parameters:
//   - arr []T: A slice of ordered elements to be sorted.
//
// Returns:
//   - []T: The sorted slice.
func Quicksort[T constraints.Ordered](arr []T) []T {
	// Base case: If the slice has 0 or 1 element, it is already sorted.
	if len(arr) < 2 {
		return arr
	}

	// Initialize pointers for partitioning.
	left, right := 0, len(arr)-1

	// Select a random pivot index within the slice bounds.
	pivot := rand.Int() % len(arr)

	// Swap pivot element with the last element to simplify partitioning.
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Partitioning: Move elements smaller than the pivot to the left side.
	for i := range arr {
		if arr[i] < arr[right] { // Compare with pivot element.
			arr[left], arr[i] = arr[i], arr[left] // Swap smaller element to the left.
			left++
		}
	}

	// Place the pivot element in its correct position.
	arr[left], arr[right] = arr[right], arr[left]

	// Recursively apply Quicksort to the left and right subarrays.
	Quicksort(arr[:left])   // Sort elements before pivot.
	Quicksort(arr[left+1:]) // Sort elements after pivot.

	// Return the sorted slice.
	return arr
}
