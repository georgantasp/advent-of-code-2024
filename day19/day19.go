package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputtest
var inputtest string

//go:embed input
var input string

func main() {
	inputs := strings.Split(input, "\n\n")

	towelsInput := strings.Split(inputs[0], ", ")
	towels := []string{}
	for _, t := range towelsInput {
		towels = append(towels, t)
	}
	designs := strings.Split(inputs[1], "\n")

	total := 0
	totalWays := 0
	towelMap := from(towels)
	impossibleDesigns := make(map[string]struct{})
	possibleDesigns := make(map[string]int)
	for _, design := range designs {
		if ways := designPossible(design, towelMap, impossibleDesigns, possibleDesigns); ways > 0 {
			total++
			totalWays += ways
		}
	}
	fmt.Println("part1", total)
	fmt.Println("part2", totalWays)
}

func from(towels []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, towel := range towels {
		m[towel] = struct{}{}
	}

	return m
}

func designPossible(design string, towels map[string]struct{}, impossibleDesigns map[string]struct{}, possibleDesigns map[string]int) int {
	if _, ok := impossibleDesigns[design]; ok {
		return 0
	}

	if d, ok := possibleDesigns[design]; ok {
		return d
	}

	ways := 0
	if _, ok := towels[design]; ok {
		ways = 1
	}
	for m := 1; m < len(design); m++ {
		a := design[:m]
		b := design[m:]

		_, aPossible := towels[a]
		bPossible := designPossible(b, towels, impossibleDesigns, possibleDesigns)

		if aPossible && bPossible > 0 {
			ways += bPossible
		}
	}

	if ways == 0 {
		impossibleDesigns[design] = struct{}{}
	} else {
		possibleDesigns[design] = ways
	}

	return ways
}

/// bbr
///  (b)(b)(r)
///  (b)(br)
