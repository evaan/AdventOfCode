package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(string(file)))
	fmt.Println(part2(string(file)))
}
