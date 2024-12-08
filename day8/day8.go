package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	part1()
	part2()
}

func part1() {
	lines := strings.Split(input, "\n")

	maxRow := len(lines)
	maxCol := len(lines[0])
	puzzMap := make(map[string][]nodeCoord)

	unique := make(map[nodeCoord]struct{})
	for i, line := range lines {
		for j, node := range strings.Split(line, "") {
			if node == "." {
				continue
			}
			coord := nodeCoord{row: i, col: j}
			if _, ok := puzzMap[node]; ok {
				for _, otherCoord := range puzzMap[node] {
					rowDist := coord.row - otherCoord.row
					colDist := coord.col - otherCoord.col

					testCoord1 := nodeCoord{row: coord.row + rowDist, col: coord.col + colDist}
					testCoord2 := nodeCoord{row: otherCoord.row - rowDist, col: otherCoord.col - colDist}
					if testCoord1.Valid(maxRow, maxCol) {
						unique[testCoord1] = struct{}{}
					}
					if testCoord2.Valid(maxRow, maxCol) {
						unique[testCoord2] = struct{}{}
					}
				}
			}
			puzzMap[node] = append(puzzMap[node], coord)
		}
	}

	fmt.Println("part1", len(unique))
}

func part2() {
	lines := strings.Split(input, "\n")

	maxRow := len(lines)
	maxCol := len(lines[0])
	puzzMap := make(map[string][]nodeCoord)

	unique := make(map[nodeCoord]struct{})
	for i, line := range lines {
		for j, node := range strings.Split(line, "") {
			if node == "." {
				continue
			}
			coord := nodeCoord{row: i, col: j}
			if _, ok := puzzMap[node]; ok {
				for _, otherCoord := range puzzMap[node] {
					unique[coord] = struct{}{}
					unique[otherCoord] = struct{}{}

					rowDist := coord.row - otherCoord.row
					colDist := coord.col - otherCoord.col

					testCoord1 := coord
					for {
						testCoord1 = nodeCoord{row: testCoord1.row + rowDist, col: testCoord1.col + colDist}
						if testCoord1.Valid(maxRow, maxCol) {
							unique[testCoord1] = struct{}{}
						} else {
							break
						}
					}

					testCoord2 := otherCoord
					for {
						testCoord2 = nodeCoord{row: testCoord2.row - rowDist, col: testCoord2.col - colDist}
						if testCoord2.Valid(maxRow, maxCol) {
							unique[testCoord2] = struct{}{}
						} else {
							break
						}
					}
				}
			}
			puzzMap[node] = append(puzzMap[node], coord)
		}
	}

	fmt.Println("part2", len(unique))
}

type nodeCoord struct {
	row int
	col int
}

func (n nodeCoord) Valid(maxRow, maxCol int) bool {
	return n.row >= 0 &&
		n.row < maxRow &&
		n.col >= 0 &&
		n.col < maxCol
}
