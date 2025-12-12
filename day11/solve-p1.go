package main

import (
	"fmt"
	"os"
	"strings"
)

func countPathsOutP1(paths map[string][]string, next string) int {
	possiblePaths := paths[next]
	totalPaths := 0
	for _, path := range possiblePaths {
		if path == "out" {
			return 1
		}

		totalPaths += countPathsOutP1(paths, path)
	}

	return totalPaths
}

func P1() {
	data, _ := os.ReadFile("inputs/big.txt")
	devices := strings.Split(string(data), "\r\n")

	paths := map[string][]string{}
	for _, device := range devices {
		tmp := strings.Split(device, ": ")
		src, tmp2 := tmp[0], tmp[1]
		dst := strings.Fields(tmp2)
		paths[src] = dst
	}

	rez := countPathsOutP1(paths, "you")
	fmt.Println("Le answer is:", rez)
}
