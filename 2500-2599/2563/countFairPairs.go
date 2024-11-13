/*
2563 : https://leetcode.com/problems/count-the-number-of-fair-pairs/description/
1. Sort the nums
2. lower - nums[j] <= nums[i] <= upper - nums[j]
*/
package main

import "sort"

func countFairPairs1(nums []int, lower, upper int) (result int64) {
	sort.Ints(nums)
	for j, x := range nums {
		r := sort.SearchInts(nums[:j], upper-x+1)
		l := sort.SearchInts(nums[:j], lower-x)
		result += int64(r - l)
	}
	return result
}

func countFairPairs2(nums []int, lower, upper int) (result int64) {
	sort.Ints(nums)
	left, right := len(nums), len(nums)

	for j, x := range nums {
		for right > 0 && nums[right-1] > upper-x {
			right--
		}
		for left > 0 && nums[left-1] >= lower-x {
			left--
		}
		result += int64(min(right, j) - min(left, j))
	}
	return
}

func main() {

}
