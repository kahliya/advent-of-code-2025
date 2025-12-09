package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func P1() {
	data, _ := os.ReadFile("inputs/big.txt")
	dataRows := strings.Fields(string(data))

	points := make([][]int, 0)
	for _, point := range dataRows {
		pair := strings.Split(point, ",")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		points = append(points, []int{x, y})
	}

	maxArea := -1.0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			area := (math.Abs(float64(x1-x2)) + 1) * (math.Abs(float64(y1-y2)) + 1)

			if area > maxArea {
				maxArea = area
			}

			// fmt.Printf("(%d, %d) & (%d, %d) = Area <%f>\n", x1, y1, x2, y2, area)
		}
	}

	fmt.Printf("Le answer is: %f\n", maxArea)
}
