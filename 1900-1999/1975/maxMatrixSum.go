/*
1975 : https://leetcode.com/problems/maximum-matrix-sum/description
*/
package main

func maxMatrixSum(matrix [][]int) int64 {
	sum, min, neg := 0, int(1e9), false
	for _, r := range matrix {
		for _, v := range r {
			if v < 0 {
				neg = !neg
				v = -v
			}
			if v < min {
				min = v
			}
			sum += v
		}
	}
	if neg {
		sum -= min * 2
	}
	return int64(sum)
}

func main() {

}
