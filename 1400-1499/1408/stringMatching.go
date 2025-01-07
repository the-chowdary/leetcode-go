/*
* 1408 : https://leetcode.com/problems/string-matching-in-an-array/
 */
package main

import "strings"

func stringMatching(words []string) (result []string) {
	for i, w := range words {
		for j, w2 := range words {
			if i != j && strings.Contains(w2, w) {
				result = append(result, w)
				break
			}
		}
	}
	return
}

func main() {

}
