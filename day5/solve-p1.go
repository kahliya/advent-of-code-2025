package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P1() {
	data, _ := os.ReadFile("inputs/big.txt")

	parts := strings.Split(string(data), "\r\n\r\n")
	rawRanges, ids := strings.Fields(parts[0]), strings.Fields(parts[1])

	freshRanges := make([][2]int, len(rawRanges))

	for i, v := range rawRanges {
		parts := strings.Split(v, "-")
		low, _ := strconv.Atoi(parts[0])
		high, _ := strconv.Atoi(parts[1])
		freshRanges[i] = [2]int{low, high}
	}

	freshCounter := 0
	for _, rawId := range ids {
		currId, _ := strconv.Atoi(rawId)
		for _, rangePair := range freshRanges {
			low, high := rangePair[0], rangePair[1]
			if currId >= low && currId <= high {
				freshCounter++
				break
			}
		}
	}

	fmt.Println("Le answer iz:", freshCounter)
}
