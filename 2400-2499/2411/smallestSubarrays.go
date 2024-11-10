/*
2411 : https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or/description/
1. if nums[j] | x != nums[j] then nums[j] can become larger.
2. Update result after increasing.
*/
package main

func smallestSubarrays(nums []int) []int {
	result := make([]int, len(nums))
	for i, x := range nums {
		result[i] = 1
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x
			result[j] = i - j + 1
		}
	}
	return result
}

func main() {

}
