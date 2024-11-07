/*
1. Have 2 loops inner loop to go till end of array or if bitCount == ones
2. Exchange the elements
*/
package main

import (
	"math/bits"
	"slices"
)

// Direct sorting
func canSortArray(nums []int) bool {
	for i, n := 0, len(nums); i < n; {
		start := i
		ones := bits.OnesCount(uint(nums[i]))
		i++
		for i < n && bits.OnesCount(uint(nums[i])) == ones {
			i++
		}
		slices.Sort(nums[start:i])
	}
	return slices.IsSorted(nums)
}

func main() {
	nums := []int{8, 4, 2, 30, 15}
	canSortArray(nums)
}
