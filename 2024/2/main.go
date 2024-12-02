package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(input int) int {
	if input < 0 {
		return -input
	} else {
		return input
	}
}

func strSliceToIntSlice(input []string) []int {
	intArr := make([]int, len(input))

	for i, str := range input {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intArr[i] = num
	}

	return intArr
}

func removeNthItem(slice []int, n int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return append(newSlice[:n], newSlice[n+1:]...)
}

func check(levels []int) bool {
	decreasing := false
	increasing := false
	for i, level := range levels {
		if i == 0 {
			continue
		}
		if abs(level-levels[i-1]) > 3 || abs(level-levels[i-1]) < 1 {
			return false
		}
		if level-levels[i-1] < 0 {
			if !increasing {
				decreasing = true
			} else {
				return false
			}
		}
		if level-levels[i-1] > 0 {
			if !decreasing {
				increasing = true
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	safeReports := 0
	safeishReports := 0

	for _, line := range strings.Split(string(file), "\n") {
		levels := strSliceToIntSlice(strings.Fields(line))
		if len(levels) == 0 {
			continue
		}
		safeIters := 0
		for i := range len(levels) {
			if check(removeNthItem(levels, i)) {
				safeIters++
			}
		}
		if check(levels) {
			safeReports++
		}
		if safeIters > 0 {
			safeishReports++
		}
	}
	fmt.Println("Safe Reports:", safeReports)
	fmt.Println("Safe-ish Reports:", safeishReports)
}
