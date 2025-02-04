package matrix

import "github.com/deeper-x/data_struct_algs/constraints"

// IsValid checks if the input matrix has consistent row lengths.
func IsValid[T constraints.Integer](elements [][]T) bool {
	if len(elements) == 0 {
		return true
	}
	columns := len(elements[0])
	for _, row := range elements {
		if len(row) != columns {
			return false
		}
	}
	return true
}
