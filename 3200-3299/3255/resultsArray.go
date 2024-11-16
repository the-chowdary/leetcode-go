/*
3255 : https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/description
1. Initialize result with size len(nums) - k + 1
2. Count number of consecutive
3. if count >= k then update result

Similar Problems:
3254 : https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/description
*/
package main

func resultsArray(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)
	for i := range result {
		result[i] = -1
	}
	count := 0
	for i, x := range nums {
		if i == 0 || nums[i] == nums[i-1]+1 { // count consecutive
			count++
		} else {
			count = 1
		}
		if count >= k { // if count >= k then update result
			result[i-k+1] = x
		}
	}
	return result
}

func main() {

}
