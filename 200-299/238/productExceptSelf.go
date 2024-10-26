package main

func productExceptSelf(nums []int) []int {
	N := len(nums)
	result := make([]int, len(nums))

	for i := range nums {
		result[i] = 1
	}

	prefix, suffix := 1, 1

	for i := range nums {
		result[i] *= prefix
		prefix *= nums[i]
		result[N-1-i] *= suffix
		suffix *= nums[N-1-i]
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	productExceptSelf(nums)
}
