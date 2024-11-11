/*
2601 : https://leetcode.com/problems/prime-subtraction-operation

1. Identify all primes using sieve of eratosthenes
2. Identify the small prime and subtract the num by traversing.

Similar problems :
2523 : https://leetcode.com/problems/closest-prime-numbers-in-range/description/
*/
package main

import "sort"

var p = []int{0}

func init() {
	const mx = 1000
	np := [mx]bool{} // non primes
	for i := 2; i < mx; i++ {
		if !np[i] {
			p = append(p, i) // save prime number
			for j := i * i; j < mx; j += i {
				np[j] = true // mark composite numbers
			}
		}
	}
}

func primeSubOperation(nums []int) bool {
	pre := 0
	for _, x := range nums {
		if x <= pre {
			return false
		}
		pre = x - p[sort.SearchInts(p, x-pre)-1] // Identify the x - pre
	}
	return true
}

func main() {
	nums := []int{4, 9, 6, 10}
	primeSubOperation(nums)
}
