/*
3264 : https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/description
*/
package main

func getFinalState(nums []int, k, multiplier int) []int {
	N := len(nums)
	for k > 0 {
		minIndex := 0
		for i := range N {
			if nums[i] < nums[minIndex] {
				minIndex = i
			}
		}
		nums[minIndex] *= multiplier
		k--
	}
	return nums
}

func main() {

}
