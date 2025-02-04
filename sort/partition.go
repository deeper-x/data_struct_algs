package sort

import "github.com/deeper-x/data_struct_algs/constraints"

// / Partition rearranges the elements of arr[low:high] around a pivot.
// / It ensures that elements less than or equal to the pivot are on the left,
// / and elements greater than the pivot are on the right. The pivot is chosen
// / as the last element in the range.
// /
// / Parameters:
// /   - arr []T: A slice of ordered elements.
// /   - low, high int: The indices defining the subarray to partition.
// /
// / Returns:
// /   - int: The final index of the pivot after partitioning.
func Partition[T constraints.Ordered](arr []T, low, high int) int {
	/// Index to track the correct position for smaller elements
	index := low - 1

	/// Choosing the pivot as the last element
	pivotElement := arr[high]

	/// Iterate through the array, moving smaller elements to the left
	for i := low; i < high; i++ {
		if arr[i] <= pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index] /// Swap to move smaller element
		}
	}

	/// Swap pivot into its correct position
	arr[index+1], arr[high] = arr[high], arr[index+1]

	/// Return the final pivot index
	return index + 1
}
