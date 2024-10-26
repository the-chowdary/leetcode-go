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

func isLetterOrDigit(s rune) bool {
	return unicode.IsDigit(s) || unicode.IsLetter(s)
}

func main() {
	isPalindrome("A man, a plan, a canal: Panama")
}
