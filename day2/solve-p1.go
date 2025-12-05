package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumInvalidsFromRangeP1(low, high int) int {
	invalids := []int{} // for debugging purposes
	invalidSum := 0
	for i := low; i <= high; i++ {
		numString := strconv.Itoa(i)

		// If the number of digits is odd, can't be invalid
		if len(numString)%2 != 0 {
			continue
		}

		halfLen := len(numString) / 2
		head, tail := numString[:halfLen], numString[halfLen:]
		if head == tail {
			invalids = append(invalids, i)
			invalidSum += i
		}
	}

	return invalidSum
}

func P1() {
	data, _ := os.ReadFile("inputs/big.txt")
	ranges := strings.Split(string(data), ",")

	invalidSum := 0
	for _, v := range ranges {
		pair := strings.Split(v, "-")
		low, _ := strconv.Atoi(pair[0])
		high, _ := strconv.Atoi(pair[1])

		invalidSum += sumInvalidsFromRangeP1(low, high)
	}

	fmt.Println("Le answer iz:", invalidSum)
}
