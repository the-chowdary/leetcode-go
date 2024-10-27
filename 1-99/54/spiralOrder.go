/*
Implemented spiralOrder : https://leetcode.com/problems/spiral-matrix/description/
1. get rows and cols and generate result array of length rows * cols
2. Iterate and store matrix[i][j] in result and mark it visited
3. go to next direction and check out of bounds if out of bounds change direction
4. finally update i and j

Similar problem : https://leetcode.com/problems/robot-bounded-in-circle/description/
*/
package main

import "math"

var DIRS = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // DIRS right, down, left, up
func spiralOrder(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	result := make([]int, (rows * cols))
	i, j, di := 0, 0, 0

	for k := range result {
		result[k] = matrix[i][j]
		matrix[i][j] = math.MaxInt // mark as visited

		r, c := i+DIRS[di][0], j+DIRS[di][1] // next location
		// out of bounds check
		if r < 0 || r >= rows || c < 0 || c >= cols || matrix[r][c] == math.MaxInt {
			// change direction
			di = (di + 1) % 4 // turn 90 degrees
		}
		i += DIRS[di][0]
		j += DIRS[di][1]
	}
	return result
}

// sprialOrder2 without modifying the source matrix
func sprialOrder2(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	result := make([]int, 0, (rows * cols))
	i, j := 0, -1 // start from left

	for di := 0; len(result) < cap(result); di = (di + 1) % 4 {
		for range cols {
			i += DIRS[di][0]
			j += DIRS[di][1]
			result = append(result, matrix[i][j])
		}
		cols, rows = rows-1, cols
	}
	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	spiralOrder(matrix)
	sprialOrder2(matrix)
}
