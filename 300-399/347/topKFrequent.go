package main

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	var result []int
	countMap := make(map[int]int)
	countSliceLength := len(nums) + 1
	countSlice := make([][]int, countSliceLength)

	for _, num := range nums {
		countMap[num]++
	}

	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	for i := countSliceLength - 1; len(result) != k; i-- {
		result = append(result, countSlice[i]...)
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	topKFrequent(nums, k)
}
