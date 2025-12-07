package main

import (
	"bytes"
	"fmt"
	"os"
)

var knownPaths = make(map[string]int)

func exploreRow(rows [][]byte, rowIdx, beamCol int) int {
	PATH_NAME := fmt.Sprintf("%d-%d", rowIdx, beamCol)

	// If previously explored, return the cached result
	result, ok := knownPaths[PATH_NAME]
	if ok {
		return result
	}

	if rowIdx == len(rows)-1 {
		return 1
	}

	row := rows[rowIdx]

	// Don't split, go down
	if row[beamCol] != SPLITTER {
		result = exploreRow(rows, rowIdx+1, beamCol)
		knownPaths[PATH_NAME] = result
		return result
	}

	// Split, go left & right
	result = 0
	left, right := beamCol-1, beamCol+1
	if left > -1 {
		result += exploreRow(rows, rowIdx+1, left)
	}
	if right < len(row) {
		result += exploreRow(rows, rowIdx+1, right)
	}

	return result
}

func P2() {
	data, _ := os.ReadFile("inputs/big.txt")
	rows := bytes.Fields(data)

	beamCol := bytes.Index(rows[0], []byte("S"))
	totalTimelines := exploreRow(rows, 1, beamCol)

	fmt.Println("Le answer is:", totalTimelines)
}
