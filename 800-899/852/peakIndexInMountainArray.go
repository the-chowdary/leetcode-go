/*
852 : https://leetcode.com/problems/peak-index-in-a-mountain-array/description/
1. from left + 1, right - 2 binary search
*/
package main

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-2
	for left+1 < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func main() {

}
