package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	submarine := strings.Split(string(file), "\n\n")[0]
	instructions := strings.Split(string(file), "\n\n")[1]

	pos := Point{0, 0}
	walls := make([]Point, 0)
	boxes := make([]Point, 0)

	for y, line := range strings.Split(submarine, "\n") {
		for x, char := range line {
			switch char {
			case '#':
				walls = append(walls, Point{x, y})
			case 'O':
				boxes = append(boxes, Point{x, y})
			case '@':
				pos = Point{x, y}
			}
		}
	}

	for _, instruction := range instructions {
		switch instruction {
		case '<':
			if !slices.Contains(walls, Point{pos.x - 1, pos.y}) && !slices.Contains(boxes, Point{pos.x - 1, pos.y}) {
				pos.x--
				break
			}
			if !slices.Contains(walls, Point{pos.x - 1, pos.y}) {
				if slices.Contains(boxes, Point{pos.x - 1, pos.y}) {
				checkLoop1:
					for x := pos.x - 1; x > 0; x-- {
						if slices.Contains(walls, Point{x, pos.y}) {
							break checkLoop1
						}
						if !slices.Contains(boxes, Point{x, pos.y}) {
							for i, box := range boxes {
								if box.x == pos.x-1 && box.y == pos.y {
									boxes[i] = Point{x, pos.y}
									pos.x--
									break checkLoop1
								}
							}
						}
					}
				}
			}
		case '>':
			if !slices.Contains(walls, Point{pos.x + 1, pos.y}) && !slices.Contains(boxes, Point{pos.x + 1, pos.y}) {
				pos.x++
				break
			}
			if !slices.Contains(walls, Point{pos.x + 1, pos.y}) {
				if slices.Contains(boxes, Point{pos.x + 1, pos.y}) {
				checkLoop2:
					for x := pos.x + 1; x < len(strings.Split(submarine, "\n")[0]); x++ {
						if slices.Contains(walls, Point{x, pos.y}) {
							break checkLoop2
						}
						if !slices.Contains(boxes, Point{x, pos.y}) {
							for i, box := range boxes {
								if box.x == pos.x+1 && box.y == pos.y {
									boxes[i] = Point{x, pos.y}
									pos.x++
									break checkLoop2
								}
							}
						}
					}
				}
			}
		case '^':
			if !slices.Contains(boxes, Point{pos.x, pos.y - 1}) && !slices.Contains(walls, Point{pos.x, pos.y - 1}) {
				pos.y--
				break
			}
			if !slices.Contains(walls, Point{pos.x, pos.y - 1}) {
				if slices.Contains(boxes, Point{pos.x, pos.y - 1}) {
				checkLoop3:
					for y := pos.y - 1; y > 0; y-- {
						if slices.Contains(walls, Point{pos.x, y}) {
							break checkLoop3
						}
						if !slices.Contains(boxes, Point{pos.x, y}) {
							for i, box := range boxes {
								if box.x == pos.x && box.y == pos.y-1 {
									boxes[i] = Point{pos.x, y}
									pos.y--
									break checkLoop3
								}
							}
						}
					}
				}
			}
		case 'v':
			if !slices.Contains(boxes, Point{pos.x, pos.y + 1}) && !slices.Contains(walls, Point{pos.x, pos.y + 1}) {
				pos.y++
				break
			}
			if !slices.Contains(walls, Point{pos.x, pos.y + 1}) {
				if slices.Contains(boxes, Point{pos.x, pos.y + 1}) {
				checkLoop4:
					for y := pos.y + 1; y < len(strings.Split(submarine, "\n")); y++ {
						if slices.Contains(walls, Point{pos.x, y}) {
							break checkLoop4
						}
						if !slices.Contains(boxes, Point{pos.x, y}) {
							for i, box := range boxes {
								if box.x == pos.x && box.y == pos.y+1 {
									boxes[i] = Point{pos.x, y}
									pos.y++
									break checkLoop4
								}
							}
						}
					}
				}
			}
		}
	}

	output := 0

	for _, box := range boxes {
		output += box.y*100 + box.x
	}

	fmt.Println(output)
}
