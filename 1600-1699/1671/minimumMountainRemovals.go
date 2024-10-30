/*
1671 : https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/description/
1. Get prefix and suffix of nums
2. Iterate over range nums and return min(len(nums) - prefix[i] - suffix[i])
*/
package main

import "leetcode-go/utils"

func minimumMountainRemovals(nums []int) int {
	N := len(nums)
	// Mountain should have atleast 3 elements
	if N < 3 {
		return 0
	}

	prefix := make([]int, N)
	for i := 0; i < N; i++ {
		prefix[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				prefix[i] = utils.MaxNum(prefix[i], prefix[j]+1)
			}
		}
	}

	suffix := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		suffix[i] = 1
		for j := N - 1; j > i; j-- {
			if nums[j] < nums[i] {
				suffix[i] = utils.MaxNum(suffix[i], suffix[j]+1)
			}
		}
	}

	result := N
	for i := 1; i < N-1; i++ {
		if prefix[i] > 0 && suffix[i] > 1 {
			result = utils.MinNum(result, N-prefix[i]-suffix[i]+1)
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 1}
	minimumMountainRemovals(nums)
}
