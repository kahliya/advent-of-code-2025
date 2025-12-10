package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseMachineP2(machine string) ([]int, [][]int) {
	split1 := strings.Split(machine, "] ")
	split2 := strings.Split(split1[1], " {")

	buttons := strings.Split(split2[0], " ")
	joltage := strings.Split(split2[1][:len(split2[1])-1], ",")

	targetMatrix := make([]int, 0)
	for _, s := range joltage {
		v, _ := strconv.Atoi(s)
		targetMatrix = append(targetMatrix, v)
	}

	buttonMatrices := make([][]int, 0)
	for _, b := range buttons {
		valStrings := strings.Split(b[1:len(b)-1], ",")

		matrix := make([]int, len(targetMatrix))
		for _, s := range valStrings {
			idx, _ := strconv.Atoi(s)
			matrix[idx] = 1
		}

		buttonMatrices = append(buttonMatrices, matrix)
	}

	return targetMatrix, buttonMatrices
}

func P2() {
	data, _ := os.ReadFile("inputs/smol.txt")
	machines := strings.Split(string(data), "\r\n")

	for _, machine := range machines {
		targetMatrix, buttonMatrices := parseMachineP2(machine)
		fmt.Println(targetMatrix, "\t\t", buttonMatrices)
	}

}
