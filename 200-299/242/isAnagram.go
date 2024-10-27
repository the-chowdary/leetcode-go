/*
Implemented isAnagram
1. Convert string to byte array
2. Sort the byte array for both s and t
3. return s == t after sorting
*/
package main

import "sort"

func isAnagram(s string, t string) bool {
	sBytes, tBytes := []byte(s), []byte(t)
	sort.Slice(sBytes, func(x, y int) bool { return sBytes[x] < sBytes[y] })
	sort.Slice(tBytes, func(x, y int) bool { return tBytes[x] < tBytes[y] })
	return string(sBytes) == string(tBytes)
}

func main() {
	s, t := "anagram", "nagaram"
	isAnagram(s, t)
}
