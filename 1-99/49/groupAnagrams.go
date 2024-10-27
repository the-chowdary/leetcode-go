/*
Implemented groupAnagrams
1. sorting each word and storing in hashMap[sortedWord] = word
2. Iterating the hasMap and storing anagram words in result array
*/
package main

import (
	"leetcode-go/utils"
)

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)

	for _, word := range strs {
		sortedWord := utils.SortString(word)
		hashMap[sortedWord] = append(hashMap[sortedWord], word)
	}

	result := make([][]string, 0, len(hashMap))
	for _, words := range hashMap {
		result = append(result, words)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}
