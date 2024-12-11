package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	stones := make(map[string]int, 0)
	step25stones := make(map[string]int, 0)

	for _, stone := range strings.Split(strings.Split(string(file), "\n")[0], " ") {
		stones[stone]++
	}

	for i := range 75 {
		if i == 25 {
			step25stones = stones
		}
		newStones := make(map[string]int, 0)
		for stone, stoneCount := range stones {
			if stone == "0" {
				newStones["1"] += stoneCount
			} else if len(stone)%2 == 0 {
				left := stone[:len(stone)/2]
				right := stone[len(stone)/2:]
				leftInt, err := strconv.Atoi(left)
				if err != nil {
					panic(err)
				}
				rightInt, err := strconv.Atoi(right)
				if err != nil {
					panic(err)
				}
				newStones[strconv.Itoa(leftInt)] += stoneCount
				newStones[strconv.Itoa(rightInt)] += stoneCount
			} else {
				stoneInt, err := strconv.Atoi(stone)
				if err != nil {
					panic(err)
				}
				newStones[strconv.Itoa(stoneInt*2024)] += stoneCount
			}
		}
		stones = newStones
	}

	count := 0
	step25count := 0

	for _, stoneCount := range stones {
		count += stoneCount
	}

	for _, stoneCount := range step25stones {
		step25count += stoneCount
	}

	fmt.Println("25 Steps:", step25count)
	fmt.Println("75 Steps:", count)
}
