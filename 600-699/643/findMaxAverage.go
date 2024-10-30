/*
643 : https://leetcode.com/problems/maximum-average-subarray-i/description/
1. Fixed length sliding window

Similar problems :
1456 : https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
*/
package main

func findMaxAverage(nums []int, k int) (result float64) {
	sum := 0

	for i, n := range nums {

		// 1. Enter the window
		sum += n

		if i < k-1 {
			continue
		}

		if i == k-1 {
			result = float64(sum) / float64(k)
		} else {
			currentAvg := float64(sum) / float64(k)
			result = max(result, currentAvg)
		}

		// Leave the window
		sum -= nums[i-k+1]
	}
	return
}

func main() {

}
