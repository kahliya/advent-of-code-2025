package main

import (
	"bytes"
	"fmt"
	"os"
)

func P1() {
	data, _ := os.ReadFile("inputs/big.txt")
	layout := bytes.Split(data, []byte("\r\n"))

	numAccessible := 0
	for rowIdx, row := range layout {
		for colIdx, item := range row {
			if string(item) != ROLL {
				continue
			}

			if checkIfAccessible(layout, rowIdx, colIdx) {
				numAccessible++
			}
		}
	}

	fmt.Println("Le answer is:", numAccessible)
}
