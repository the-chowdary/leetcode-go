/*
1390 - https://leetcode.com/problems/four-divisors/?envType=daily-question
1. Calculate the number of divisors and the sum of divisors for all numbers up to 100000 using a modified sieve approach.
2. For each number in the input array, check if it has exactly four divisors.
3. If it does, add the sum of its divisors to the result.
4. Return the total sum of divisors for all numbers with exactly four divisors.
*/
package main

const mx = 100_001

var divisorNum, divisorSum [mx]int

func init() {
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i {
			divisorNum[j]++
			divisorSum[j] += i
		}
	}
}

func sumFourDivisors(nums []int) (result int) {
	for _, x := range nums {
		if divisorNum[x] == 4 {
			result += divisorSum[x]
		}
	}
	return
}

func main() {

}
