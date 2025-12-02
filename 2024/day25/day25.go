package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

//go:embed inputtest
var inputtest string

type lock [5]int

type key [5]int

func (k key) fits(l lock) bool {
	return 5-l[0] >= k[0] &&
		5-l[1] >= k[1] &&
		5-l[2] >= k[2] &&
		5-l[3] >= k[3] &&
		5-l[4] >= k[4]
}

func main() {
	inputBlocks := strings.Split(input, "\n\n")

	locks := []lock{}
	keys := []key{}

	for _, block := range inputBlocks {
		lines := strings.Split(block, "\n")

		if lines[0] == "#####" {
			locks = append(locks, parse(lines))
		} else {
			keys = append(keys, parse(lines))
		}
	}

	part1 := 0
	for _, l := range locks {
		for _, k := range keys {
			if k.fits(l) {
				part1++
			}
		}
	}

	fmt.Println("part1", part1)
}

func parse(lines []string) [5]int {
	var item [5]int
	for i, line := range lines {
		if i == 0 || i == len(lines)-1 {
			continue
		}
		for j, char := range line {
			if char == '#' {
				item[j] += 1
			}
		}
	}
	return item
}
