package main

import (
	"strconv"
	"strings"
)

type Block struct {
	data int
	size int
}

func part2(input string) int {
	blocks := make([]Block, 0)

	for i, char := range strings.Split(input, "\n")[0] {
		count, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			blocks = append(blocks, Block{i / 2, count})
		} else {
			blocks = append(blocks, Block{-1, count})
		}
	}

	for {
		moved := false

		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i].data != -1 {
				for j, block := range blocks {
					if j == i {
						break
					}
					if block.data == -1 && block.size >= blocks[i].size {
						blocks[j].data = blocks[i].data
						blocks[j].size = blocks[i].size
						blocks[i].data = -1
						remainingSize := block.size - blocks[i].size
						if remainingSize > 0 {
							extraDotBlock := Block{-1, remainingSize}
							blocks = append(blocks[:j+1], append([]Block{extraDotBlock}, blocks[j+1:]...)...)
						}
						moved = true
						break
					}
				}
			}
		}

		if !moved {
			break
		}
	}

	checksum := 0
	blockArr := make([]int, 0)

	for _, block := range blocks {
		for range block.size {
			blockArr = append(blockArr, block.data)
		}
	}

	for i, num := range blockArr {
		if num != -1 {
			checksum += i * num
		}
	}

	return checksum
}
