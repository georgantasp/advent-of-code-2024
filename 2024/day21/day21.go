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

	a := arrowPad{cache: make(map[robotMove][]string)}
	n := numberPad{}

	total := 0
	for _, i := range inputs {
		fmt.Println(i)
		all := n.encode(i)
		for rNum := range numberRobots {
			all = a.encode(all)
			fmt.Println("done rNum", rNum)
		}
		num, _ := strconv.Atoi(i[0 : len(i)-1])
		fmt.Println("len", all[0].Length(), "n", num)

		total += all[0].Length() * num
	}
	return total
}

type numberPad struct{}

type robotMove struct {
	start coord
	end   coord
}

type arrowPad struct {
	cache map[robotMove][]string
}

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

type padCombinations struct {
	sequences []sequence
	length    int
}

func (p *padCombinations) Length() int {
	if p.length != 0 {
		return p.length
	}

	for _, s := range p.sequences {
		p.length += len(s[0])
	}

	return p.length
}

type sequence []string

func (a arrowPad) encode(padCombos []padCombinations) []padCombinations {
	var result []padCombinations
	for _, pc := range padCombos {
		for _, s := range pc.sequences {

			// sequence cache

			buttonsArray := strings.Split(s, "")

			state := arrowPadButtondMap["A"]
			var combinationResult padCombinations
			for b := range buttonsArray {
				next := arrowPadButtondMap[buttonsArray[b]]

				if combo, ok := a.cache[robotMove{state, next}]; ok {
					combinationResult.sequences = append(combinationResult.sequences, combo)
				} else {
					combo = getMoveCombinations(state, next, coord{x: 0, y: 0})
					a.cache[robotMove{state, next}] = combo
					combinationResult.sequences = append(combinationResult.sequences, combo)
				}

				state = next
			}

			if result == nil || combinationResult.Length() < result[0].Length() {
				result = []padCombinations{combinationResult}
			} else if combinationResult.Length() == result[0].Length() {
				result = append(result, combinationResult)
			}
		}
	}

	return result
}

func (n numberPad) encode(code string) []padCombinations {
	buttons := strings.Split(code, "")

	var result []sequence
	state := numberPadButtonMap["A"]
	for b := range buttons {
		next := numberPadButtonMap[buttons[b]]
		combo := getMoveCombinations(state, next, coord{x: 0, y: 3})
		result = append(result, combo)
		state = next
	}

	return []padCombinations{{
		sequences: result,
	}}
}

func getMoveCombinations(start, end, emptyCoord coord) sequence {

	if start == end {
		return sequence{"A"}
	}

	var lr string
	if end.x-start.x > 0 {
		lr = strings.Repeat(">", end.x-start.x)
	} else if start.x-end.x > 0 {
		lr = strings.Repeat("<", start.x-end.x)
	}

	var ud string
	if start.y-end.y > 0 {
		ud = strings.Repeat("^", start.y-end.y)
	} else if end.y-start.y > 0 {
		ud = strings.Repeat("v", end.y-start.y)
	}

	var results sequence
	if lr == "" {
		results = sequence{ud + "A"}
	} else if ud == "" {
		results = sequence{lr + "A"}
	} else {
		results = sequence{
			lr + ud + "A",
			ud + lr + "A",
		}
	}

	var resultPermutations sequence
checkLoop:
	for _, r := range results {
		testState := start
		for _, button := range r {
			testState.applyButton(string(button))
			if testState == emptyCoord {
				continue checkLoop
			}
		}

		resultPermutations = append(resultPermutations, r)
	}

	return resultPermutations
}

//func combinations(perm1, perm2 []string) []string {
//	results := make(map[string]struct{})
//	if len(perm1) == 0 {
//		for _, s2 := range perm2 {
//			results[s2] = struct{}{}
//		}
//	} else if len(perm2) == 0 {
//		for _, s1 := range perm1 {
//			results[s1] = struct{}{}
//		}
//	} else {
//		for _, s1 := range perm1 {
//			for _, s2 := range perm2 {
//				results[s1+s2] = struct{}{}
//			}
//		}
//	}
//
//	unique := make([]string, 0, len(results))
//	for k := range results {
//		unique = append(unique, k)
//	}
//
//	return unique
//}
