/*
Implemented numIslands
1. Get rows and cols length
2. Iterate the grid and if grid[i][j] == 1 then dfs(i, j) and result++
3. dfs() -> check out of bounds cases and if grid[i][j] != 1 then return
4. Mark grid[i][j] = 1 -> Making it as visited
5. dfs in all directions(left, right top, bottom)
*/
package main

func numIslands(grid [][]byte) (result int) {
	rows, cols := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '1' {
				dfs(i, j)
				result++
			}
		}
	}
	return result
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	numIslands(grid)
}
