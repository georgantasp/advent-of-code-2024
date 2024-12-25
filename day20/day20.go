package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputtest
var inputtest string

//go:embed input
var input string

type maze [][]record

type record struct {
	c         coord
	mazeValue string
	score     int
	next      coord
	cheats    map[coord]int
}

func (r record) isUnscored() bool {
	return (r.mazeValue == "." || r.mazeValue == "E") && r.score == 0
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
			m[r][c] = record{mazeValue: v, c: coord{y: r, x: c}, cheats: map[coord]int{}}
		}
	}

	m.printMaze()
	m.traverse()
	part1 := m.findCheats(2, 100)
	fmt.Println("part1", part1)
	part2 := m.findCheats(20, 100)
	fmt.Println("part1", part2)
}

func (m maze) traverse() {
	y, x := m.findStart()
	score := 0

	for {
		m[y][x].score = score

		score++
		switch {
		case m[y][x-1].isUnscored():
			m[y][x].next = coord{x: x - 1, y: y}
			x = x - 1
		case m[y][x+1].isUnscored():
			m[y][x].next = coord{x: x + 1, y: y}
			x = x + 1
		case m[y-1][x].isUnscored():
			m[y][x].next = coord{x: x, y: y - 1}
			y = y - 1
		case m[y+1][x].isUnscored():
			m[y][x].next = coord{x: x, y: y + 1}
			y = y + 1
		}

		if m[y][x].mazeValue == "E" {
			break
		}
	}

	m[y][x].score = score
}

func (m maze) findCheats(cheatLen, minSave int) int {

	y, x := m.findStart()

	total := 0
	for {
		m.findCheatsFor(y, x, cheatLen)

		for _, v := range m[y][x].cheats {
			if v >= minSave {
				total++
			}
		}

		n := m[y][x].next
		y, x = n.y, n.x

		if m[y][x].mazeValue == "E" {
			break
		}
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (m maze) findCheatsFor(y, x, cheatLenMax int) {
	for yCheat := y - cheatLenMax; yCheat <= y+cheatLenMax; yCheat++ {
		if yCheat < 0 || yCheat >= len(m) {
			continue
		}
		for xCheat := x - cheatLenMax; xCheat <= x+cheatLenMax; xCheat++ {
			if xCheat < 0 || xCheat >= len(m[0]) {
				continue
			}

			cheatLen := abs(y-yCheat) + abs(x-xCheat)
			if cheatLen > cheatLenMax {
				continue
			}

			if m[yCheat][xCheat].mazeValue != "." && m[yCheat][xCheat].mazeValue != "E" {
				continue
			}

			saved := m[yCheat][xCheat].score - m[y][x].score - cheatLen
			if saved > 0 {
				m[y][x].cheats[coord{y: yCheat, x: xCheat}] = saved
			}
		}
	}
}

func (m maze) findStart() (int, int) {
	for y, _ := range m {
		for x, _ := range m[y] {
			if m[y][x].mazeValue == "S" {
				return y, x
			}
		}
	}
	panic("no start")
}

func (m maze) printMaze() int {
	bestCount := 0
	for _, row := range m {
		for _, c := range row {
			fmt.Print(c.mazeValue)
		}
		fmt.Println()
	}
	return bestCount
}
