package main

import (
	"fmt"
	"os"
	"strings"
)

func P2() {
	data, _ := os.ReadFile("inputs/big.txt")
	banks := strings.Fields(string(data))

	totalJoltage := 0
	for _, bank := range banks {
		maxJolt := make([]int, 12)
		prevIdx := -1

		// Iterate a total of 12 times to find 12 digits
		for iter := 0; iter < 12; iter++ {
			// idxLimit is the number of remaining digits required
			idxLimit := len(bank) - (11 - iter)
			for i := prevIdx + 1; i < idxLimit; i++ {
				battNum := int(bank[i] - '0')
				if battNum > maxJolt[iter] {
					maxJolt[iter] = battNum
					prevIdx = i
				}
			}
		}

		// Build the bankJoltage digit
		bankJoltage := 0
		for _, d := range maxJolt {
			bankJoltage = bankJoltage*10 + d
		}

		totalJoltage += bankJoltage
	}

	fmt.Println("Le answer is:", totalJoltage)
}
