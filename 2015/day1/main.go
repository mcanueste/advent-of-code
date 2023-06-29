package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func p1(input string) int {
	left := strings.Count(input, "(")
	right := strings.Count(input, ")")
	return left - right
}

func p2(input string) int {
	count := 0
	for pos, char := range input {
		if char == '(' {
			count += 1
		} else {
			count -= 1
		}

		if count < 0 {
			return pos + 1
		}
	}
	return 0
}

func main() {
	input := readInput("./input.txt")
	p1_solution := p1(input)
	p2_solution := p2(input)
	fmt.Println(p1_solution)
	fmt.Println(p2_solution)
}
