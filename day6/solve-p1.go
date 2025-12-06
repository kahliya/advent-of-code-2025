package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P1() {
	data, _ := os.ReadFile("inputs/big.txt")
	dataRows := strings.Split(string(data), "\r\n")

	rows := make([][]string, len(dataRows))
	for i, row := range dataRows {
		rows[i] = strings.Fields(row)
	}

	// Process column by column
	finalSum := 0
	for i := 0; i < len(rows[0]); i++ {
		lastRowIdx := len(rows) - 1
		operator := rows[lastRowIdx][i]

		// Iterate until second last row
		columnRez, _ := strconv.Atoi(rows[0][i])
		for _, row := range rows[1:lastRowIdx] {
			intVal, _ := strconv.Atoi(row[i])
			columnRez = performOperation(columnRez, intVal, operator)
		}

		finalSum += columnRez
	}

	fmt.Println("Le answer is:", finalSum)
}
