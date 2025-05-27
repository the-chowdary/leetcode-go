//go:build go1.23

/*
1338 : https://leetcode.com/problems/reduce-array-size-to-the-half/description/
1. Count the numbers
2. Sort the counted values
3. Iterate over the values individually and if sum >= len(arr) / 2 return that index.
*/
package main

import (
	"maps"
	"slices"
)

func minSetSize(arr []int) int {
	freq := map[int]int{}
	for _, x := range arr {
		freq[x]++
	}

	s := 0
	cnt := slices.SortedFunc(maps.Values(freq), func(a, b int) int { return b - a })
	for i, c := range cnt {
		s += c
		if s >= len(arr)/2 {
			return i + 1
		}
	}
	panic("impossible")
}

func main() {

}
