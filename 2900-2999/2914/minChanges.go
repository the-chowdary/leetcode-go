/*
2914 : https://leetcode.com/problems/minimum-number-of-changes-to-make-binary-string-beautiful/description
1. Compare even places see if there are same or not and return
*/
package main

func minChanges(s string) (result int) {
	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			result++
		}
	}
	return
}

func main() {

}
