/*
2825 : https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/description/?envType=daily-question&envId=2024-12-04
1. Two pointers
2. If str1[0] can match str2[0] then it is a match else there is no match later.
*/
package main

func canMakeSubsequence(str1, str2 string) bool {
	if len(str1) < len(str2) {
		return false
	}
	j := 0
	for _, i := range str1 {
		c := byte(i) + 1
		if i == 'z' {
			c = 'a'
		}
		if byte(i) == str2[j] || c == str2[j] {
			j++
			if j == len(str2) {
				return true
			}
		}
	}
	return false
}

func main() {

}
