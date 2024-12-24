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
	cheats    map[coord]int
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
	part1 := m.findCheats(100)

	fmt.Println("part1", part1)
}

func (m maze) traverse() {
	y, x := m.findStart()
	score := 0

	isDotOrE := func(v record) bool {
		return (v.mazeValue == "." || v.mazeValue == "E") && v.score == 0
	}

	for {
		m[y][x].score = score

		score++
		switch {
		case isDotOrE(m[y][x-1]):
			x = x - 1
		case isDotOrE(m[y][x+1]):
			x = x + 1
		case isDotOrE(m[y-1][x]):
			y = y - 1
		case isDotOrE(m[y+1][x]):
			y = y + 1
		}

		if m[y][x].mazeValue == "E" {
			break
		}
	}

	m[y][x].score = score
}

func (m maze) findCheats(min int) int {

	checkCheats := func(s, e record) int {
		if e.mazeValue != "#" {
			return -1
		}
		diffY := e.c.y - s.c.y
		diffX := e.c.x - s.c.x

		if e.c.y+diffY < 0 || e.c.x+diffX < 0 || e.c.y+diffY >= len(m) || e.c.x+diffX >= len(m[0]) {
			return -1
		}

		landingRecord := m[e.c.y+diffY][e.c.x+diffX]
		if landingRecord.mazeValue != "." && landingRecord.mazeValue != "E" {
			return -1
		}
		saved := landingRecord.score - s.score - 2
		if saved <= 0 {
			return -1
		}
		s.cheats[e.c] = saved
		return saved
	}

	isNext := func(s, e record) bool {
		return (e.mazeValue == "." || e.mazeValue == "E") && s.score+1 == e.score
	}

	y, x := m.findStart()

	total := 0
	for {
		if checkCheats(m[y][x], m[y][x-1]) >= min {
			total++
		}
		if checkCheats(m[y][x], m[y][x+1]) >= min {
			total++
		}
		if checkCheats(m[y][x], m[y-1][x]) >= min {
			total++
		}
		if checkCheats(m[y][x], m[y+1][x]) >= min {
			total++
		}

		switch {
		case isNext(m[y][x], m[y][x-1]):
			x = x - 1
		case isNext(m[y][x], m[y][x+1]):
			x = x + 1
		case isNext(m[y][x], m[y-1][x]):
			y = y - 1
		case isNext(m[y][x], m[y+1][x]):
			y = y + 1
		}

		if m[y][x].mazeValue == "E" {
			break
		}
	}
	return total
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
