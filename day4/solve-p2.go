package main

import (
	"bytes"
	"fmt"
	"os"
)

func P2() {
	data, _ := os.ReadFile("inputs/big.txt")
	layout := bytes.Split(data, []byte("\r\n"))

	numAccessible := 0
	for {
		numAccessibleThisIter := 0
		for rowIdx, row := range layout {
			for colIdx, item := range row {
				if string(item) != ROLL {
					continue
				}

				if checkIfAccessible(layout, rowIdx, colIdx) {
					numAccessible++
					numAccessibleThisIter++
					layout[rowIdx][colIdx] = '.'
				}
			}
		}

		if numAccessibleThisIter == 0 {
			break
		}
	}

	fmt.Println("Le answer is:", numAccessible)
}
