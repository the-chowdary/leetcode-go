/*
796 :https://leetcode.com/problems/rotate-string/description
1. Compare length if not same return false
2. goal should be there in s + s then return true else false
*/
package main

import "strings"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}

func main() {
	s := "abcde"
	goal := "cdeab"
	rotateString(s, goal)
}
