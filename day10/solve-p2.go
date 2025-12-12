package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseMachineP2(machine string) ([]int, [][]int) {
	split1 := strings.Split(machine, "] ")
	split2 := strings.Split(split1[1], " {")

	buttons := strings.Split(split2[0], " ")
	joltage := strings.Split(split2[1][:len(split2[1])-1], ",")

	target := make([]int, 0)
	for _, s := range joltage {
		v, _ := strconv.Atoi(s)
		target = append(target, v)
	}

	buttonMatrices := make([][]int, 0)
	for _, b := range buttons {
		valStrings := strings.Split(b[1:len(b)-1], ",")

		matrix := make([]int, len(target))
		for _, s := range valStrings {
			idx, _ := strconv.Atoi(s)
			matrix[idx] = 1
		}

		buttonMatrices = append(buttonMatrices, matrix)
	}

	return target, buttonMatrices
}

// Idea: BFS with heuristics to reduce complexity
// - H1: End tree if any counter overshoots
// 		- Potentially: Stop trying this button in the future if it causes H1 from "current" state

// Complexity too high, doesn't work

func P2BFSFailed() {
	data, _ := os.ReadFile("inputs/smol.txt")
	machines := strings.Split(string(data), "\r\n")

	for _, machine := range machines {
		target, buttons := parseMachineP2(machine)
		fmt.Println("----\n", target, buttons)

		queue := Queue{}
		queue.push(State{make([]int, len(target)), []int{}, 0})

		for soln := -1; soln < 0; {
			s := queue.pop()
			excludedButtons := make([]int, len(s.excludedButtons))
			copy(excludedButtons, s.excludedButtons)

			nextCounters := make([][]int, 0)
			for btnIdx, button := range buttons {
				if slices.Contains(excludedButtons, btnIdx) {
					continue
				}

				impossible := false

				thisCounter := make([]int, len(s.counter))
				copy(thisCounter, s.counter)

				for i := 0; i < len(thisCounter); i++ {
					if button[i] == 1 {
						thisCounter[i]++
					}

					if thisCounter[i] > target[i] {
						impossible = true
					}
				}

				if impossible {
					// fmt.Println("found impossible @", s.counters, s.nextPress, "(", s.presses+1, "presses so far )")
					excludedButtons = append(excludedButtons, btnIdx)
					continue
				}

				if slices.Equal(thisCounter, target) {
					fmt.Println("found soln:", s.presses+1)
					soln = s.presses + 1
					break
				}

				nextCounters = append(nextCounters, thisCounter)
			}

			for _, c := range nextCounters {
				queue.push(State{c, excludedButtons, s.presses + 1})
			}
		}
	}
}
