/*
2185 : https://leetcode.com/problems/counting-words-with-a-given-prefix/description
*/
package main

import "strings"

func prefixCount(words []string, prefix string) (result int) {
	for _, word := range words {
		if strings.HasPrefix(word, prefix) {
			result++
		}
	}
	return
}

func main() {

}
