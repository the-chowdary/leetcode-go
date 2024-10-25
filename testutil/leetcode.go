package main

import (
	"fmt"
	"reflect"
	"strconv"
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

func parseRawArg(tp reflect.Type, rawData string) (v reflect.Value, err error) {
	rawData = strings.TrimSpace(rawData)
	invalidError := fmt.Errorf("invalid test data %s", rawData)

	switch tp.Kind() {
	case reflect.String:
		if len(rawData) <= 1 || rawData[0] != '"' && rawData[0] != '\'' || rawData[len(rawData) - 1] != rawData[0] {
			return reflect.Value{}, invalidError
		}
		// Remove " or ' from start and end
		v = reflect.ValueOf(rawData[1 : len(rawData) - 1])
	
	case reflect.Uint8: // byte
		// rawData like "a" or 'a'
		if len(rawData) != 3 || rawData[0] != '"' && rawData[0] != '\'' || rawData[2] != rawData[0] {
			return reflect.Value{}, invalidError
		}
		v = reflect.ValueOf(rawData[1])

	case reflect.Int: 
		i, err := strconv.Atoi(rawData)
		if err != nil {
			return reflect.Value{}, invalidError
		}
		v = reflect.ValueOf(i)

	case reflect.Uint: 
		i, err := strconv.Atoi(rawData)
		if err != nil {
			return reflect.Value{}, invalidError
		}
		v = reflect.ValueOf(uint(i))

	case reflect.Int64:
		i, err := strconv.ParseInt(rawData, 10, 64)
		if err != nil {
			return reflect.Value{}, invalidError
		}
		v = reflect.ValueOf(i)
	
	case reflect.Uint64:
		i, err := strconv.ParseUint(rawData, 10, 64)
		if err != nil {
			return reflect.Value{}, invalidError
		}
		v = reflect.ValueOf(i)
	
	case reflect.Float64:
		f, err := strconv.ParseFloat(rawData, 64)
		if err != nil {
			return reflect.Value{}, invalidError
		}
		v = reflect.ValueOf(f)
	
	case reflect.Bool:
		b, err := strconv.ParseBool(rawData)
		if err != nil {
			return reflect.Value{}, invalidError
		}
		v = reflect.ValueOf(b)
	
	case reflect.Slice:
		splits, err := parseRawArray(rawData)
		if err != nil {
			return reflect.Value{}, err
		}
		v = reflect.New(tp).Elem()
		for _, s := range splits {
			_v, err := parseRawArg(tp.Elem(), s)
			if err != nil {
				return reflect.Value{}, err
			}
			v = reflect.Append(v, _v)
		}
	
	case reflect.Pointer: //*TreeNode, *ListNode, *Point, *Interval
		switch tmpName := tp.Elem().Name(); tmpName {
		case "TreeNode": 
			
		}
	}

}