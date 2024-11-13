/*
1343 : https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/description/
1. Calculate prefix sum
2. Fixed length sliding window

Similar Problems :
643 : https://leetcode.com/problems/maximum-average-subarray-i/description/
1456 : https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
*/
package main

func numOfSubarrays(arr []int, k int, threshold int) (result int) {
	N := len(arr)
	sum := make([]int, N+1)

	// Prefix sum
	for i, v := range arr {
		sum[i+1] = sum[i] + v
	}

	// Sliding window
	for r := k; r <= N; r++ {
		if sum[r]-sum[r-k] >= threshold*k {
			result++
		}
	}
	return
}

func main() {

}
