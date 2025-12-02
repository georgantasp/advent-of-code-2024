package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	ranges := strings.Split(input, ",")

	invalidIdSumPart1 := 0
	invalidIdSumPart2 := 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		// iterate the range
		for i := start; i <= end; i++ {
			iAsString := strconv.Itoa(i)
			if len(iAsString)%2 == 0 && iAsString[0:len(iAsString)/2] == iAsString[len(iAsString)/2:] {
				invalidIdSumPart1 += i
			}

			for j := 1; j < len(iAsString); j++ {
				if checkEqualByPart(iAsString, j) {
					invalidIdSumPart2 += i
					break
				}
			}
		}
	}

	println("part1:", invalidIdSumPart1)
	println("part2:", invalidIdSumPart2)
}

func checkEqualByPart(s string, partLength int) bool {
	if len(s)%partLength != 0 {
		return false
	}

	part0 := s[0:partLength]
	for i := partLength; i < len(s); i += partLength {
		checkPart := s[i : i+partLength]
		if checkPart != part0 {
			return false
		}
	}

	println(s)
	return true
}

// 17318342683 too high
// 17298174201
// 17298173706 too low
