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

	inputs := strings.Split(input, "\n")

	a := arrowPad{
		e: arrowPad{
			e: numberPad{},
		},
	}

	total := 0
	for _, i := range inputs {
		fmt.Println(i)
		all := a.encode(i)
		r := all[0]
		l := len(r)
		n, _ := strconv.Atoi(i[0 : len(i)-1])
		fmt.Println("len", len(r), "n", n)

		total += l * n
	}

	fmt.Println("part1", total)
}

type padEncoder interface {
	encode(code string) []string
	decode(code string) string
}

type numberPad struct{}

type arrowPad struct {
	e padEncoder
}

type coord struct {
	x, y int
}

func arrowPadCood(button string) coord {
	switch button {
	case "^":
		return coord{x: 1, y: 0}
	case "A":
		return coord{x: 2, y: 0}
	case "<":
		return coord{x: 0, y: 1}
	case "v":
		return coord{x: 1, y: 1}
	case ">":
		return coord{x: 2, y: 1}
	default:
		return coord{x: 0, y: 0}
	}
}

func arrowPadButton(c coord) string {
	switch c {
	case coord{x: 1, y: 0}:
		return "^"
	case coord{x: 2, y: 0}:
		return "A"
	case coord{x: 0, y: 1}:
		return "<"
	case coord{x: 1, y: 1}:
		return "v"
	case coord{x: 2, y: 1}:
		return ">"
	default:
		return "S"
	}
}

func numberPadCoord(button string) coord {
	switch button {
	case "7":
		return coord{x: 0, y: 0}
	case "8":
		return coord{x: 1, y: 0}
	case "9":
		return coord{x: 2, y: 0}
	case "4":
		return coord{x: 0, y: 1}
	case "5":
		return coord{x: 1, y: 1}
	case "6":
		return coord{x: 2, y: 1}
	case "1":
		return coord{x: 0, y: 2}
	case "2":
		return coord{x: 1, y: 2}
	case "3":
		return coord{x: 2, y: 2}
	case "0":
		return coord{x: 1, y: 3}
	case "A":
		return coord{x: 2, y: 3}
	default:
		return coord{x: 0, y: 3}
	}
}

func numberPadButton(c coord) string {
	switch c {
	case coord{x: 0, y: 0}:
		return "7"
	case coord{x: 1, y: 0}:
		return "8"
	case coord{x: 2, y: 0}:
		return "9"
	case coord{x: 0, y: 1}:
		return "4"
	case coord{x: 1, y: 1}:
		return "5"
	case coord{x: 2, y: 1}:
		return "6"
	case coord{x: 0, y: 2}:
		return "1"
	case coord{x: 1, y: 2}:
		return "2"
	case coord{x: 2, y: 2}:
		return "3"
	case coord{x: 1, y: 3}:
		return "0"
	case coord{x: 2, y: 3}:
		return "A"
	default:
		return "S"
	}
}

func (a arrowPad) encode(code string) []string {
	buttonsCombinations := a.e.encode(code)

	var result []string
	for c, _ := range buttonsCombinations {
		buttonsArray := strings.Split(buttonsCombinations[c], "")
		state := arrowPadCood("A")
		var combinationResult []string
		for b, _ := range buttonsArray {
			next := arrowPadCood(buttonsArray[b])
			combo := from(state, next, coord{x: 0, y: 0})
			combinationResult = combinations(combinationResult, combo)
			state = next
		}
		result = append(result, combinationResult...)
		result = shortest(result)
	}

	return result
}

func (n numberPad) encode(code string) []string {
	buttons := strings.Split(code, "")

	var result []string
	state := numberPadCoord("A")
	for b, _ := range buttons {
		next := numberPadCoord(buttons[b])
		result = combinations(result, from(state, next, coord{x: 0, y: 3}))
		state = next
	}

	result = shortest(result)
	return result
}

func (a arrowPad) decode(code string) string {
	fmt.Println("decode", code, "len", len(code), "countA", countA(code))
	buttons := strings.Split(code, "")

	result := ""
	state := arrowPadCood("A")
	for _, button := range buttons {
		if button == "A" {
			result += arrowPadButton(state)
		} else {
			to(&state, button)
		}
	}

	return a.e.decode(result)
}

func (a numberPad) decode(code string) string {
	fmt.Println("decode", code, "len", len(code), "countA", countA(code))
	buttons := strings.Split(code, "")

	result := ""
	state := numberPadCoord("A")
	for _, button := range buttons {
		if button == "A" {
			result += numberPadButton(state)
		} else {
			to(&state, button)
		}
	}

	return result
}

func from(state, buttonCoord, spaceCoord coord) []string {
	var result string

	if state == buttonCoord {
		return []string{"A"}
	}

	if buttonCoord.x-state.x > 0 {
		result += strings.Repeat(">", buttonCoord.x-state.x)
	}
	if state.x-buttonCoord.x > 0 {
		result += strings.Repeat("<", state.x-buttonCoord.x)
	}
	if state.y-buttonCoord.y > 0 {
		result += strings.Repeat("^", state.y-buttonCoord.y)
	}
	if buttonCoord.y-state.y > 0 {
		result += strings.Repeat("v", buttonCoord.y-state.y)
	}

	var resultPermutations []string
checkLoop:
	for _, r := range permute(result) {
		copyState := state
		for _, button := range r {
			to(&copyState, string(button))
			if copyState == spaceCoord {
				continue checkLoop
			}
		}

		resultPermutations = append(resultPermutations, r+"A")
	}

	return resultPermutations
}

func to(state *coord, button string) {
	switch button {
	case "^":
		state.y -= 1
	case ">":
		state.x += 1
	case "v":
		state.y += 1
	case "<":
		state.x -= 1
	}
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
	for k, _ := range results {
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
	for k, _ := range results {
		unique = append(unique, k)
	}

	return unique
}

func countA(s string) int {
	count := 0
	for _, char := range s {
		if char == 'A' {
			count++
		}
	}
	return count
}

func shortest(results []string) []string {
	var s []string
	for i, r := range results {
		if i == 0 {
			s = []string{r}
		} else {
			if len(r) < len(s[0]) {
				s = []string{r}
			} else if len(r) == len(s[0]) {
				s = append(s, r)
			}
		}
	}
	return s
}
