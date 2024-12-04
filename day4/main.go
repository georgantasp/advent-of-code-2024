package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")

	puzz := make([][]string, len(lines))

	for i, line := range lines {
		puzz[i] = strings.Split(line, "")
	}

	//part1(puzz)
	part2(puzz)
}

func part1(puzz [][]string) {

	total := checkLinesForXMAS(puzz)                          // gets forward and backward
	total += checkLinesForXMAS(rowsToCols(puzz))              // gets up and down
	total += checkLinesForXMAS(rowsToDiag(puzz))              // gets diag ne -> sw, sw -> ne
	total += checkLinesForXMAS(rowsToDiag(reverseRows(puzz))) // gets diag nw -> se, se -> nw

	fmt.Println(total)
}

// FRIGGIN BRUTE FORCE
func part2(puzz [][]string) {
	total := 0
	for x := range puzz {
		if x == 0 || x == len(puzz)-1 {
			continue
		}
		for y := range puzz[x] {
			if y == 0 || y == len(puzz[x])-1 {
				continue
			}
			if puzz[x][y] == "A" {

				cross := strings.Join([]string{puzz[x-1][y-1], puzz[x-1][y+1], puzz[x+1][y-1], puzz[x+1][y+1]}, "")
				if cross == "MMSS" || cross == "SSMM" || cross == "MSMS" || cross == "SMSM" {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func checkLinesForXMAS(puzz [][]string) int {
	total := 0
	for _, puzzLine := range puzz {
		line := strings.Join(puzzLine, "")

		total += strings.Count(line, "XMAS")
		total += strings.Count(line, "SAMX")
	}
	return total
}

func rowsToCols(puzz [][]string) [][]string {
	cols := make([][]string, len(puzz[0]))
	for x := range len(puzz) {
		for y := range len(puzz[0]) {
			if x == 0 {
				cols[y] = make([]string, len(puzz))
			}
			cols[y][x] = puzz[x][y]
		}
	}

	return cols
}

func rowsToDiag(puzz [][]string) [][]string {
	diagRows := len(puzz) + len(puzz[0]) - 1
	diag := make([][]string, diagRows)
	for x := range len(puzz) {
		for y := range len(puzz[0]) {
			i := x + y
			diag[i] = append(diag[i], puzz[x][y])
		}
	}

	return diag
}

func reverseRows(puzz [][]string) [][]string {
	rows := make([][]string, len(puzz))
	for x, line := range puzz {
		rows[x] = make([]string, len(line))
		copy(rows[x], line)
		slices.Reverse(rows[x])
	}

	return rows
}
