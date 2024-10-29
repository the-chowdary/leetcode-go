/*
2684 : https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/description
1. For every col(i, 0) iterate to r - 1, r, r + 1
2. DFS the solution
*/
package main

func maxMoves(grid [][]int) (result int) {
	// DFS the solution
	rows, cols := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(r, c int) {
		result = max(result, c)
		if result == cols-1 {
			return
		}

		for row := max(r-1, 0); row < min(r+2, rows); row++ {
			if grid[row][c+1] > grid[r][c] {
				dfs(row, c+1)
			}
		}

		grid[r][c] = 0 // visited
	}

	for r := range rows {
		dfs(r, 0)
	}
	return
}

func main() {
	grid := [][]int{
		{2, 4, 3, 5},
		{5, 4, 9, 3},
		{3, 4, 2, 11},
	}
	maxMoves(grid)
}
