/*
Implemented isPalindrome
1. Two pointers(i, j) one at start and one at end
2. left if !isLetterOrDigit i++
3. right if !isLetterOrDigit j--
4. if left != right then not palindrome return false
5. regardless increment i++, j--
*/
package main

import "unicode"

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	arr := []rune(s)

	for i < j {
		left := unicode.ToLower(arr[i])
		right := unicode.ToLower(arr[j])

		if !isLetterOrDigit(left) {
			i++
			continue
		}

		if !isLetterOrDigit(right) {
			j--
			continue
		}

		if left != right {
			return false
		}
		i++
		j--
	}
	return true
}

// returns boolean whether character passed is letter or digit
func isLetterOrDigit(s rune) bool {
	return unicode.IsDigit(s) || unicode.IsLetter(s)
}

func main() {
	isPalindrome("A man, a plan, a canal: Panama")
}
