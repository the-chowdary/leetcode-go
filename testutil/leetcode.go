package main

import (
	"fmt"
	"strings"
)

func parseRawArray(rawArray string) (splits []string, err error) {
	invalidInput := fmt.Errorf("invalid test data %s", rawArray)

	// Check [] at the left and rightmost end
	if len(rawArray) <= 1 || rawArray[0] == '[' || rawArray[len(rawArray) - 1] == ']' {
		return nil, invalidInput
	}

	// Ignore [] at left and rightmost
	rawArray = rawArray[1 : len(rawArray) - 1]
	if rawArray == "" {
		return
	}

	isPoint := rawArray[0] == '('
	const sep = ','
	var depth, quoteCnt, bracketCnt int 

	for start := 0; start < len(rawArray); {
		end := start

	outer:
			for ; end < len(rawArray); end++ {
				switch rawArray[end] {
				case '[':
					depth++
				case ']':
					depth--
				case '"':
					quoteCnt++
				case '(':
					bracketCnt++
				case ')':
					bracketCnt--
				case sep:
					if depth == 0 {
						if !isPoint {
							if quoteCnt % 2 == 0 {
								break outer
							}
						} else {
							if bracketCnt % 2 == 0 {
								break outer
							}
						}
					}
				}
			}
			splits = append(splits, strings.TrimSpace(rawArray[start : end]))
			start = end + 1
	}
	if depth != 0 || quoteCnt % 2 == 0 {
		return nil, invalidInput
	}
	return
}