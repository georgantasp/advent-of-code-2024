package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed inputtest
var inputtest string

//go:embed input
var input string

func main() {
	inputs := strings.Split(inputtest, "\n\n")

	towelsInput := strings.Split(inputs[0], ", ")
	towels := []string{}
	for _, t := range towelsInput {
		towels = append(towels, t)
	}
	designs := strings.Split(inputs[1], "\n")

	total := 0
	totalWays := 0
	possibleDesigns := from(towels)
	impossibleDesigns := make(map[string]struct{})
	for i, design := range designs {
		fmt.Println(design)
		fmt.Print("i ", i)
		if ways := designPossible(design, possibleDesigns, impossibleDesigns); ways > 0 {
			fmt.Print(" ", ways)
			total++
			totalWays += ways
		}
		fmt.Println()
	}
	fmt.Println("part1", total)
	fmt.Println("part2", totalWays)
}

func from(towels []string) map[string]int {

	sort.Slice(towels, func(i, j int) bool {
		if len(towels[i]) != len(towels[j]) {
			return len(towels[i]) < len(towels[j])
		}
		return towels[i] < towels[j]
	})

	possibleDesigns := make(map[string]int)
	for _, towel := range towels {
		if len(towel) == 1 {
			possibleDesigns[towel] = 1
		} else {
			possibleDesigns[towel] = 1 + collectWays(towel, possibleDesigns)
		}
	}

	return possibleDesigns
}

func collectWays(towelPart string, possibleDesigns map[string]int) int {
	if ways, ok := possibleDesigns[towelPart]; ok {
		return ways
	}

	if len(towelPart) == 1 {
		return 0
	}

	ways := 0
	for m := 1; m < len(towelPart); m++ {
		a := towelPart[:m]
		b := towelPart[m:]

		aWays := collectWays(a, possibleDesigns)
		bWays := collectWays(b, possibleDesigns)

		if aWays > 0 && bWays > 0 {
			ways += 1
		}
	}
	return ways
}

func designPossible(design string, possibleDesigns map[string]int, impossibleDesigns map[string]struct{}) int {
	if ways, ok := possibleDesigns[design]; !ok && len(design) == 1 {
		return 0
	} else if ok {
		return ways
	}

	ways := 0
	for m := 1; m < len(design); m++ {
		a := design[:m]
		b := design[m:]

		aPossible := 0
		if _, ok := impossibleDesigns[a]; !ok {
			aPossible = designPossible(a, possibleDesigns, impossibleDesigns)
		}
		bPossible := 0
		if _, ok := impossibleDesigns[b]; !ok {
			bPossible = designPossible(b, possibleDesigns, impossibleDesigns)
		}

		switch {
		case aPossible < 1:
			impossibleDesigns[a] = struct{}{}
		case bPossible < 1:
			impossibleDesigns[b] = struct{}{}
		default:
			ways += 1
		}
	}

	possibleDesigns[design] = ways

	return ways
}
