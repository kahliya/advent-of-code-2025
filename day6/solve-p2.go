package main

import (
	"fmt"
	"os"
	"strings"
)

func P2() {
	data, _ := os.ReadFile("inputs/big.txt")
	rows := strings.Split(string(data), "\r\n")

	// Process column by column
	// Right-to-left can be ignored
	RESET_VALUE := -1        // safe to use negative, since all operands are positive
	finalSum := -RESET_VALUE // initialized as negative RESET_VALUE to offset the first "reset", I know it's hacky but it works :D
	setSum := RESET_VALUE

	lastRowIdx := len(rows) - 1
	operator := ""
	for i := 0; i < len(rows[0]); i++ {
		operatorRowVal := string(rows[lastRowIdx][i])

		// This is the first "reset" I'm talking about
		// If operator row contains an operator, it is a new calculation
		if operatorRowVal != " " {
			finalSum += setSum
			setSum = RESET_VALUE
			operator = operatorRowVal
		}

		// Iterate until second last row (less operator row)
		operand := 0
		ignoreColumn := true
		for _, row := range rows[:lastRowIdx] {
			if string(row[i]) == " " {
				continue
			}

			digit := int(row[i] - '0')
			operand = addNextDigit(operand, digit)

			// Column will only be ignored if it's empty
			ignoreColumn = false
		}

		// If its the first operand of an operation, ignore the operation
		if setSum == RESET_VALUE {
			setSum = operand
			ignoreColumn = true
		}

		if !ignoreColumn {
			setSum = performOperation(setSum, operand, operator)
		}
	}

	// Add the final operation
	finalSum += setSum

	fmt.Println("Le answer is:", finalSum)
}
