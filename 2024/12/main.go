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

func findAdj(lines []string, visited map[Point]bool, char rune, x int, y int, area int, perimeter int, points []Point) (int, int, []Point) {
	if x < 0 || y < 0 || y >= len(lines) || x >= len(lines[y]) {
		return area, perimeter, points
	}

	if visited[Point{x, y}] || rune(lines[y][x]) != char {
		return area, perimeter, points
	}

	visited[Point{x, y}] = true
	points = append(points, Point{x, y})

	area++

	if y-1 < 0 || rune(lines[y-1][x]) != char {
		perimeter++
	} else {
		area, perimeter, points = findAdj(lines, visited, char, x, y-1, area, perimeter, points)
	}
	if y+1 >= len(lines) || rune(lines[y+1][x]) != char {
		perimeter++
	} else {
		area, perimeter, points = findAdj(lines, visited, char, x, y+1, area, perimeter, points)
	}
	if x-1 < 0 || rune(lines[y][x-1]) != char {
		perimeter++
	} else {
		area, perimeter, points = findAdj(lines, visited, char, x-1, y, area, perimeter, points)
	}
	if x+1 >= len(lines[y]) || rune(lines[y][x+1]) != char {
		perimeter++
	} else {
		area, perimeter, points = findAdj(lines, visited, char, x+1, y, area, perimeter, points)
	}

	return area, perimeter, points
}

func calculateSides(points []Point) int {
	sides := 0
	for _, point := range points {
		x := point.x
		y := point.y
		if !slices.Contains(points, Point{x + 1, y}) && !slices.Contains(points, Point{x, y - 1}) {
			sides++
		}
		if !slices.Contains(points, Point{x + 1, y}) && !slices.Contains(points, Point{x, y + 1}) {
			sides++
		}
		if !slices.Contains(points, Point{x - 1, y}) && !slices.Contains(points, Point{x, y - 1}) {
			sides++
		}
		if !slices.Contains(points, Point{x - 1, y}) && !slices.Contains(points, Point{x, y + 1}) {
			sides++
		}
		if slices.Contains(points, Point{x + 1, y}) && slices.Contains(points, Point{x, y + 1}) && !slices.Contains(points, Point{x + 1, y + 1}) {
			sides++
		}
		if slices.Contains(points, Point{x - 1, y}) && slices.Contains(points, Point{x, y + 1}) && !slices.Contains(points, Point{x - 1, y + 1}) {
			sides++
		}
		if slices.Contains(points, Point{x + 1, y}) && slices.Contains(points, Point{x, y - 1}) && !slices.Contains(points, Point{x + 1, y - 1}) {
			sides++
		}
		if slices.Contains(points, Point{x - 1, y}) && slices.Contains(points, Point{x, y - 1}) && !slices.Contains(points, Point{x - 1, y - 1}) {
			sides++
		}
	}
	return sides
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	visited := make(map[Point]bool)
	output := 0
	output1 := 0

	for y, line := range lines {
		for x, char := range line {
			if !visited[Point{x, y}] {
				area, perimeter, points := findAdj(lines, visited, char, x, y, 0, 0, make([]Point, 0))
				output += area * perimeter
				output1 += area * calculateSides(points)
			}
		}
	}

	fmt.Println(output)
	fmt.Println(output1)
}
