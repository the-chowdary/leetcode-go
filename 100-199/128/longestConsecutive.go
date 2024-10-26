package main

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	longest := 0

	for _, num := range nums {
		set[num] = true
	}

	for _, num := range nums {
		if !set[num-1] {
			length := 0
			for set[num+length] {
				length++
			}
			longest = max(longest, length)
		}
	}
	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	longestConsecutive(nums)
}
