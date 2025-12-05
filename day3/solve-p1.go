package main

import (
	"fmt"
	"os"
	"strings"
)

func P1() {
	data, _ := os.ReadFile("inputs/big.txt")
	banks := strings.Fields(string(data))

	totalJoltage := 0
	for _, bank := range banks {
		max1 := -1
		idx1 := -1
		max2 := -1

		// First number cannot be the last battery, stop @ len-1
		for i := 0; i < len(bank)-1; i++ {
			battNum := int(bank[i] - '0') // convert ascii char to int

			if battNum > max1 {
				max1 = battNum
				idx1 = i
			}
		}

		for i := idx1 + 1; i < len(bank); i++ {
			battNum := int(bank[i] - '0')

			if battNum > max2 {
				max2 = battNum
			}
		}

		totalJoltage += (max1*10 + max2)
	}

	fmt.Println("Le answer is:", totalJoltage)
}
