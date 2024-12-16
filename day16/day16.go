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

type maze [][]string

type coord struct {
	x, y int
}

type path map[coord]struct{}

func (p path) contains(x, y int) bool {
	_, ok := p[coord{x, y}]
	return ok
}

func (p path) copyWith(x, y int) path {
	n := make(path)
	for k, v := range p {
		n[k] = v
	}

	n[coord{x, y}] = struct{}{}

	return n
}

func main() {
	splitinput := strings.Split(inputtest, "\n")

	m := make(maze, len(splitinput))
	for r, line := range splitinput {
		m[r] = strings.Split(line, "")
	}

	m.printMaze()

	x, y := m.findStart()
	p := make(path)
	score := 0

	part1 := m.traverse(1, 0, x, y, p, score)
	fmt.Println("part1", part1)
}

func (m maze) traverse(xMove, yMove, x, y int, p path, score int) int {
	if p.contains(x, y) {
		return math.MaxInt
	}
	p = p.copyWith(x, y)

	switch m[y][x] {
	case "#":
		return math.MaxInt
	case "E":
		return score
	case "S", ".":
		straight := m.traverse(xMove, yMove, x+xMove, y+yMove, p, score+1)
		turn1 := math.MaxInt
		turn2 := math.MaxInt
		if yMove == 0 {
			turn1 = m.traverse(0, 1, x, y+1, p, score+1001)
			turn2 = m.traverse(0, -1, x, y-1, p, score+1001)
		} else if xMove == 0 {
			turn1 = m.traverse(1, 0, x+1, y, p, score+1001)
			turn2 = m.traverse(-1, 0, x-1, y, p, score+1001)
		}
		return min(straight, turn1, turn2)
	}
	panic("should not reach here")
}

func (m maze) findStart() (int, int) {
	for y, _ := range m {
		for x, _ := range m[y] {
			if m[y][x] == "S" {
				return x, y
			}
		}
	}
	panic("no start")
}

func (m maze) printMaze() {
	for _, row := range m {
		for _, c := range row {
			fmt.Print(c)
		}
		fmt.Println()
	}
}
