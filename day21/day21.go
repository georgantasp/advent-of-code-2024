package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

//go:embed inputtest
var inputtest string

func main() {
	part1 := calculate(2)
	fmt.Println("part1", part1)

	part2 := calculate(25)
	fmt.Println("part1", part2)
}

func calculate(numberRobots int) int {
	inputs := strings.Split(input, "\n")

	a := arrowPad{}
	n := numberPad{}

	total := 0
	for _, i := range inputs {
		fmt.Println(i)
		all := n.encode(i)
		for range numberRobots {
			all = a.encode(all)
		}
		r := all[0]
		l := len(r)
		num, _ := strconv.Atoi(i[0 : len(i)-1])
		fmt.Println("len", len(r), "n", num)

		total += l * num
	}
	return total
}

type numberPad struct{}

type arrowPad struct{}

type coord struct {
	x, y int
}

func (c *coord) applyButton(button string) {
	switch button {
	case "^":
		c.y -= 1
	case ">":
		c.x += 1
	case "v":
		c.y += 1
	case "<":
		c.x -= 1
	}
}

var arrowPadButtondMap = map[string]coord{
	"^": {x: 1, y: 0},
	"A": {x: 2, y: 0},
	"<": {x: 0, y: 1},
	"v": {x: 1, y: 1},
	">": {x: 2, y: 1},
}

var numberPadButtonMap = map[string]coord{
	"7": {x: 0, y: 0},
	"8": {x: 1, y: 0},
	"9": {x: 2, y: 0},
	"4": {x: 0, y: 1},
	"5": {x: 1, y: 1},
	"6": {x: 2, y: 1},
	"1": {x: 0, y: 2},
	"2": {x: 1, y: 2},
	"3": {x: 2, y: 2},
	"0": {x: 1, y: 3},
	"A": {x: 2, y: 3},
}

func (a arrowPad) encode(buttonsCombinations []string) []string {
	var result []string
	for c := range buttonsCombinations {
		buttonsArray := strings.Split(buttonsCombinations[c], "")

		state := arrowPadButtondMap["A"]
		var combinationResult []string
		for b := range buttonsArray {
			next := arrowPadButtondMap[buttonsArray[b]]
			combo := getMoveCombinations(state, next, coord{x: 0, y: 0})
			combinationResult = combinations(combinationResult, combo)
			state = next
		}

		if result == nil || len(combinationResult[0]) < len(result[0]) {
			result = combinationResult
		} else if len(combinationResult[0]) == len(result[0]) {
			result = append(result, combinationResult...)
		} else {
			fmt.Println("here")
		}
	}

	return result
}

func (n numberPad) encode(code string) []string {
	buttons := strings.Split(code, "")

	var result []string
	state := numberPadButtonMap["A"]
	for b := range buttons {
		next := numberPadButtonMap[buttons[b]]
		combo := getMoveCombinations(state, next, coord{x: 0, y: 3})
		result = combinations(result, combo)
		state = next
	}

	return result
}

func getMoveCombinations(start, end, emptyCoord coord) []string {
	var result string

	if start == end {
		return []string{"A"}
	}

	if end.x-start.x > 0 {
		result += strings.Repeat(">", end.x-start.x)
	}
	if start.x-end.x > 0 {
		result += strings.Repeat("<", start.x-end.x)
	}
	if start.y-end.y > 0 {
		result += strings.Repeat("^", start.y-end.y)
	}
	if end.y-start.y > 0 {
		result += strings.Repeat("v", end.y-start.y)
	}

	var resultPermutations []string
checkLoop:
	for _, r := range permute(result) {
		testState := start
		for _, button := range r {
			testState.applyButton(string(button))
			if testState == emptyCoord {
				continue checkLoop
			}
		}

		resultPermutations = append(resultPermutations, r+"A")
	}

	return resultPermutations
}

func permute(buttons string) []string {
	if len(buttons) == 1 {
		return []string{buttons}
	}

	results := make(map[string]struct{})

	for i, button := range buttons {
		next := buttons[:i] + buttons[i+1:]
		permutations := permute(next)
		for _, p := range permutations {
			results[string(button)+p] = struct{}{}
		}
	}

	unique := make([]string, 0, len(results))
	for k := range results {
		unique = append(unique, k)
	}

	return unique
}

func combinations(perm1, perm2 []string) []string {
	results := make(map[string]struct{})
	if len(perm1) == 0 {
		for _, s2 := range perm2 {
			results[s2] = struct{}{}
		}
	} else if len(perm2) == 0 {
		for _, s1 := range perm1 {
			results[s1] = struct{}{}
		}
	} else {
		for _, s1 := range perm1 {
			for _, s2 := range perm2 {
				results[s1+s2] = struct{}{}
			}
		}
	}

	unique := make([]string, 0, len(results))
	for k := range results {
		unique = append(unique, k)
	}

	return unique
}
