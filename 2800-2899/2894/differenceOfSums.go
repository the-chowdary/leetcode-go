/*
2894 : https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/description
1. Iterate and use mod and return num1 - num2
*/
package main

func differenceOfSums(n int, m int) int {
	num1, num2 := 0, 0
	for i := 1; i <= n; i++ {
		if i%m != 0 {
			num1 += i
		} else {
			num2 += i
		}
	}
	return num1 - num2
}

func main() {

}
