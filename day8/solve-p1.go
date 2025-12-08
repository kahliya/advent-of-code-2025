package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func P1BruteForce() {
	data, _ := os.ReadFile("inputs/big.txt")
	points := strings.Fields(string(data))

	NUM_CONNECTIONS := 1000 // REMEMBER TO SET THIS! From the scenario, smol = 10, big = 1000

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
	for _, pair := range distances[:NUM_CONNECTIONS] {
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
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	result := len(circuits[0]) * len(circuits[1]) * len(circuits[2])
	fmt.Println("Le answer is:", result)
}
