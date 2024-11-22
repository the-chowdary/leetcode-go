/*
1072 : https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/description
*/
package main

func maxEqualRowsAfterFlips(matrix [][]int) (result int) {
	m := len(matrix[0])
	mp := map[string]int{}
	s1, s2 := make([]byte, m), make([]byte, m)

	for _, row := range matrix {
		for j, mij := range row {
			s1[j] = byte(mij)
			s2[j] = byte(mij) ^ 1
		}
		mp[string(s1)]++
		mp[string(s2)]++
	}
	for _, v := range mp {
		result = max(result, v)
	}
	return
}

func main() {

}
