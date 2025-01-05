package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

type secretNumber int

func (sn secretNumber) mix(in int) secretNumber {
	return secretNumber(int(sn) ^ in)
}

func (sn secretNumber) prune() secretNumber {
	return secretNumber(int(sn) % 16777216)
}

func (sn secretNumber) next() secretNumber {
	r := sn

	r = r.mix(int(r * 64)).prune()

	r = r.mix(int(r / 32)).prune()

	r = r.mix(int(r * 2048)).prune()

	return r
}

func main() {

	buyersInput := strings.Split(input, "\n")

	total := 0
	for _, buyerInput := range buyersInput {
		buyerInitial, _ := strconv.Atoi(buyerInput)
		buyerSn := secretNumber(buyerInitial)
		fmt.Println(buyerSn)
		for range 2000 {
			buyerSn = buyerSn.next()
		}
		total += int(buyerSn)
	}

	fmt.Println("part1", total)
}
