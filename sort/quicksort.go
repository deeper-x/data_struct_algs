// quicksort.go
// description: Implementation of in-place quicksort algorithm
// details:
// A simple in-place quicksort algorithm implementation. [Wikipedia](https://en.wikipedia.org/wiki/Quicksort)
// worst time complexity: O(n^2)
// average time complexity: O(n log n)
// space complexity: O(log n)
// see sort_test.go for a test implementation, test function TestQuickSort.

package sort

import (
	"math/rand"

	"github.com/deeper-x/data_struct_algs/constraints"
)

func Quicksort[T constraints.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	Quicksort(arr[:left])
	Quicksort(arr[left+1:])

	return arr
}
