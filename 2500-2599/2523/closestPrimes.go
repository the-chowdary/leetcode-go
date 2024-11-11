/*
2523 : https://leetcode.com/problems/closest-prime-numbers-in-range/description/
1. Identify all primes with sieve of eratosthenes
2. Find left and right minimum

Similar Problems:
204 : https://leetcode.com/problems/count-primes/
2601 : https://leetcode.com/problems/prime-subtraction-operation
*/
package main

import "sort"

const mx = 1e6 + 1

var primes = make([]int, 0, 78500)

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
	primes = append(primes, mx, mx) // make sure not to cross the boundary
}

func closestPrimes(left, right int) []int {
	p, q := -1, -1
	for i := sort.SearchInts(primes, left); primes[i+1] <= right; i++ {
		if p < 0 || primes[i+1]-primes[i] < q-p {
			p, q = primes[i], primes[i+1]
		}
	}
	return []int{p, q}
}

func main() {
	left, right := 10, 19
	closestPrimes(left, right)
}
