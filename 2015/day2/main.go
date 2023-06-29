package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type box struct {
	l, w, h int
}

func (b box) surfaceArea() int {
	return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l
}

func (b box) volume() int {
	return b.l * b.w * b.h
}

func (b box) sortSides() []int {
	sides := []int{b.l, b.w, b.h}
	sort.Slice(sides, func(i, j int) bool {
		return sides[i] < sides[j]
	})
	return sides
}

func (b box) slack() int {
	sides := b.sortSides()
	return sides[0] * sides[1]
}

func (b box) smallestFacePerimeter() int {
	sides := b.sortSides()
	return (2 * sides[0]) + (2 * sides[1])
}

func (b box) paperSize() int {
	return b.surfaceArea() + b.slack()
}

func (b box) ribbonSize() int {
	return b.volume() + b.smallestFacePerimeter()
}

func newBox(sizes string) *box {
	splitted := strings.Split(sizes, "x")
	l, err := strconv.Atoi(splitted[0])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(splitted[1])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(splitted[2])
	if err != nil {
		panic(err)
	}
	return &box{
		l: l,
		w: w,
		h: h,
	}
}

func readInput(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func convertToBoxes(lines []string) []*box {
	boxes := make([]*box, len(lines))
	for idx, line := range lines {
		boxes[idx] = newBox(line)
	}
	return boxes
}

func p1(boxes []*box) int {
	sum := 0
	for _, box := range boxes {
		sum += box.paperSize()
	}
	return sum
}

func p2(boxes []*box) int {
	sum := 0
	for _, box := range boxes {
		sum += box.ribbonSize()
	}
	return sum
}

func main() {
	lines := readInput("./input.txt")
	boxes := convertToBoxes(lines)
	p1Solution := p1(boxes)
	p2Solution := p2(boxes)
	fmt.Println(p1Solution)
	fmt.Println(p2Solution)
}
