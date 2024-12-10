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

func isValid(p Point, height, width int) bool {
	return p.x >= 0 && p.x < height && p.y >= 0 && p.y < width
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	antiNodes := make([]Point, 0)
	antiNodes2 := make([]Point, 0)
	nodes := make([]Point, 0)

	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				nodes = append(nodes, Point{i, j})
			}
		}
	}

	for _, node1 := range nodes {
		for _, node2 := range nodes {
			if node1 != node2 && lines[node1.x][node1.y] == lines[node2.x][node2.y] {
				dx := node2.x - node1.x
				dy := node2.y - node1.y

				antinode1 := Point{node2.x + dx, node2.y + dy}
				if isValid(antinode1, len(lines), len(lines[0])) && !slices.Contains(antiNodes, antinode1) {
					antiNodes = append(antiNodes, antinode1)
				}

				i := 0
				for {
					point := Point{node2.x + i*dx, node2.y + i*dy}
					if !isValid(point, len(lines), len(lines[0])) {
						break
					}
					if !slices.Contains(antiNodes2, point) {
						antiNodes2 = append(antiNodes2, point)
					}
					i++
				}

				antinode2 := Point{node1.x - dx, node1.y - dy}
				if isValid(antinode2, len(lines), len(lines[0])) && !slices.Contains(antiNodes, antinode2) {
					antiNodes = append(antiNodes, antinode2)
				}

				i = 0
				for {
					point := Point{node1.x - i*dx, node1.y - i*dy}
					if !isValid(point, len(lines), len(lines[0])) {
						break
					}
					if !slices.Contains(antiNodes2, point) {
						antiNodes2 = append(antiNodes2, point)
					}
					i++
				}
			}
		}
	}

	fmt.Println(len(antiNodes))
	fmt.Println(len(antiNodes2))
}
