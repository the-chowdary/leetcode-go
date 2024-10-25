package main 

// https://leetcode.com/problems/contains-duplicate/
func containsDuplicates(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	count := make(map[int] struct{})

	for _, num := range nums {
		if _, ok := count[num]; ok {
			return true
		}
		count[num] = struct{}{}
	}
	return false
}