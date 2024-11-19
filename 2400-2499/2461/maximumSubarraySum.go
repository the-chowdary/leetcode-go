/*
2461 : https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/description
Sliding window
*/
package main

func maximumSubarraySum(nums []int, k int) int64 {
	sum, result := 0, 0
	count := map[int]int{}

	for _, x := range nums[:k-1] {
		count[x]++
		sum += x
	}

	for i := k - 1; i < len(nums); i++ {
		count[nums[i]]++
		sum += nums[i]

		if len(count) == k && sum > result {
			result = sum
		}

		x := nums[i+1-k]
		count[x]--

		if count[x] == 0 {
			delete(count, x)
		}
		sum -= x
	}
	return int64(result)
}

func main() {

}
