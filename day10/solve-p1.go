package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseMachineP1(machine string) (int, []int) {
	split1 := strings.Split(machine, "] ")
	split2 := strings.Split(split1[1], " {")

	lights := strings.Split(split1[0][1:], "")
	buttons := strings.Split(split2[0], " ")

	// Represent the light (target) via binary
	targetBinaryValue := 0.0
	for i, s := range lights {
		if s == "#" {
			targetBinaryValue += math.Pow(2, float64(len(lights)-1-i))
		}
	}

	// Represent the presses via binary
	buttonBinaryValues := make([]int, 0)
	for _, b := range buttons {
		powStrings := strings.Split(b[1:len(b)-1], ",")
		binaryValue := 0.0
		for _, s := range powStrings {
			x, _ := strconv.Atoi(s)
			binaryValue += math.Pow(2, float64(len(lights)-1-x))
		}
		buttonBinaryValues = append(buttonBinaryValues, int(binaryValue))
	}

	return int(targetBinaryValue), buttonBinaryValues
}

func P1() {
	data, _ := os.ReadFile("inputs/big.txt")
	machines := strings.Split(string(data), "\r\n")

	totalPresses := 0
	for _, machine := range machines {
		targetBinaryValue, buttonBinaryValues := parseMachineP1(machine)

		minPresses := len(buttonBinaryValues) + 1
		maxCombinations := int(math.Pow(2, float64(len(buttonBinaryValues))) - 1)
		for iter := 1; iter <= maxCombinations; iter++ {
			mask := fmt.Sprintf("%0*b", len(buttonBinaryValues), iter)
			current := 0
			presses := 0
			for idx := 0; idx < len(mask); idx++ {
				press := mask[idx]
				if press == '1' {
					presses++
					current ^= buttonBinaryValues[idx]
				}
			}

			if current == targetBinaryValue && minPresses > presses {
				minPresses = presses
			}
		}

		totalPresses += minPresses
	}

	fmt.Println("Le answer is:", totalPresses)
}
