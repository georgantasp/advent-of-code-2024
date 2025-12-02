package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input
var input string

func main() {
	//part1()
	part2()
}

func part1() {
	re := regexp.MustCompile(`(?m)(mul\((\d+?),(\d+?)\))`)

	var total int
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		l, _ := strconv.Atoi(match[2])
		r, _ := strconv.Atoi(match[3])

		total += l * r
	}
	fmt.Println(total)
}

func part2() {
	re := regexp.MustCompile(`(mul\((\d+?),(\d+?)\))|(don't\(\))|(do\(\))`)

	var total int
	enabled := true
	for _, match := range re.FindAllStringSubmatch(input, -1) {

		switch {
		case enabled && match[1] != "":
			l, _ := strconv.Atoi(match[2])
			r, _ := strconv.Atoi(match[3])

			total += l * r
		case match[4] != "":
			enabled = false
		case match[5] != "":
			enabled = true
		}
	}
	fmt.Println(total)
}
