package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func P1() {
	data, _ := os.ReadFile("inputs/big.txt")
	rows := bytes.Fields(data)

	timesSplit := 0
	beamColumns := []int{bytes.Index(rows[0], []byte("S"))}
	for _, row := range rows[1:] {
		nextBeamColumns := []int{}
		for _, col := range beamColumns {
			if row[col] != SPLITTER {
				if !slices.Contains(nextBeamColumns, col) {
					nextBeamColumns = append(nextBeamColumns, col)
				}
				continue
			}
			timesSplit++

			left, right := col-1, col+1
			if left > -1 && !slices.Contains(nextBeamColumns, left) {
				nextBeamColumns = append(nextBeamColumns, left)
			}
			if right < len(row) && !slices.Contains(nextBeamColumns, right) {
				nextBeamColumns = append(nextBeamColumns, right)
			}
		}

		beamColumns = nextBeamColumns
	}

	fmt.Println("Le answer is:", timesSplit)
}
