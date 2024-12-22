package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func part2() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

outerLoop:
	for it := range 10000 {
		if it%1000 == 0 {
			fmt.Println("iter prog", it)
		}
		posArr := make([]Point, 0)

		for _, line := range lines {
			p := strings.Split(strings.Split(line, " ")[0][2:], ",")
			v := strings.Split(strings.Split(line, " ")[1][2:], ",")
			px, _ := strconv.Atoi(p[0])
			py, _ := strconv.Atoi(p[1])
			vx, _ := strconv.Atoi(v[0])
			vy, _ := strconv.Atoi(v[1])

			px = (px + vx*it) % 101
			py = (py + vy*it) % 103

			if px < 0 {
				px = px + 101
			}
			if py < 0 {
				py = py + 103
			}

			px %= 101
			py %= 103

			posArr = append(posArr, Point{px, py})
		}

		grid := make(map[Point]bool)
		for _, p := range posArr {
			grid[p] = true
		}

		for y := range 94 {
			for x := range 92 {
				count := 0
				for i := 0; i < 10; i++ {
					for j := 0; j < 10; j++ {
						if grid[Point{y + i, x + j}] {
							count++
						}
					}
				}
				if count > 50 {
					fmt.Println(it)
					for y1 := range 103 {
						for x1 := range 101 {
							if grid[Point{x1, y1}] {
								fmt.Print("#")
							} else {
								fmt.Print(".")
							}
						}
						fmt.Println()
					}
					continue outerLoop
				}
			}
		}
	}
}
