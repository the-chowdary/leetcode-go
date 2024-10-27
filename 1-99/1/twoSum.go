/*
Implments a solution for twoSum
1. Iterate nums and get found -> (target - value)
2. If it already exists in hashMap return [found, key]
3. Else store the hashmap[value] = key
4. If nothing found return [-1, -1]
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
