package utils

import (
	"sort"
	"strings"
)

// sort the given string
func SortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}