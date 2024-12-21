package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solveLinearEquations(x1, x2, x3, y1, y2, y3 float64) (float64, float64) {
	det := x1*y2 - x2*y1

	x := (x3*y2 - y3*y1) / det
	y := (x1*y3 - x2*x3) / det

	return float64(x), float64(y)
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	output := 0
	output1 := 0

	for i := range lines {
		if i%4 == 0 {
			x1, _ := strconv.Atoi(strings.Split(strings.Split(lines[i], "+")[1], ", ")[0])
			y1, _ := strconv.Atoi(strings.Split(strings.Split(lines[i], "+")[2], ", ")[0])
			x2, _ := strconv.Atoi(strings.Split(strings.Split(lines[i+1], "+")[1], ", ")[0])
			y2, _ := strconv.Atoi(strings.Split(strings.Split(lines[i+1], "+")[2], ", ")[0])
			x3, _ := strconv.Atoi(strings.Split(strings.Split(lines[i+2], "=")[1], ", ")[0])
			y3, _ := strconv.Atoi(strings.Split(strings.Split(lines[i+2], "=")[2], ", ")[0])
			x, y := solveLinearEquations(float64(x1), float64(y1), float64(x3), float64(x2), float64(y2), float64(y3)) //these variable names are wrong im not fixing it
			if x == float64(int(x)) && y == float64(int(y)) {
				output += int(x)*3 + int(y)
			}
			x4, y4 := solveLinearEquations(float64(x1), float64(y1), 10000000000000+float64(x3), float64(x2), float64(y2), 10000000000000+float64(y3))
			if x4 == float64(int(x4)) && y4 == float64(int(y4)) {
				output1 += int(x4)*3 + int(y4)
			}
		}
	}

	fmt.Println(output)
	fmt.Println(output1)
}
