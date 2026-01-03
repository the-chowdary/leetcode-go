/*
3512 - https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k
1. Calculate the total sum of the array.
2. Return the remainder of the total sum when divided by k.
*/
package main

func minOperations(nums []int, k int) int {
	total := 0
	for i := 0; i <= len(nums)-1; i++ {
		total += nums[i]
	}
	return total % k
}

func main() {

}
