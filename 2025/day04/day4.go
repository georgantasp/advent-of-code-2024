package main

import (
	_ "embed"
	"strings"
)

//go:embed input
var input string

var neighbors = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1} /*self*/, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func main() {
	var warehouse [][]rune

	for _, line := range strings.Split(input, "\n") {
		warehouse = append(warehouse, []rune(line))
	}

	println("part 1: ", len(getForkliftable(warehouse))) //1437

	var totalForkliftable int
	for {
		forkliftable := getForkliftable(warehouse)
		if len(forkliftable) == 0 {
			break
		}
		totalForkliftable += len(forkliftable)
		for _, coord := range forkliftable {
			warehouse[coord[0]][coord[1]] = 'x'
		}
	}
	println("part 2: ", totalForkliftable)
}

func getForkliftable(warehouse [][]rune) [][2]int {
	var forkliftable [][2]int
	for i, row := range warehouse {
		for j, col := range row {
			if col == '@' {
				count := 0
				for _, n := range neighbors {
					if isPaper(warehouse, i+n[0], j+n[1]) {
						count++
						if count >= 4 {
							break
						}
					}
				}

				if count < 4 {
					forkliftable = append(forkliftable, [2]int{i, j})
				}
			}
		}
	}
	return forkliftable
}

func isPaper(warehouse [][]rune, i, j int) bool {
	if i < 0 || i >= len(warehouse) {
		return false
	}
	if j < 0 || j >= len(warehouse[i]) {
		return false
	}
	return warehouse[i][j] == '@'
}
