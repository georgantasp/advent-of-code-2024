package main

import (
	_ "embed"
	"fmt"
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

	curX, curY := 0, 0
startLoop:
	for y, row := range puzz {
		for x, col := range row {
			if col == "^" {
				curX, curY = x, y
				break startLoop
			}
		}
	}

	part1(copyPuzz(puzz), curX, curY)
	part2(copyPuzz(puzz), curX, curY)
}

func part1(puzz [][]string, curX, curY int) {
	total, _ := moveLoop(puzz, curX, curY)

	fmt.Println("part1", total)
}

func part2(puzz [][]string, curX, curY int) {
	total := 0
	for i := range len(puzz) {
		for j := range len(puzz[i]) {
			if puzz[i][j] == "." {
				cPuzz := copyPuzz(puzz)
				cPuzz[i][j] = "#"
				if _, isLoop := moveLoop(cPuzz, curX, curY); isLoop {
					total++
				}
			}
		}
	}

	fmt.Println("part2", total)
}

func copyPuzz(puzz [][]string) [][]string {
	c := make([][]string, len(puzz))
	for i := range c {
		c[i] = make([]string, len(puzz[i]))
		for j := range c[i] {
			c[i][j] = puzz[i][j]
		}
	}
	return c
}

func moveLoop(puzz [][]string, curX, curY int) (int, bool) {
	total := 1
	direction := "n"

	for {
		//fmt.Println(len(puzz[0]), len(puzz), direction, curX, curY, total)
		switch direction {
		case "n":
			if curY-1 == -1 {
				return total, false
			} else if puzz[curY-1][curX] == "#" {
				direction = "e"
			} else {
				if puzz[curY-1][curX] == "." {
					total++
					puzz[curY-1][curX] = direction
				} else if strings.Contains(puzz[curY-1][curX], direction) {
					return total, true
				} else {
					puzz[curY-1][curX] = puzz[curY-1][curX] + direction
				}
				curY = curY - 1
			}
		case "e":
			if curX+1 == len(puzz) {
				return total, false
			} else if puzz[curY][curX+1] == "#" {
				direction = "s"
			} else {
				if puzz[curY][curX+1] == "." {
					puzz[curY][curX+1] = direction
					total++
				} else if strings.Contains(puzz[curY][curX+1], direction) {
					return total, true
				} else {
					puzz[curY][curX+1] = puzz[curY][curX+1] + direction
				}
				curX = curX + 1
			}
		case "s":
			if curY+1 == len(puzz[0]) {
				return total, false
			} else if puzz[curY+1][curX] == "#" {
				direction = "w"
			} else {
				if puzz[curY+1][curX] == "." {
					puzz[curY+1][curX] = direction
					total++
				} else if strings.Contains(puzz[curY+1][curX], direction) {
					return total, true
				} else {
					puzz[curY+1][curX] = puzz[curY+1][curX] + direction
				}
				curY = curY + 1
			}
		case "w":
			if curX-1 == -1 {
				return total, false
			} else if puzz[curY][curX-1] == "#" {
				direction = "n"
			} else {
				if puzz[curY][curX-1] == "." {
					puzz[curY][curX-1] = direction
					total++
				} else if strings.Contains(puzz[curY][curX-1], direction) {
					return total, true
				} else {
					puzz[curY][curX-1] += direction
				}
				curX = curX - 1
			}
		}
	}
}
