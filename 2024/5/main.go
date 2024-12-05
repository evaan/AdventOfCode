package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func fixInvalid(input []string, rules map[string][]string) []string {
	tmp := input
	output := make([]string, 0)
	for len(tmp) > 0 {
		for i, num := range tmp {
			valid := true
			for _, rule := range rules[num] {
				if slices.Contains(tmp[i+1:], rule) {
					valid = false
					break
				}
			}
			if valid {
				output = append(output, num)
				tmp = append(tmp[:i], tmp[i+1:]...)
				break
			}
		}
	}
	return output
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rules := make(map[string][]string)
	correct := make([][]string, 0)
	fixed := make([][]string, 0)

lineLoop:
	for _, line := range strings.Split(string(file), "\n") {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "|") {
			tmp := strings.Split(line, "|")
			rules[tmp[1]] = append(rules[tmp[1]], tmp[0])
		} else {
			nums := strings.Split(line, ",")
			for i, num := range nums {
				for _, one := range rules[num] {
					for _, two := range nums[i:] {
						if one == two {
							fixed = append(fixed, fixInvalid(nums, rules))
							continue lineLoop
						}
					}
				}
			}
			correct = append(correct, nums)
		}
	}

	output := 0
	output1 := 0

	for _, list := range correct {
		tmp, err := strconv.Atoi(list[len(list)/2])
		if err != nil {
			panic(err)
		}
		output += tmp
	}

	fmt.Println(output)

	for _, list := range fixed {
		tmp, err := strconv.Atoi(list[len(list)/2])
		if err != nil {
			panic(err)
		}
		output1 += tmp
	}

	fmt.Println(output1)
}
