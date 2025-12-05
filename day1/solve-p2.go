package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P2() {
	data, _ := os.ReadFile("inputs/big-in.txt")

	zeroCounter := 0
	pointer := 50
	for _, v := range strings.Fields(string(data)) {
		direction := string(v[0])
		amnt, _ := strconv.Atoi((v[1:]))

		// Check times looped
		if amnt > 99 {
			zeroCounter += amnt / 100
			amnt %= 100
		}

		// Get absolute pointer value
		nextPtr := 0
		if direction == "L" {
			nextPtr = (pointer - amnt) % 100
			if nextPtr < 0 {
				nextPtr = 100 + nextPtr
			}

			if pointer != 0 && nextPtr > pointer {
				zeroCounter++
			}

		} else if direction == "R" {
			nextPtr = (pointer + amnt) % 100
			if nextPtr != 0 && nextPtr < pointer {
				zeroCounter++
			}
		}

		if pointer != 0 && nextPtr == 0 {
			zeroCounter++
		}

		pointer = nextPtr
	}

	fmt.Printf("Le answer iz: %d\n", zeroCounter)
}
