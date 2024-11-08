/*
2974 : https://leetcode.com/problems/minimum-number-game/description/
*/
package main

import "slices"

func numberGame(nums []int) []int {
	slices.Sort(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}
	return nums
}

func main() {
	nums := []int{5, 4, 2, 3}
	numberGame(nums)
}
