package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	x, y int
}

var visitedNines = make([]Point, 0)

func checkSides(length int, width int, lines []string, x int, y int) int {
	output := 0
	if lines[y][x] == '9' {
		if !slices.Contains(visitedNines, Point{x, y}) {
			visitedNines = append(visitedNines, Point{x, y})
			return 1
		}
		return 0
	}

	char := lines[y][x]

	if y-1 >= 0 && int(lines[y-1][x]) == int(char)+1 {
		output += checkSides(length, width, lines, x, y-1)
	}

	if y+1 < length && int(lines[y+1][x]) == int(char)+1 {
		output += checkSides(length, width, lines, x, y+1)
	}

	if x-1 >= 0 && int(lines[y][x-1]) == int(char)+1 {
		output += checkSides(length, width, lines, x-1, y)
	}

	if x+1 < width && int(lines[y][x+1]) == int(char)+1 {
		output += checkSides(length, width, lines, x+1, y)
	}

	return output
}

// id make this one function but im too lazy i need to study
func checkSides1(length int, width int, lines []string, x int, y int) int {
	output := 0
	if lines[y][x] == '9' {
		return 1
	}

	char := lines[y][x]

	if y-1 >= 0 && int(lines[y-1][x]) == int(char)+1 {
		output += checkSides1(length, width, lines, x, y-1)
	}

	if y+1 < length && int(lines[y+1][x]) == int(char)+1 {
		output += checkSides1(length, width, lines, x, y+1)
	}

	if x-1 >= 0 && int(lines[y][x-1]) == int(char)+1 {
		output += checkSides1(length, width, lines, x-1, y)
	}

	if x+1 < width && int(lines[y][x+1]) == int(char)+1 {
		output += checkSides1(length, width, lines, x+1, y)
	}

	return output
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	output := 0
	output1 := 0

	for y, line := range lines {
		for x, char := range line {
			if char == '0' {
				visitedNines = make([]Point, 0)
				output += checkSides(len(lines), len(line), lines, x, y)
				output1 += checkSides1(len(lines), len(line), lines, x, y)
			}
		}
	}

	fmt.Println(output)
	fmt.Println(output1)
}
