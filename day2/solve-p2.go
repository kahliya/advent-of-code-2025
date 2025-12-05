package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumInvalidsFromRangeP2(low, high int) int {
	invalids := []int{} // for debugging purposes
	invalidSum := 0
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)

		// Stop at half, that is the max possible size of the pattern
		isInvalid := false
		for j := 1; j <= len(s)/2; j++ {
			// If the length is not divisible by j, that pattern is not possible
			if len(s)%j != 0 {
				continue
			}

			pattern := s[:j]
			for k := 0; k < len(s); k += j {
				if s[k:k+j] != pattern {
					break
				}

				// Last iteration, if it reaches here, means it has found a pattern
				if k == len(s)-j {
					isInvalid = true
				}
			}
		}

		if isInvalid {
			invalids = append(invalids, i)
			invalidSum += i
		}
	}

	return invalidSum
}

func P2() {
	data, _ := os.ReadFile("inputs/big.txt")
	ranges := strings.Split(string(data), ",")

	invalidSum := 0
	for _, v := range ranges {
		pair := strings.Split(v, "-")
		low, _ := strconv.Atoi(pair[0])
		high, _ := strconv.Atoi(pair[1])

		invalidSum += sumInvalidsFromRangeP2(low, high)
	}

	fmt.Println("Le answer iz:", invalidSum)
}
