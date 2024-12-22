package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, line := range lines {
		p := strings.Split(strings.Split(line, " ")[0][2:], ",")
		v := strings.Split(strings.Split(line, " ")[1][2:], ",")
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])

		px = (px + vx*100) % 101
		py = (py + vy*100) % 103

		if px < 0 {
			px = px + 101
		}
		if py < 0 {
			py = py + 103
		}

		px %= 101
		py %= 103

		if py < 51 {
			if px < 50 {
				q1++
			} else if px > 50 {
				q2++
			}
		} else if py > 51 {
			if px < 50 {
				q3++
			} else if px > 50 {
				q4++
			}
		}
	}

	fmt.Println(q1 * q2 * q3 * q4)
}
