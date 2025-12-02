package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

type counts struct {
	part1 int
	part2 int
}

func main() {
	inputLines := strings.Split(input, "\n")

	dial := 50
	c := counts{}
	for _, line := range inputLines {
		direction := string(line[0:1])
		num, _ := strconv.Atoi(line[1:])

		dial = c.run(dial, direction, num)
	}

	println("part1:", c.part1)
	println("part2:", c.part2)
}

func (c *counts) run(dial int, direction string, num int) int {
	prevDial := dial
	switch direction {
	case "L":
		dial -= num
	case "R":
		dial += num
	}

	switch {
	case dial < 0:
		if prevDial != 0 {
			c.part2++
		}
		c.part2 += (-dial / 100)

		if dial%100 == 0 {
			c.part1++
			dial = 0
		} else {
			dial = dial%100 + 100
		}
	case dial == 0:
		c.part1++
		c.part2++
	case dial > 99:
		c.part2 += dial / 100

		if dial%100 == 0 {
			c.part1++
			dial = 0
		} else {
			dial = dial % 100
		}
	}

	return dial
}
