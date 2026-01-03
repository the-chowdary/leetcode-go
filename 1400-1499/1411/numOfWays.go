/*
1411 : https://leetcode.com/problems/number-of-ways-to-paint-n-3-grid/description
*/
package main

func numOfWays(n int) int {
	const mod = 1_000_000_007
	f := make([]int, n+1)
	f[0] = 3
	f[1] = 12

	for i := 2; i <= n; i++ {
		f[i] = (f[i-1]*5 - f[i-2]*2) % mod
	}
	return (f[n] + mod) % mod // non -ve result
}

func main() {

}
