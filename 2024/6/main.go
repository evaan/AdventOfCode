package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pos struct {
	x, y int
}

type TurnState struct {
	pos       Pos
	direction int
}

func simulate(initialPos Pos, lines []string, barrier *Pos) int {
	xBound := len(lines[0])
	yBound := len(lines)
	pos := initialPos
	visitedPos := make([]Pos, 0)
	turnStates := make([]TurnState, 0)
	direction := 0 //0 - up, 1 - right, 2 - down, 3 - left

	distinctTiles := 1

mainLoop:
	for {
		switch direction % 4 {
		case 0:
			if pos.y-1 < 0 {
				break mainLoop
			} else if lines[pos.y-1][pos.x] == '#' || (barrier != nil && (pos.y-1 == barrier.y && pos.x == barrier.x)) {
				if !slices.Contains(turnStates, TurnState{pos, direction % 4}) {
					turnStates = append(turnStates, TurnState{pos, direction % 4})
				} else {
					return -1
				}
				direction++
			} else {
				if !slices.Contains(visitedPos, Pos{pos.x, pos.y}) {
					visitedPos = append(visitedPos, pos)
					distinctTiles++
				}
				pos.y--
			}
		case 1:
			if pos.x+1 >= xBound {
				break mainLoop
			} else if lines[pos.y][pos.x+1] == '#' || (barrier != nil && (pos.y == barrier.y && pos.x+1 == barrier.x)) {
				if !slices.Contains(turnStates, TurnState{pos, direction % 4}) {
					turnStates = append(turnStates, TurnState{pos, direction % 4})
				} else {
					return -1
				}
				direction++
			} else {
				if !slices.Contains(visitedPos, Pos{pos.x, pos.y}) {
					visitedPos = append(visitedPos, pos)
					distinctTiles++
				}
				pos.x++
			}
		case 2:
			if pos.y+1 >= yBound {
				break mainLoop
			} else if lines[pos.y+1][pos.x] == '#' || (barrier != nil && (pos.y+1 == barrier.y && pos.x == barrier.x)) {
				if !slices.Contains(turnStates, TurnState{pos, direction % 4}) {
					turnStates = append(turnStates, TurnState{pos, direction % 4})
				} else {
					return -1
				}
				direction++
			} else {
				if !slices.Contains(visitedPos, Pos{pos.x, pos.y}) {
					visitedPos = append(visitedPos, pos)
					distinctTiles++
				}
				pos.y++
			}
		case 3:
			if pos.x-1 < 0 {
				break mainLoop
			} else if lines[pos.y][pos.x-1] == '#' || (barrier != nil && (pos.y == barrier.y && pos.x-1 == barrier.x)) {
				if !slices.Contains(turnStates, TurnState{pos, direction % 4}) {
					turnStates = append(turnStates, TurnState{pos, direction % 4})
				} else {
					return -1
				}
				direction++
			} else {
				if !slices.Contains(visitedPos, Pos{pos.x, pos.y}) {
					visitedPos = append(visitedPos, pos)
					distinctTiles++
				}
				pos.x--
			}
		}
	}

	return distinctTiles
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	pos := Pos{0, 0}

	for y, line := range lines {
		for x, line := range line {
			if line == '^' {
				pos = Pos{x, y}
			}
		}
	}

	possibleTraps := 0

	for x, line := range lines {
		for y := range line {
			if simulate(pos, lines, &Pos{x, y}) == -1 {
				possibleTraps++
			}
		}
	}

	fmt.Println("Possible Traps:", possibleTraps)
	fmt.Println("Distinct Tiles:", simulate(pos, lines, nil))
}
