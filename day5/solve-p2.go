package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func P2() {
	data, _ := os.ReadFile("inputs/big.txt")

	parts := strings.Split(string(data), "\r\n\r\n")
	rawRanges := strings.Fields(parts[0])

	freshRanges := make([][2]int, len(rawRanges))

	for i, v := range rawRanges {
		parts := strings.Split(v, "-")
		low, _ := strconv.Atoi(parts[0])
		high, _ := strconv.Atoi(parts[1])
		freshRanges[i] = [2]int{low, high}
	}

	// Sort freshRanges by "low" ascending
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0]
	})

	// Solve the overlapping ranges
	finalRanges := [][2]int{freshRanges[0]}
	for _, v := range freshRanges[1:] {
		prev := &finalRanges[len(finalRanges)-1]
		_, prevHigh := &prev[0], &prev[1]
		low, high := v[0], v[1]

		// Case #1: Low overlaps into prev range, extend the range
		// Case #2: Low overlaps into prev range, high also overlaps, ignore
		if low <= *prevHigh && *prevHigh < high {
			*prevHigh = high
		}

		// Case #3: Low does not overlap, new range
		if low > *prevHigh {
			finalRanges = append(finalRanges, v)
		}
	}

	// Count the number of fresh ids
	freshCounter := 0
	for _, v := range finalRanges {
		low, high := v[0], v[1]
		freshCounter += high - low + 1
	}

	fmt.Println("Le answer iz:", freshCounter)
}
