package main

import (
	"fmt"
	"os"
	"strings"
)

// Map of strings of []int
// For the []int, 0 = DAC only, 1 = FFT only, 2 = Both visited
var KNOWN_PATHS = make(map[string]map[string]int)

func countPathsOutP2(paths map[string][]string, next string, dac, fft bool) int {
	visitStatus := fmt.Sprintf("%t-%t", dac, fft)
	answer, ok := KNOWN_PATHS[next][visitStatus]
	if ok {
		return answer
	}

	possiblePaths := paths[next]
	totalPaths := 0
	for _, path := range possiblePaths {
		thisDac := dac
		thisFft := fft
		switch path {
		case "out":
			if dac && fft {
				return 1
			}
			return 0
		case "dac":
			thisDac = true
		case "fft":
			thisFft = true
		}
		totalPaths += countPathsOutP2(paths, path, thisDac, thisFft)
	}

	_, ok2 := KNOWN_PATHS[next]
	if !ok2 {
		KNOWN_PATHS[next] = map[string]int{}
	}

	KNOWN_PATHS[next][visitStatus] = totalPaths
	return totalPaths
}

func P2() {
	data, _ := os.ReadFile("inputs/big.txt")
	devices := strings.Split(string(data), "\r\n")

	paths := map[string][]string{}
	for _, device := range devices {
		tmp := strings.Split(device, ": ")
		src, tmp2 := tmp[0], tmp[1]
		dst := strings.Fields(tmp2)
		paths[src] = dst
	}

	rez := countPathsOutP2(paths, "svr", false, false)
	fmt.Println("Le answer is:", rez)
}
