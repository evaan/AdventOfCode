package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	leftArr := make([]int, 1000)
	rightArr := make([]int, 1000)
	rightCount := make(map[int]int)

	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	for i, line := range strings.Split(string(file), "\n") {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		rightCount[rightNum] = rightCount[rightNum] + 1

		leftArr[i] = leftNum
		rightArr[i] = rightNum
	}

	slices.Sort(leftArr)
	slices.Sort(rightArr)

	output := 0
	output1 := 0

	for i := range 1000 {
		//part 1
		if leftArr[i] > rightArr[i] {
			output += leftArr[i] - rightArr[i]
		} else {
			output += rightArr[i] - leftArr[i]
		}
		output1 += leftArr[i] * rightCount[leftArr[i]] //part 2
	}

	fmt.Println("Part 1 Output:", output)
	fmt.Println("Part 2 Output:", output1)

}
