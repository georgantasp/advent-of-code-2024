package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	banks := strings.Split(input, "\n")

	part1Sum := 0
	for _, bank := range banks {
		num := getBankMax(bank, 2)

		println(bank)
		println(num)
		println("")
		part1Sum += num
	}

	println("Part 1:", part1Sum)

	part2Sum := 0
	for _, bank := range banks {
		num := getBankMax(bank, 12)

		println(bank)
		println(num)
		println("")
		part2Sum += num
	}

	println("Part 2:", part2Sum)
}

func getBankMax(bank string, battSize int) int {
	finalBatt := ""
	start := 0
	for i := range battSize {
		end := len(bank) - (battSize - i - 1)
		batt, battI := getNextMaxBatt(bank, start, end)
		finalBatt += batt
		start = battI + 1
	}

	num, _ := strconv.Atoi(finalBatt)
	return num
}

func getNextMaxBatt(bank string, start int, end int) (string, int) {
	var maxBatt string
	var maxI int
	for i := start; i < end; i++ {
		if bank[i:i+1] > maxBatt {
			maxBatt = bank[i : i+1]
			maxI = i
		}
	}
	return maxBatt, maxI
}
