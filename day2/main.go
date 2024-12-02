package main

import (
	"bufio"
	_ "embed"
	"fmt"
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

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Fields(line)

		if safe, _ := isReportSafe(s); safe {
			total += 1
		}
	}

	fmt.Println("total:", total)
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Fields(line)

		if safe, badIndex := isReportSafe(s); safe {
			total += 1
		} else {
			newReport := []string{}
			newReport = append(newReport, s[:badIndex]...)
			newReport = append(newReport, s[badIndex+1:]...)
			if safe, _ = isReportSafe(newReport); safe {
				total += 1
				continue
			}

			// try with the first record removed
			newReport = s[1:]
			if safe, _ = isReportSafe(newReport); safe {
				total += 1
				continue
			}

			// try with the second record removed
			newReport = append([]string{s[0]}, s[2:]...)
			if safe, _ = isReportSafe(newReport); safe {
				total += 1
				continue
			}
		}
	}

	fmt.Println("total:", total)
}

func isReportSafe(s []string) (bool, int) {
	safe := true
	var last int
	var inc bool

	badIndex := -1
	for i, ns := range s {
		n, _ := strconv.Atoi(ns)

		if i == 0 {
			last = n
			continue
		}

		if i == 1 {
			inc = last < n
		}

		if inc {
			if n <= last {
				safe = false
				badIndex = i
				break
			}
			if n-last > 3 {
				safe = false
				badIndex = i
				break
			}
		} else {
			if n >= last {
				safe = false
				badIndex = i
				break
			}
			if last-n > 3 {
				safe = false
				badIndex = i
				break
			}
		}

		last = n
	}

	return safe, badIndex
}
