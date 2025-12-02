package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input
var input string

//go:embed inputtest
var inputtest string

type maze [][]record

type record struct {
	mazeValue string
	bestScore int
}

type coord struct {
	x, y int
}

func main() {
	size := 70
	inputLines := strings.Split(input, "\n")
	start := 1024

	bytes := make([]coord, len(inputLines))
	for i, line := range inputLines {
		s := strings.Split(line, ",")
		bytes[i] = coord{}
		bytes[i].x, _ = strconv.Atoi(s[0])
		bytes[i].y, _ = strconv.Atoi(s[1])
	}

	m := make(maze, size+1)
	for i := range m {
		m[i] = make([]record, size+1)
		for j := range m[i] {
			val := "."
			if i == 0 && j == 0 {
				val = "S"
			}
			if i == size && j == size {
				val = "E"
			}
			m[i][j].mazeValue = val
			m[i][j].bestScore = math.MaxInt
		}
	}

	for b := range start {
		m[bytes[b].y][bytes[b].x].mazeValue = "#"
	}

	m.print()
	score := m.traverse(1, 0, 0, 0, 0)
	fmt.Println("part1", score)

	for b := start; b < len(bytes); b++ {
		m[bytes[b].y][bytes[b].x].mazeValue = "#"
		m.clearScores()
		score = m.traverse(1, 0, 0, 0, 0)
		if score == math.MaxInt {
			fmt.Println("part2", bytes[b])
			break
		}
	}
}

func (m maze) traverse(xMove, yMove, x, y int, score int) int {
	if x < 0 || y < 0 || x >= len(m[0]) || y >= len(m) {
		return math.MaxInt
	}
	if score >= m[y][x].bestScore {
		return math.MaxInt
	}
	m[y][x].bestScore = score

	switch m[y][x].mazeValue {
	case "#":
		return math.MaxInt
	case "E":
		return score
	case "S", ".":
		straight := m.traverse(xMove, yMove, x+xMove, y+yMove, score+1)
		if yMove == 0 {
			turn1 := m.traverse(0, 1, x, y+1, score+1)
			turn2 := m.traverse(0, -1, x, y-1, score+1)

			minScore := min(straight, turn1, turn2)

			return minScore
		} else if xMove == 0 {
			turn1 := m.traverse(1, 0, x+1, y, score+1)
			turn2 := m.traverse(-1, 0, x-1, y, score+1)

			minScore := min(straight, turn1, turn2)

			return minScore
		}
		panic("should not reach here")
	}
	panic("should not reach here")
}

func (m maze) print() {
	for _, row := range m {
		for _, r := range row {
			fmt.Print(r.mazeValue)
		}
		fmt.Println()
	}
}

func (m maze) clearScores() {
	for i := range len(m) {
		for j := range len(m[i]) {
			m[j][i].bestScore = math.MaxInt
		}
	}
}
