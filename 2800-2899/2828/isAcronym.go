/*
2828 : https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/description/
1. Iterate and compare
2. If len(words) != len(s) then false
*/
package main

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i, w := range words {
		if w[0] != s[i] {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"alice", "bob", "charlie"}
	s := "abc"
	isAcronym(words, s)
}
