/*
Implemented repeatedNTimes : https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
1. Create a map to store seen elements
2. Iterate through the nums array
3. If the element is already in the map, return it
4. Otherwise, add the element to the map
5. If no element is found, return -1
*/
package main

func repeatedNTimes(nums []int) int {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return -1
}

func main() {

}
