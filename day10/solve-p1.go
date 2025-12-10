package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseMachine(machine string) ([]string, [][]int) {
	split1 := strings.Split(machine, "] ")
	split2 := strings.Split(split1[1], " {")

	lights := strings.Split(split1[0][1:], "")
	rawButtons := strings.Split(split2[0], " ")

	toggles := make([][]int, 0)
	for _, b := range rawButtons {
		togglesString := strings.Split(b[1:len(b)-1], ",")
		tmp := make([]int, 0)
		for _, s := range togglesString {
			t, _ := strconv.Atoi(s)
			tmp = append(tmp, t)
		}
		toggles = append(toggles, tmp)
	}

	return lights, toggles
}

func P1() {
	data, _ := os.ReadFile("inputs/smol.txt")
	machines := strings.Split(string(data), "\r\n")

	for _, machine := range machines {
		lights, toggles := parseMachine(machine)
		fmt.Println(lights, toggles)

		maxCombinations := 
	}
}
