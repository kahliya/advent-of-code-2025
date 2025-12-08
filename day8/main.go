package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {
	P2BruteForce()
}

type pairDist struct {
	idx1, idx2 int
	dist       float64
}

func getXYZFromPoint(point string) (int, int, int) {
	triplet := strings.Split(point, ",")
	x, _ := strconv.Atoi(triplet[0])
	y, _ := strconv.Atoi(triplet[1])
	z, _ := strconv.Atoi(triplet[2])

	return x, y, z
}

func get3DEuclideanDist(x1, y1, z1, x2, y2, z2 int) float64 {
	fx1 := float64(x1)
	fy1 := float64(y1)
	fz1 := float64(z1)
	fx2 := float64(x2)
	fy2 := float64(y2)
	fz2 := float64(z2)
	return math.Sqrt(math.Pow((fx1-fx2), 2) + math.Pow((fy1-fy2), 2) + math.Pow((fz1-fz2), 2))
}
