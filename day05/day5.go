package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input_rules
var inputRules string

//go:embed input_updates
var inputUpdaets string

type rule struct {
	first  int
	second int
}

type rulesList []rule

func (rules rulesList) checkUpdate(update []int) int {
	indexes := make(map[int]int, len(update))
	for i, u := range update {
		indexes[u] = i
	}

	for _, r := range rules {
		f, fok := indexes[r.first]
		s, sok := indexes[r.second]
		if !fok || !sok {
			continue
		}
		if f > s {
			return -1
		}
	}

	return update[len(update)/2]
}

func main() {
	ruleLines := strings.Split(inputRules, "\n")
	rules := make(rulesList, len(ruleLines))
	for i, line := range ruleLines {
		ruleS := strings.Split(line, "|")
		first, _ := strconv.Atoi(ruleS[0])
		second, _ := strconv.Atoi(ruleS[1])
		rules[i] = rule{first: first, second: second}
	}

	part1(rules)
	part2(rules)
}

func part1(rules rulesList) {
	updateLines := strings.Split(inputUpdaets, "\n")
	total := 0
	for ul, line := range updateLines {
		updatesS := strings.Split(line, ",")
		updates := make([]int, len(updatesS))
		for i, update := range updatesS {
			updates[i], _ = strconv.Atoi(update)
		}

		ret := rules.checkUpdate(updates)
		if ret != -1 {
			total += ret
		} else {
			fmt.Println("bad line", ul)
		}
	}

	fmt.Println("part1", total)
}

func part2(rules rulesList) {

	updateLines := strings.Split(inputUpdaets, "\n")
	total := 0

	for _, line := range updateLines {
		updatesS := strings.Split(line, ",")
		updates := make([]int, len(updatesS))
		for i, update := range updatesS {
			updates[i], _ = strconv.Atoi(update)
		}

		if rules.checkUpdate(updates) != -1 {
			continue
		}

		sort.Slice(updates, func(i, j int) bool {
			return rules.checkUpdate([]int{updates[i], updates[j]}) != -1
		})

		total += updates[len(updates)/2]

	}

	fmt.Println("part2", total)
}
