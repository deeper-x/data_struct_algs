package sort

import (
	"math/rand"

	"github.com/deeper-x/data_struct_algs/constraints"
)

// / Quicksort sorts an array using the quicksort algorithm.
// / It selects a pivot, partitions the array, and recursively sorts the subarrays.
// /
// / Parameters:
// /   - arr []T: A slice of ordered elements to be sorted.
// /
// / Returns:
// /   - []T: The sorted slice.
func Quicksort[T constraints.Ordered](arr []T) []T {
	/// Base case: If the array has 0 or 1 element, it's already sorted.
	if len(arr) < 2 {
		return arr
	}

	/// Initialize pointers for partitioning
	left, right := 0, len(arr)-1

	/// Select a random pivot index
	pivot := rand.Int() % len(arr)

	/// Move pivot element to the end of the array
	arr[pivot], arr[right] = arr[right], arr[pivot]

	/// Partitioning: Move elements smaller than pivot to the left side
	for i := range arr {
		if arr[i] < arr[right] { /// Compare with pivot element
			arr[left], arr[i] = arr[i], arr[left] /// Swap smaller element to the left
			left++
		}
	}

	/// Place the pivot element in its correct position
	arr[left], arr[right] = arr[right], arr[left]

	/// Recursively apply Quicksort to the left and right subarrays
	Quicksort(arr[:left])   /// Sort elements before pivot
	Quicksort(arr[left+1:]) /// Sort elements after pivot

	/// Return the sorted array
	return arr
}
