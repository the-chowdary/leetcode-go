/*
204 : https://leetcode.com/problems/count-primes/
1. Identify all the primes with sieve of eratosthenes
2. Return count

Similar Problems:
2523 : https://leetcode.com/problems/closest-prime-numbers-in-range/description/
2601 : https://leetcode.com/problems/prime-subtraction-operation
*/
package main

import "sort"

const mx = 5*1e6 + 1

var primes = make([]int, 0, mx)

func init() {
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
	primes = append(primes, mx, mx) // don't cross the boundary
}

func countPrimes(n int) int {
	return sort.SearchInts(primes, n)
}

func main() {
	n := 4
	countPrimes(n)
}
