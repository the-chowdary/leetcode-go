/*
1574 : https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted
1. If already sorted return 0
2. Increment left and right pointers to correct places (arr[l] <= arr[l + 1] l++) and (arr[r] >= arr[r - 1] r--)
*/
package main

import (
	"sort"
)

func findLengthOfShortestSubarray(arr []int) (result int) {
	if sort.IntsAreSorted(arr) {
		return
	}

	N := len(arr)
	l, r := 0, N-1

	for ; arr[l] <= arr[l+1]; l++ {

	}

	for ; arr[r] >= arr[r-1]; r-- {

	}

	result = min(N-1-l, r)
	for ; r < N; r++ {
		result = min(result, r-sort.SearchInts(arr[:l+1], arr[r]+1))
	}
	return
}

func main() {
	
}
