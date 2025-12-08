package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func P2BruteForce() {
	data, _ := os.ReadFile("inputs/big.txt")
	points := strings.Fields(string(data))

	distances := make([]pairDist, 0)
	for i := 0; i < len(points); i++ {
		x1, y1, z1 := getXYZFromPoint(points[i])
		for j := i + 1; j < len(points); j++ {
			x2, y2, z2 := getXYZFromPoint(points[j])
			dist := get3DEuclideanDist(x1, y1, z1, x2, y2, z2)
			distances = append(distances, pairDist{i, j, dist})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	circuits := make([][]int, 0)
	finalPair := pairDist{}
	for _, pair := range distances[:5492] {
		nextCircuits := make([][]int, 0)

		// Use map keys as a "Set" to remove duplicates
		// Values dont matter
		connectedBoxes := map[int]string{
			pair.idx1: "",
			pair.idx2: "",
		}

		for _, circuit := range circuits {
			if slices.Contains(circuit, pair.idx1) || slices.Contains(circuit, pair.idx2) {
				for _, v := range circuit {
					connectedBoxes[v] = ""
				}
			} else {
				nextCircuits = append(nextCircuits, circuit)
			}
		}

		newCircuit := make([]int, 0)
		for k := range connectedBoxes {
			newCircuit = append(newCircuit, k)
		}
		nextCircuits = append(nextCircuits, newCircuit)
		circuits = nextCircuits

		if len(circuits) == 1 && len(circuits[0]) == len(points) {
			finalPair = pair
			break
		}
	}

	finalX1, _, _ := getXYZFromPoint(points[finalPair.idx1])
	finalX2, _, _ := getXYZFromPoint(points[finalPair.idx2])
	fmt.Println("Le answer is:", finalX1*finalX2)
}
