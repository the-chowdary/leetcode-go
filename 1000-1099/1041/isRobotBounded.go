/*
1041 : https://leetcode.com/problems/robot-bounded-in-circle/description/
1. Define directions and if instruction is 'G' then update x, y with direction
2. If instruction is 'L' then turn counter clockwise else clockwise

Similar Problem : https://leetcode.com/problems/spiral-matrix/description/
*/
package main

// DIRS is east, south, west, north
var DIRS = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func isRobotBounded(instructions string) bool {
	x, y, d := 0, 0, 0
	for _, instruction := range instructions {
		if instruction == 'G' {
			x += DIRS[d][0]
			y += DIRS[d][1]
		} else {
			if instruction == 'L' {
				d = (d + 3) % 4 // turn counter clock wise
			} else {
				d = (d + 1) % 4 // turn clock wise
			}
		}
	}
	return (x == 0 && y == 0) || (d != 0)
}

func main() {
	instructions := "GGLLGG"
	isRobotBounded(instructions)
}
