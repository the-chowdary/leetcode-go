/*
2923 : https://leetcode.com/problems/find-champion-i/description/
1. See the mountain and return
i.e sum(row) == len(grid) - 1 then return i
*/
package main

func findChampion(grid [][]int) int {
next:
	for i, row := range grid {
		for j, x := range row {
			if j != i && x == 0 {
				continue next
			}
		}
		return i
	}
	panic(-1)
}

func main() {

}
