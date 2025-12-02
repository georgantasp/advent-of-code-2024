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
	isBest     map[int]struct{}
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
			m[r][c] = record{mazeValue: v, bestScores: make(map[coord]int), isBest: make(map[int]struct{})}
		}
	}

	m.printMaze(0)

	x, y := m.findStart()
	score := 0

	part1, _ := m.traverse(1, 0, x, y, score)
	fmt.Println("part1", part1)
	best := m.printMaze(part1)
	fmt.Println("part2", best)
}

func (m maze) traverse(xMove, yMove, x, y int, score int) (int, bool) {
	if s, ok := m[y][x].bestScores[coord{xMove, yMove}]; ok && score > s {
		return math.MaxInt, false
	}
	m[y][x].bestScores[coord{xMove, yMove}] = score

	switch m[y][x].mazeValue {
	case "#":
		return math.MaxInt, false
	case "E":
		m[y][x].isBest[score] = struct{}{}
		return score, true
	case "S", ".":
		straight, endS := m.traverse(xMove, yMove, x+xMove, y+yMove, score+1)
		if yMove == 0 {
			turn1, end1 := m.traverse(0, 1, x, y+1, score+1001)
			turn2, end2 := m.traverse(0, -1, x, y-1, score+1001)

			minScore := min(straight, turn1, turn2)
			if endS || end1 || end2 {
				m[y][x].isBest[minScore] = struct{}{}
			}

			return minScore, endS || end1 || end2
		} else if xMove == 0 {
			turn1, end1 := m.traverse(1, 0, x+1, y, score+1001)
			turn2, end2 := m.traverse(-1, 0, x-1, y, score+1001)

			minScore := min(straight, turn1, turn2)
			if endS || end1 || end2 {
				m[y][x].isBest[minScore] = struct{}{}
			}

			return minScore, endS || end1 || end2
		}
		panic("should not reach here")
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

func (m maze) printMaze(best int) int {
	bestCount := 0
	for _, row := range m {
		for _, c := range row {
			if _, ok := c.isBest[best]; ok {
				fmt.Print("O")
				bestCount++
			} else {
				fmt.Print(c.mazeValue)
			}
		}
		fmt.Println()
	}
	return bestCount
}
