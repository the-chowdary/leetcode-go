/*
Implemented plusOne : https://leetcode.com/problems/plus-one/description/
1. Iterate from the last digit to the first digit
2. If digit < 9, increment and return
3. If digit == 9, set to 0 and continue
4. If all digits are 9, prepend 1 to the result
*/
package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	digits := []int{1, 2, 3}
	plusOne(digits)
}
