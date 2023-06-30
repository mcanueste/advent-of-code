package main

import (
	"fmt"
	"os"
	"strconv"
)

func readInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func moveNorth(loc []int) []int {
	loc[0] += 1
	return loc
}

func moveSouth(loc []int) []int {
	loc[0] -= 1
	return loc
}

func moveEast(loc []int) []int {
	loc[1] += 1
	return loc
}

func moveWest(loc []int) []int {
	loc[1] -= 1
	return loc
}

func makeHash(loc []int) string {
	return strconv.Itoa(loc[0]) + "-" + strconv.Itoa(loc[1])
}

func visit(hist map[string]int, loc []int) {
	hash := makeHash(loc)
	val, exists := hist[hash]
	if exists {
		hist[hash] = val + 1
	} else {
		hist[hash] = 1
	}
}

func atLeastOne(hist map[string]int) int {
	count := 0
	for _, val := range hist {
		if val > 0 {
			count += 1
		}
	}
	return count
}

func traverse(moves string, location []int, history map[string]int) {
	visit(history, location)
	for _, move := range moves {
		switch move {
		case '^':
			location = moveNorth(location)
			visit(history, location)
		case 'v':
			location = moveSouth(location)
			visit(history, location)
		case '>':
			location = moveEast(location)
			visit(history, location)
		case '<':
			location = moveWest(location)
			visit(history, location)
		default:
			fmt.Println("uknown move" + string(move))
		}
	}
}

func splitMoves(start int, moves string) string {
	newMoves := ""
	for i := start; i < len(moves); i = i + 2{
		newMoves += string(moves[i])
	}
	return newMoves
}

func p1(moves string) int {
	location := []int{0, 0}
	history := make(map[string]int)
	traverse(moves, location, history)
	return atLeastOne(history)
}

func p2(moves string) int {
	history := make(map[string]int)
	santaLocation := []int{0, 0}
	santaMoves:= splitMoves(0, moves)
	robotLocation := []int{0, 0}
	robotMoves := splitMoves(1, moves)
	traverse(santaMoves, santaLocation, history)
	traverse(robotMoves, robotLocation, history)
	return atLeastOne(history)
}

func main() {
	moves := readInput("./input.txt")
	p1Solution := p1(moves)
	p2Solution := p2(moves)
	fmt.Println(p1Solution)
	fmt.Println(p2Solution)
}
