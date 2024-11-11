/*
Implemented 874 : https://leetcode.com/problems/walking-robot-simulation/description/
1. Maintain obstacleSet and iterate over commands
2. If command is -2 turn left if it is -1 turn right else
3. if x + i, y + j not in obstacleSet get (x * x + y * y)
4. return max(result, distance)

Similar Problems:
54 : https://leetcode.com/problems/spiral-matrix/description/
1041 : https://leetcode.com/problems/robot-bounded-in-circle/description/
*/
package main

// DIRS east, south, west, north
var DIRS = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func robotSim(commands []int, obstacles [][]int) int {
	obstacleSet := make(map[[2]int]struct{})
	for _, obs := range obstacles {
		obstacleSet[[2]int{obs[0], obs[1]}] = struct{}{}
	}

	x, y, d, result := 0, 0, 0, 0

	for _, c := range commands {
		if c == -2 {
			// Turn left
			d = (d + 3) % 4
		} else if c == -1 {
			// Turn right
			d = (d + 1) % 4
		} else {
			i, j := DIRS[d][0], DIRS[d][1]
			for c > 0 {
				if _, found := obstacleSet[[2]int{x + i, y + j}]; found {
					break
				}
				x += i
				y += j
				c--
				distance := (x*x + y*y)
				if distance > result {
					result = distance
				}
			}
		}
	}
	return result
}

func main() {
	commands := []int{4, -1, 3}
	obstacles := [][]int{}
	robotSim(commands, obstacles)
}
