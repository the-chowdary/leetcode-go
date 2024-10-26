package main

// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/description/
var result int
var pre map[string]struct{}

func maxUniqueSplit(s string) int {
    result = 0
    pre = make(map[string]struct{})
    N := len(s)
    dfs("", 0, s, N)
    return result
}

func dfs(st string, i int, s string, N int) {
    if i == N {
        if st != "" {
            if _, exists := pre[st]; !exists {
                pre[st] = struct{}{}
                if len(pre) > result {
                    result = len(pre)
                }
                delete(pre, st)
            }
            return
        }
    }
    if st != "" {
        if _, exists := pre[st]; !exists {
            pre[st] = struct{}{}
            dfs(string(s[i]), i + 1, s, N)
            delete(pre, st)
        }
    }
    dfs(st + s[i : i + 1], i + 1, s, N)
}