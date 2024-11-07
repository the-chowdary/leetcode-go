/*
1. Compare each and every element since (AND of 0 gives 0) consider only 1
2. Traverse elements and count number of 1 bit and return its sum
3. 2^23 < 10^7 < 2^24 so looping till 24
*/
package main

import "leetcode-go/utils"

func largestCombination(candidates []int) (result int) {
	for i := 0; i < 24; i++ {
		s := 0
		for _, v := range candidates {
			s += v >> i & 1
		}
		result = utils.MaxNum(result, s)
	}
	return
}

func main() {
	candidates := []int{16, 17, 71, 62, 12, 24, 14}
	largestCombination(candidates)
}
