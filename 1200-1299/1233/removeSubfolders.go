/*
Implemented removeSubfolders
1. sort the folder array and assign result with folder[0]
2. Iterate from 1 to len(folder)
3. Get last folder and append '/'
4. Check if folder[i] has lastFolder as prefix if not it is a valid subfolder
*/
package main

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)

	result := []string{folder[0]}

	for i := 1; i < len(folder); i++ {
		lastFolder := result[len(result)-1] + "/"

		if !strings.HasPrefix(folder[i], lastFolder) {
			result = append(result, folder[i])
		}
	}
	return result
}

func main() {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	removeSubfolders(folder)
}
