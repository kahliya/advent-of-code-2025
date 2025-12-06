package main

func main() {
	P2()
}

var DIRECTIONS = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1}, // UP-LEFT, UP, UP-RIGHT
	{0, -1}, {0, 1}, // LEFT, RIGHT
	{1, -1}, {1, 0}, {1, 1}, // DOWN-LEFT, DOWN, DOWN-RIGHT
}

var ROLL = "@"

func checkIfAccessible(layout [][]byte, rowIdx, colIdx int) bool {
	rollsFound := 0
	for _, d := range DIRECTIONS {
		tmpRowIdx := rowIdx + d[0]
		tmpColIdx := colIdx + d[1]

		// Skip if idx out of bounds
		if tmpRowIdx < 0 || tmpRowIdx >= len(layout) || tmpColIdx < 0 || tmpColIdx >= len(layout[0]) {
			continue
		}

		if string(layout[tmpRowIdx][tmpColIdx]) == ROLL {
			rollsFound++
		}
	}

	if rollsFound < 4 {
		return true
	}

	return false
}
