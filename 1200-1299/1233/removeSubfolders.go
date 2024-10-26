package main

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem
func removeSubfolders(folder []string) []string {
    sort.Strings(folder)

    result := []string { folder[0] }

    for i := 1; i < len(folder); i++ {
        lastFolder := result[len(result) - 1] + "/"

        if !strings.HasPrefix(folder[i], lastFolder) {
            result = append(result, folder[i])
        }
    }
    return result
}