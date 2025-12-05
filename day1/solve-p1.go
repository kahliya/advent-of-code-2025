package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P1() {
	data, _ := os.ReadFile("inputs/big-in.txt")

	zeroCounter := 0
	pointer := 50
	for _, v := range strings.Fields(string(data)) {
		direction := string(v[0])
		amnt, _ := strconv.Atoi((v[1:]))

		if direction == "L" {
			pointer = (pointer - amnt) % 100
		} else if direction == "R" {
			pointer = (pointer + amnt) % 100
		}

		if pointer < 0 {
			pointer = 100 + pointer
		}

		if pointer == 0 {
			zeroCounter++
		}
	}

	fmt.Printf("Ze answer iz: %d\n", zeroCounter)
}
