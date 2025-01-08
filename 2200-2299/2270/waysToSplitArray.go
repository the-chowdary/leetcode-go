/*
2270 : https://leetcode.com/problems/number-of-ways-to-split-array/description/
*/
package main

func waysToSplitArray(nums []int) (result int) {
	total := 0
	for _, x := range nums {
		total += x
	}
	s := 0
	for _, x := range nums[:len(nums)-1] {
		s += x
		if s*2 >= total {
			result++
		}
	}
	return
}

func main() {

}
