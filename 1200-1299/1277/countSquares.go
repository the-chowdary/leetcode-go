/*
Implemented countSquares
1. Add first row and col to count as nothing will there before them
2. traverse matrix from and update matrix[i][j] with min(matrix[i - 1][j], min(matrix[i][j - 1], matrix[i - 1][j - 1]))
3. Update count with matrix[i][j]
*/
package main

func countSquares(matrix [][]int) (count int) {
	rows, cols := len(matrix), len(matrix[0])

	// Add first col to count
	for col := range cols {
		count += matrix[0][col]
	}

	// Add first row to count
	for i := 1; i < rows; i++ {
		count += matrix[i][0]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 1 && matrix[i-1][j] > 0 && matrix[i][j-1] > 0 && matrix[i-1][j-1] > 0 {
				matrix[i][j] += min(matrix[i-1][j], min(matrix[i][j-1], matrix[i-1][j-1]))
			}
			count += matrix[i][j]
		}
	}
	return
}

func main() {
	matrix := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}
	countSquares(matrix)
}
