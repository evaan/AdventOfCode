package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generateCombinations(n int) []string {
	combos := make([]string, 0)
	total := 1 << n

	for i := 0; i < total; i++ {
		combos = append(combos, fmt.Sprintf("%0*b", n, i))
	}

	return combos
}

func toBase3(num, length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result = strconv.Itoa(num%3) + result
		num /= 3
	}
	return result
}

func generateCombinations1(n int) []string {
	combos := make([]string, 0)
	total := 1
	for i := 0; i < n; i++ {
		total *= 3
	}

	for i := 0; i < total; i++ {
		combos = append(combos, fmt.Sprintf("%0*b", n, i))
	}

	// Convert binary-like format to base-3 strings
	for i := range combos {
		combos[i] = toBase3(i, n)
	}

	return combos
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	output := 0
	output1 := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		split := strings.Split(line, ": ")
		goal, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		segments := strings.Fields(split[1])
		n := len(segments)

		//part 1
		for _, combo := range generateCombinations(n - 1) {
			current, err := strconv.Atoi(segments[0])
			if err != nil {
				panic(err)
			}

			for i := range n - 1 {
				next, err := strconv.Atoi(segments[i+1])
				if err != nil {
					panic(err)
				}

				if combo[i] == '0' {
					current += next
				} else {
					current *= next
				}
			}

			if current == goal {
				output += goal
				break
			}
		}

		for _, combo := range generateCombinations1(n - 1) {
			current, err := strconv.Atoi(segments[0])
			if err != nil {
				panic(err)
			}

			for i := 0; i < n-1; i++ {
				next, err := strconv.Atoi(segments[i+1])
				if err != nil {
					panic(err)
				}

				switch combo[i] {
				case '0':
					current += next
				case '1':
					current *= next
				case '2':
					concatValue, err := strconv.Atoi(fmt.Sprintf("%d%d", current, next))
					if err != nil {
						panic(err)
					}
					current = concatValue
				}
			}

			if current == goal {
				output1 += goal
				break
			}
		}
	}

	fmt.Println(output)
	fmt.Println(output1)
}
