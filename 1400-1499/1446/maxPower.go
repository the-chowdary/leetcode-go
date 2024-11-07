/*
1446 : https://leetcode.com/problems/consecutive-characters/description/
1. Two loops one for setup and one to calculate diving the work
Similar problems :
3011 :https://leetcode.com/problems/find-if-array-can-be-sorted/
*/
package main

import "leetcode-go/utils"

func maxPower(s string) (result int) {
	for i, n := 0, len(s); i < n; {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		result = utils.MaxNum(result, j-i)
		i = j
	}
	return
}

func main() {

}
