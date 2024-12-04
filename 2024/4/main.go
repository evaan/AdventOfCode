package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	xmasCount := 0
	crossMasCount := 0

	for i, line := range lines {
		for j, char := range line {
			if char == 'A' {
				if i-1 >= 0 && i+1 < len(lines) && j-1 >= 0 && j+1 < len(line) {
					if lines[i-1][j-1] == 'M' && lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' && lines[i+1][j+1] == 'S' {
						crossMasCount++
					}
					if lines[i-1][j-1] == 'S' && lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' && lines[i+1][j+1] == 'M' {
						crossMasCount++
					}
					if lines[i-1][j-1] == 'M' && lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' && lines[i+1][j+1] == 'S' {
						crossMasCount++
					}
					if lines[i-1][j-1] == 'S' && lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' && lines[i+1][j+1] == 'M' {
						crossMasCount++
					}
				}
			}
			if char == 'X' {
				if (j+4) <= len(line) && line[j:j+4] == "XMAS" {
					xmasCount++
				}
				if (j-3) >= 0 && line[j-3:j+1] == "SAMX" {
					xmasCount++
				}
				if i >= 3 {
					if j >= 3 {
						if lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
							xmasCount++
						}
					}
					if j+3 < len(line) {
						if lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
							xmasCount++
						}
					}
					if lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
						xmasCount++
					}
				}
				if i+3 < len(lines) {
					if j >= 3 {
						if lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
							xmasCount++
						}
					}
					if j+3 < len(line) {
						if lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
							xmasCount++
						}
					}
					if lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
						xmasCount++
					}
				}
			}
		}
	}

	fmt.Println(xmasCount)
	fmt.Println(crossMasCount)
}
