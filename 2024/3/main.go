package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	//48-57 == 0-9
	//44 = ,
	//41 = )

	output := 0
	output1 := 0

	i := 0
	do := true
	for i < len(string(file)) {
		if (i+3 < len(string(file))) && (string(file)[i:i+4]) == "do()" {
			do = true
		}
		if (i+7 < len(string(file))) && (string(file)[i:i+7]) == "don't()" {
			do = false
		}
		if (i+4 < len(string(file))) && (string(file)[i:i+4]) == "mul(" {
			i += 4
			hasComma := false
			num1 := "0"
			num2 := "0"
			for {
				if i > len(string(file)) {
					break
				}
				char := string(file)[i]
				if (char >= 48 && char <= 57) || (char == 44 && !hasComma) || char == 41 {
					if char == 44 {
						hasComma = true
					}
					if char >= 48 && char <= 57 {
						if hasComma {
							num2 += string(char)
						} else {
							num1 += string(char)
						}
					}
					if char == 41 {
						int1, err := strconv.Atoi(num1)
						if err != nil {
							panic(err)
						}
						int2, err := strconv.Atoi(num2)
						if err != nil {
							panic(err)
						}
						output += int1 * int2
						if do {
							output1 += int1 * int2
						}
						break
					}
					i++
				} else {
					break
				}
			}
		} else {
			i++
		}
	}

	fmt.Println("Part 1:", output)
	fmt.Println("Part 2:", output1)
}
