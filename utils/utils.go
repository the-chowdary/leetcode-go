/*
Package utils provides utility functions.
*/
package utils

import (
	"sort"
	"strings"
)

// SortString sorts the given string
func SortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// Max returns max values in the int array
func Max(numbers []int) (result int) {
	for _, num := range numbers {
		if num > result {
			result = num
		}
	}
	return
}

// Sum returns total sum of the int array
func Sum(numbers []int) (result int) {
	for _, num := range numbers {
		result += num
	}
	return
}

// MaxNum returns max(x, y)
func MaxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MinNum return min(x, y)
func MinNum(x, y int) int {
	if x < y {
		return x
	}
	return y
}
