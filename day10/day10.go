package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

type puzzMap [][]coord
type coord struct {
	row    int
	col    int
	height int
}
type trail [10]coord

func (puzz puzzMap) findTrailHeads() []coord {
	var trailHeads []coord
	for r := range len(puzz) {
		for j := range puzz[r] {
			if puzz[r][j].height == 0 {
				trailHeads = append(trailHeads, puzz[r][j])
			}
		}
	}
	return trailHeads
}

func (puzz puzzMap) findTopsFrom(c coord) map[coord]struct{} {
	tops := make(map[coord]struct{})

	if c.height == 9 {
		tops[c] = struct{}{}
		return tops
	}

	if c.row+1 < len(puzz) && puzz[c.row+1][c.col].height == c.height+1 {
		for k, v := range puzz.findTopsFrom(puzz[c.row+1][c.col]) {
			tops[k] = v
		}
	}
	if c.row-1 >= 0 && puzz[c.row-1][c.col].height == c.height+1 {
		for k, v := range puzz.findTopsFrom(puzz[c.row-1][c.col]) {
			tops[k] = v
		}
	}
	if c.col+1 < len(puzz[0]) && puzz[c.row][c.col+1].height == c.height+1 {
		for k, v := range puzz.findTopsFrom(puzz[c.row][c.col+1]) {
			tops[k] = v
		}
	}
	if c.col-1 >= 0 && puzz[c.row][c.col-1].height == c.height+1 {
		for k, v := range puzz.findTopsFrom(puzz[c.row][c.col-1]) {
			tops[k] = v
		}
	}

	return tops
}

func (puzz puzzMap) findTrailsFrom(c coord, t trail) map[trail]struct{} {
	trails := make(map[trail]struct{})

	t[c.height] = c

	if c.height == 9 {
		trails[t] = struct{}{}
		return trails
	}

	if c.row+1 < len(puzz) && puzz[c.row+1][c.col].height == c.height+1 {
		for k, v := range puzz.findTrailsFrom(puzz[c.row+1][c.col], t) {
			trails[k] = v
		}
	}
	if c.row-1 >= 0 && puzz[c.row-1][c.col].height == c.height+1 {
		for k, v := range puzz.findTrailsFrom(puzz[c.row-1][c.col], t) {
			trails[k] = v
		}
	}
	if c.col+1 < len(puzz[0]) && puzz[c.row][c.col+1].height == c.height+1 {
		for k, v := range puzz.findTrailsFrom(puzz[c.row][c.col+1], t) {
			trails[k] = v
		}
	}
	if c.col-1 >= 0 && puzz[c.row][c.col-1].height == c.height+1 {
		for k, v := range puzz.findTrailsFrom(puzz[c.row][c.col-1], t) {
			trails[k] = v
		}
	}

	return trails
}

func main() {
	lines := strings.Split(input, "\n")

	puzz := make([][]coord, len(lines))

	for row, line := range lines {
		split := strings.Split(line, "")
		puzz[row] = make([]coord, len(split))
		for col, num := range split {
			height, _ := strconv.Atoi(num)
			puzz[row][col] = coord{row: row, col: col, height: height}
		}
	}

	part1(puzz)
	part2(puzz)
}

func part1(p puzzMap) {
	trailHeads := p.findTrailHeads()
	totalScore := 0
	for _, t := range trailHeads {
		tops := p.findTopsFrom(t)
		totalScore += len(tops)
	}

	fmt.Println("part1", totalScore)
}

func part2(p puzzMap) {
	trailHeads := p.findTrailHeads()
	totalScore := 0
	for _, c := range trailHeads {
		trails := p.findTrailsFrom(c, trail{})
		totalScore += len(trails)
	}

	fmt.Println("part2", totalScore)
}
