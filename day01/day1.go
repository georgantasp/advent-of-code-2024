package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	//part1()
	part2()
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0

	var left []int
	var right []int
	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Fields(line)

		l, _ := strconv.Atoi(s[0])
		left = append(left, l)
		r, _ := strconv.Atoi(s[1])
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := range len(left) {
		total += abs(left[i] - right[i])
	}

	fmt.Println("total:", total)
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0

	var left []int
	right := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Fields(line)

		l, _ := strconv.Atoi(s[0])
		left = append(left, l)
		r, _ := strconv.Atoi(s[1])
		right[r] = right[r] + 1

	}

	for _, l := range left {
		n := right[l]
		total += l * n
	}

	fmt.Println("total:", total)
}
