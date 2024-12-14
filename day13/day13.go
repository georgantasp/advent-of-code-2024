package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

var re = regexp.MustCompile(`(?m).+?X[+,=](\d+), Y[+,=](\d+)`)

const factor = 10000000000000

func main() {

	machineS := strings.Split(input, "\n\n")

	totalPart1 := 0
	totalPart2 := 0
	for _, m := range machineS {
		lines := strings.Split(m, "\n")
		buttonA := lineToCoord(lines[0])
		buttonB := lineToCoord(lines[1])
		prize := lineToCoord(lines[2])
		totalPart1 += solveBrute(buttonA, buttonB, prize)
		totalPart2 += solveBetter(buttonA, buttonB, coord{prize.x + factor, prize.y + factor})
	}

	fmt.Println("part1", totalPart1)
	fmt.Println("part2", totalPart2)
}

type coord struct {
	x, y int
}

func lineToCoord(line string) coord {
	matches := re.FindAllStringSubmatch(line, -1)
	x, _ := strconv.Atoi(matches[0][1])
	y, _ := strconv.Atoi(matches[0][2])

	return coord{x, y}
}

func solveBrute(a, b, p coord) int {
	aPresses := 0 //costs 3
	for {
		bPresses := 0 //costs 1
		for {
			if aPresses*a.x+bPresses*b.x == p.x &&
				aPresses*a.y+bPresses*b.y == p.y {
				return aPresses*3 + bPresses
			}
			if aPresses*a.x+bPresses*b.x > p.x &&
				aPresses*a.y+bPresses*b.y > p.y {
				break
			}
			bPresses++
		}
		if aPresses*a.x > p.x ||
			aPresses*a.y > p.y {
			break
		}
		aPresses++
	}

	return 0
}

func solveBetter(a coord, b coord, p coord) int {
	det := a.x*b.y - a.y*b.x
	if det == 0 {
		return 0
	}

	aNumerator := p.x*b.y - p.y*b.x
	bNumerator := a.x*p.y - a.y*p.x

	if aNumerator%det != 0 || bNumerator%det != 0 {
		return 0
	}

	aPresses := aNumerator / det
	bPresses := bNumerator / det

	return 3*aPresses + bPresses
}
