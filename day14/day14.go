package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed inputtest
var inputtest string

//go:embed input
var input string

var re = regexp.MustCompile(`(?m)p=(\d+?),(\d+?) v=(-?\d+?),(-?\d+?)$`)

type robot struct {
	x  int
	y  int
	vX int
	vY int
}

func (r *robot) move(x int, y int) {
	r.x = wrap(r.x, r.vX, x)
	r.y = wrap(r.y, r.vY, y)
}

func wrap(p, v, max int) int {
	newP := p + v
	if newP >= 0 && newP < max {
		return newP
	} else if newP < 0 {
		return newP + max
	} else {
		return newP - max
	}
}

func (r *robot) quad(xMax, yMax int) int {
	xMid := xMax / 2
	yMid := yMax / 2
	switch {
	case r.x < xMid && r.y < yMid:
		return 1
	case r.x > xMid && r.y < yMid:
		return 2
	case r.x < xMid && r.y > yMid:
		return 3
	case r.x > xMid && r.y > yMid:
		return 4
	default:
		return 0
	}
}

func main() {

	matches := re.FindAllStringSubmatch(input, -1)

	robots := make([]robot, len(matches))
	for i, _ := range matches {
		robots[i] = robot{}
		robots[i].x, _ = strconv.Atoi(matches[i][1])
		robots[i].y, _ = strconv.Atoi(matches[i][2])
		robots[i].vX, _ = strconv.Atoi(matches[i][3])
		robots[i].vY, _ = strconv.Atoi(matches[i][4])
	}

	xMax := 101
	yMax := 103

	sep := "-"
	for range 6 {
		sep = sep + sep
	}

	for sec := range xMax * yMax {
		printField(robots, xMax, yMax)
		fmt.Println("sec", sec, sep)
		fmt.Println()
		for i, _ := range robots {
			robots[i].move(xMax, yMax)
		}
	}

	part1Quads := make(map[int]int)
	part1Quads[0] = 0
	part1Quads[1] = 0
	part1Quads[2] = 0
	part1Quads[3] = 0
	part1Quads[4] = 0
	for _, r := range robots {
		part1Quads[r.quad(xMax, yMax)]++
	}

	total := part1Quads[1] * part1Quads[2] * part1Quads[3] * part1Quads[4]

	fmt.Println("part1", total)
}

func printField(robots []robot, xMax, yMax int) {
	for i := range xMax {
		for j := range yMax {
			found := false
			for r := range robots {
				if robots[r].x == i && robots[r].y == j {
					found = true
					break
				}
			}
			if found {
				fmt.Printf("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
