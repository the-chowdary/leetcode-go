/*
2490 : https://leetcode.com/problems/circular-sentence/description
1. First and last letter should be same if not return false
2. If space encountered check letter before space and after space they should be same if not return false
*/
package main

func isCircularSentence(s string) bool {
	// First letter and last letter of sentence should match
	if s[0] != s[len(s)-1] {
		return false
	}
	// If space, and it's previous letter and next letter should be same
	for i, c := range s {
		if c == ' ' && s[i-1] != s[i+1] {
			return false
		}
	}
	return true
}

func main() {
	s := "leetcode exercises sound delightful"
	isCircularSentence(s)
}
