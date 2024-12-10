package main

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	blocks := make([]int, 0)

	for i, char := range strings.Split(input, "\n")[0] {
		count, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			for range count {
				blocks = append(blocks, i/2)
			}
		} else {
			for range count {
				blocks = append(blocks, -1)
			}
		}
	}

outerLoop:
	for {
		firstDotPos := -1
		for i, num := range blocks {
			if firstDotPos == -1 && num == -1 {
				firstDotPos = i
				break
			}
		}

		if firstDotPos == -1 {
			break outerLoop
		}

		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i] != -1 {
				blocks[firstDotPos], blocks[i] = blocks[i], -1
				break
			}
			if i == firstDotPos {
				break outerLoop
			}
		}
	}

	checksum := 0

	for i, num := range blocks {
		if num != -1 {
			checksum += i * num
		}
	}

	return checksum
}
