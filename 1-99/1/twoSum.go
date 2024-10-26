/*
Implments a solution for twoSum
*/
package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for key, value := range nums {
		found := target - value

		if found, ok := hashMap[found]; ok {
			return []int{found, key}
		}
		hashMap[value] = key
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
}
