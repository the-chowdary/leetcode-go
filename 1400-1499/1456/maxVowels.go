/*
1456 : https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
1. Fixed length sliding window
2. Enter the window
3. Update the result mostly will be max or min
4. Leave the window

Similar problems:
1. 643 : https://leetcode.com/problems/maximum-average-subarray-i/description/
2. 1343 : https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/description/
*/
package main

import "leetcode-go/utils"

func maxVowels(s string, k int) (result int) {
	vowels := 0
	for i, c := range s {
		// 1. Enter the window
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels++
		}

		if i < k-1 {
			continue
		}

		// 2. Update the result
		result = utils.MaxNum(result, vowels)

		// 3. Leave the window
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			vowels--
		}
	}
	return
}

func main() {

}
