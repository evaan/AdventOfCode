package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	leftArr := []int{}
	rightArr := []int{}
	rightCount := make(map[int]int)

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		rightCount[rightNum] = rightCount[rightNum] + 1

		leftArr = append(leftArr, leftNum)
		rightArr = append(rightArr, rightNum)
	}

	slices.Sort(leftArr)
	slices.Sort(rightArr)

	output := 0
	output1 := 0

	for i := range 1000 {
		if leftArr[i] > rightArr[i] {
			output += leftArr[i] - rightArr[i]
		} else {
			output += rightArr[i] - leftArr[i]
		}
		output1 += leftArr[i] * rightCount[leftArr[i]]
	}

	fmt.Println(output)
	fmt.Println(output1)

}
