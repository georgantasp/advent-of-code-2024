package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input
var input string

//go:embed inputtest
var inputtest string

type maze [][]record

type record struct {
	mazeValue  string
	bestScores map[coord]int
}

type coord struct {
	x, y int
}

func main() {
	splitinput := strings.Split(input, "\n")

	m := make(maze, len(splitinput))
	for r, line := range splitinput {
		m[r] = make([]record, len(line))
		for c, v := range strings.Split(line, "") {
			m[r][c] = record{mazeValue: v, bestScores: make(map[coord]int)}
		}
	}

	m.printMaze()

	x, y := m.findStart()
	score := 0

	part1 := m.traverse(1, 0, x, y, score)
	fmt.Println("part1", part1)
}

func (m maze) traverse(xMove, yMove, x, y int, score int) int {
	if s, ok := m[y][x].bestScores[coord{xMove, yMove}]; ok && s <= score {
		return math.MaxInt
	}
	m[y][x].bestScores[coord{xMove, yMove}] = score

	switch m[y][x].mazeValue {
	case "#":
		return math.MaxInt
	case "E":
		return score
	case "S", ".":
		fmt.Println(x, y)
		straight := m.traverse(xMove, yMove, x+xMove, y+yMove, score+1)
		turn1 := math.MaxInt
		turn2 := math.MaxInt
		if yMove == 0 {
			turn1 = m.traverse(0, 1, x, y+1, score+1001)
			turn2 = m.traverse(0, -1, x, y-1, score+1001)
		} else if xMove == 0 {
			turn1 = m.traverse(1, 0, x+1, y, score+1001)
			turn2 = m.traverse(-1, 0, x-1, y, score+1001)
		}
		return min(straight, turn1, turn2)
	}
	panic("should not reach here")
}

func (m maze) findStart() (int, int) {
	for y, _ := range m {
		for x, _ := range m[y] {
			if m[y][x].mazeValue == "S" {
				return x, y
			}
		}
	}
	panic("no start")
}

func (m maze) printMaze() {
	for _, row := range m {
		for _, c := range row {
			fmt.Print(c.mazeValue)
		}
		fmt.Println()
	}
}
