/*
Implemented containsDuplicates
1. hashmap to count all the elementes
2. Iterate the hashmap if num already exist return true
3. If no duplicate found return false
*/
package main

// https://leetcode.com/problems/contains-duplicate/
func containsDuplicates(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	count := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := count[num]; ok {
			return true
		}
		count[num] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	containsDuplicates(nums)
}
