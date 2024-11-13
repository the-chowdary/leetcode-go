/*
3345 : https://leetcode.com/problems/smallest-divisible-digit-product-i/
*/
package main

func smallestNumber(n, t int) int {
	for i := n; ; i++ {
		prod := 1
		for x := i; x > 0; x /= 10 {
			prod *= x % 10
		}
		if prod%t == 0 {
			return i
		}
	}
}

func main() {

}
