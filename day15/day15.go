package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputsmall
var inputsmall string

//go:embed inputcustom
var inputcustom string

//go:embed inputtest
var inputtest string

//go:embed input
var input string

type puzzle [][]string
type robot struct {
	x int
	y int
}

func main() {
	splitinput := strings.Split(input, "\n\n")

	moveLines := strings.Split(splitinput[1], "\n")
	moves := []string{}
	for _, line := range moveLines {
		moves = append(moves, strings.Split(line, "")...)
	}

	puzzleLines := strings.Split(splitinput[0], "\n")
	p := make(puzzle, len(puzzleLines))
	for i, line := range puzzleLines {
		p[i] = strings.Split(line, "")
	}

	part1(p, moves)
	part2(p, moves)
}

func part1(p puzzle, moves []string) {
	p = p.copy()
	r := p.findRobot()
	for _, m := range moves {
		xMove := 0
		yMove := 0
		switch m {
		case "^":
			yMove = -1
		case "v":
			yMove = 1
		case ">":
			xMove = 1
		case "<":
			xMove = -1
		}

		moved := p.move(xMove, yMove, r.x, r.y)
		if moved {
			r.x += xMove
			r.y += yMove
		}
	}

	total := 0
	for y, line := range p {
		for x, v := range line {
			if v == "O" {
				total += (100 * y) + x
			}
		}
	}
	fmt.Println("part1", total)
}

func part2(p puzzle, moves []string) {
	p = p.scaleUp()

	r := p.findRobot()
	for _, m := range moves {
		var moved bool
		xMove := 0
		yMove := 0

		switch m {
		case "^":
			yMove = -1
			if p.canMoveY(yMove, r.x, r.y) {
				p.moveY(yMove, r.x, r.y)
				moved = true
			}
		case "v":
			yMove = 1
			if p.canMoveY(yMove, r.x, r.y) {
				p.moveY(yMove, r.x, r.y)
				moved = true
			}
		case ">":
			xMove = 1
			moved = p.move(xMove, yMove, r.x, r.y)
		case "<":
			xMove = -1
			moved = p.move(xMove, yMove, r.x, r.y)
		}

		if moved {
			r.x += xMove
			r.y += yMove
		}
	}

	total := 0
	for y, row := range p {
		for x, v := range row {
			if v == "[" {
				total += (100 * y) + x
			}
		}
	}
	fmt.Println("part2", total)
}

func (p puzzle) printPuzzle() {
	for _, row := range p {
		for _, c := range row {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func (p puzzle) findRobot() robot {
	for y, _ := range p {
		for x, _ := range p[y] {
			if p[y][x] == "@" {
				return robot{x: x, y: y}
			}
		}
	}
	panic("no robot")
}

func (p puzzle) move(xMove, yMove, x, y int) bool {
	if p[y+yMove][x+xMove] == "#" {
		return false
	}
	if p[y+yMove][x+xMove] == "." || p.move(xMove, yMove, x+xMove, y+yMove) {
		tmp := p[y+yMove][x+xMove]
		p[y+yMove][x+xMove] = p[y][x]
		p[y][x] = tmp
		return true
	}
	return false
}

func (p puzzle) copy() puzzle {
	n := make(puzzle, len(p))
	for y := range len(p) {
		n[y] = make([]string, len(p[y]))
		for x := range len(p[y]) {
			n[y][x] = p[y][x]
		}
	}
	return n
}

func (p puzzle) scaleUp() puzzle {
	n := make(puzzle, len(p))
	for y := range len(p) {
		n[y] = make([]string, len(p[y])*2)
		for x := range len(p[y]) {
			switch p[y][x] {
			case "#":
				n[y][2*x] = "#"
				n[y][2*x+1] = "#"
			case ".":
				n[y][2*x] = "."
				n[y][2*x+1] = "."
			case "O":
				n[y][2*x] = "["
				n[y][2*x+1] = "]"
			case "@":
				n[y][2*x] = "@"
				n[y][2*x+1] = "."
			}
		}
	}
	return n
}

func (p puzzle) canMoveY(yMove, x, y int) bool {
	if p[y+yMove][x] == "#" {
		return false
	}
	if p[y+yMove][x] == "." {
		return true
	}
	if p[y+yMove][x] == "[" {
		return p.canMoveY(yMove, x, y+yMove) && p.canMoveY(yMove, x+1, y+yMove)
	}
	if p[y+yMove][x] == "]" {
		return p.canMoveY(yMove, x-1, y+yMove) && p.canMoveY(yMove, x, y+yMove)
	}
	return false
}

func (p puzzle) moveY(yMove, x, y int) {
	if p[y+yMove][x] == "[" {
		p.moveY(yMove, x, y+yMove)
		p.moveY(yMove, x+1, y+yMove)
	}
	if p[y+yMove][x] == "]" {
		p.moveY(yMove, x-1, y+yMove)
		p.moveY(yMove, x, y+yMove)
	}
	tmp := p[y+yMove][x]
	p[y+yMove][x] = p[y][x]
	p[y][x] = tmp
}
