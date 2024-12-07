package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	equations := strings.Split(input, "\n")

	totalPart1 := 0
	totalPart2 := 0
	for _, equation := range equations {
		s := strings.Split(equation, ":")

		testValue, _ := strconv.Atoi(s[0])
		numbersS := strings.Split(strings.Trim(s[1], " "), " ")
		numbers := make([]int, len(numbersS))
		for i, n := range numbersS {
			numbers[i], _ = strconv.Atoi(n)
		}

		if canCombinePart1(testValue, numbers) {
			totalPart1 += testValue
		}
		if canCombinePart2(testValue, 0, numbers) {
			totalPart2 += testValue
		}
	}

	fmt.Println("total part1", totalPart1)
	fmt.Println("total part2", totalPart2)
}

func canCombinePart1(testValue int, numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}
	if len(numbers) == 2 {
		return testValue == numbers[0]+numbers[1] ||
			testValue == numbers[0]*numbers[1]
	}

	return canCombinePart1(testValue-numbers[len(numbers)-1], numbers[:len(numbers)-1]) ||
		(testValue%numbers[len(numbers)-1] == 0 &&
			canCombinePart1(testValue/numbers[len(numbers)-1], numbers[:len(numbers)-1]))
}

func canCombinePart2(testValue int, runningTest int, numbers []int) bool {
	if len(numbers) == 0 {
		return testValue == runningTest
	}

	return canCombinePart2(testValue, runningTest+numbers[0], numbers[1:]) ||
		canCombinePart2(testValue, runningTest*numbers[0], numbers[1:]) ||
		canCombinePart2(testValue, concatNumbers(runningTest, numbers[0]), numbers[1:])

}

func concatNumbers(numbers1, numbers2 int) int {
	i, _ := strconv.Atoi(strconv.Itoa(numbers1) + strconv.Itoa(numbers2))
	return i
}
