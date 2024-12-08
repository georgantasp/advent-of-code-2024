package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input
var input string

func main() {
	fmt.Print("\033[?25l")
	//part1()
	part2()
	fmt.Print("\033[?25h")
}

type uniqCoordMap map[nodeCoord]struct{}

func (u uniqCoordMap) Add(n nodeCoord) {
	if _, ok := u[n]; !ok {
		u[n] = struct{}{}
	} else {
		return
	}

	// animate
	if len(u)%60 != 1 {
		n.Print()
	} else {
		// complete reprint
		clearScreen()
		moveCursor(1, 1)
		fmt.Print(input)
		for k, _ := range u {
			k.Print()
		}
	}

	time.Sleep(10 * time.Millisecond)
}

func part1() {
	lines := strings.Split(input, "\n")

	maxRow := len(lines)
	maxCol := len(lines[0])
	antennaMap := make(map[string][]nodeCoord)
	uniqueAntinodes := make(uniqCoordMap)

	for i, line := range lines {
		for j, node := range strings.Split(line, "") {
			if node == "." {
				continue
			}
			coord := nodeCoord{row: i, col: j}
			if _, ok := antennaMap[node]; ok {
				for _, otherCoord := range antennaMap[node] {
					rowDist := coord.row - otherCoord.row
					colDist := coord.col - otherCoord.col

					testCoord1 := nodeCoord{row: coord.row + rowDist, col: coord.col + colDist}
					testCoord2 := nodeCoord{row: otherCoord.row - rowDist, col: otherCoord.col - colDist}
					if testCoord1.Valid(maxRow, maxCol) {
						uniqueAntinodes.Add(testCoord1)
					}
					if testCoord2.Valid(maxRow, maxCol) {
						uniqueAntinodes.Add(testCoord2)
					}
				}
			}
			antennaMap[node] = append(antennaMap[node], coord)
		}
	}

	fmt.Println("part1", len(uniqueAntinodes))
}

func part2() {
	lines := strings.Split(input, "\n")

	maxRow := len(lines)
	maxCol := len(lines[0])
	antennaMap := make(map[string][]nodeCoord)
	uniqueAntinodes := make(uniqCoordMap)

	for i, line := range lines {
		for j, node := range strings.Split(line, "") {
			if node == "." {
				continue
			}
			coord := nodeCoord{row: i, col: j}
			if _, ok := antennaMap[node]; ok {
				for _, otherCoord := range antennaMap[node] {
					uniqueAntinodes.Add(coord)
					uniqueAntinodes.Add(otherCoord)

					rowDist := coord.row - otherCoord.row
					colDist := coord.col - otherCoord.col

					testCoord1 := coord
					for {
						testCoord1 = nodeCoord{row: testCoord1.row + rowDist, col: testCoord1.col + colDist}
						if testCoord1.Valid(maxRow, maxCol) {
							uniqueAntinodes.Add(testCoord1)
						} else {
							break
						}
					}

					testCoord2 := otherCoord
					for {
						testCoord2 = nodeCoord{row: testCoord2.row - rowDist, col: testCoord2.col - colDist}
						if testCoord2.Valid(maxRow, maxCol) {
							uniqueAntinodes.Add(testCoord2)
						} else {
							break
						}
					}
				}
			}
			antennaMap[node] = append(antennaMap[node], coord)
		}
	}

	moveCursor(maxRow+2, 0)
	fmt.Println("part2", len(uniqueAntinodes))
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

func (n nodeCoord) Print() {
	fmt.Printf("\033[%d;%dH#", n.row+1, n.col+1)
}

func moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J") // Clear the screen
}
