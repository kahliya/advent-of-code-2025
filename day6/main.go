package main

import "fmt"

func main() {
	P1()
}

func performOperation(accumulator, x int, operator string) int {
	switch operator {
	case "+":
		return accumulator + x
	case "*":
		return accumulator * x
	default:
		fmt.Println("UNKNOWN OPERATOR!")
		return -676767
	}
}

func addNextDigit(accumulator, nextDigit int) int {
	return accumulator*10 + nextDigit // yes, think about it
}
