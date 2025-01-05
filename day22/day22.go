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

func (sn secretNumber) cost() int {
	return int(sn) % 10
}

func main() {

	buyersInput := strings.Split(input, "\n")

	// changes to buyer num to banana count
	changes := map[[4]int]map[int]int{}

	total := 0
	for buyerNum, buyerInput := range buyersInput {
		buyerInitial, _ := strconv.Atoi(buyerInput)
		buyerSn := secretNumber(buyerInitial)
		prevCost := buyerSn.cost()
		prevChanges := [4]int{}
		for i := range 2000 {
			buyerSn = buyerSn.next()
			currCost := buyerSn.cost()
			currChanges := [4]int{prevChanges[1], prevChanges[2], prevChanges[3], currCost - prevCost}
			if i >= 3 {
				if m, ok := changes[currChanges]; ok {
					if _, ok2 := m[buyerNum]; !ok2 {
						m[buyerNum] = currCost
					}
				} else {
					changes[currChanges] = map[int]int{buyerNum: currCost}
				}

			}
			prevCost = buyerSn.cost()
			prevChanges = currChanges
		}
		total += int(buyerSn)
	}

	fmt.Println("part1", total)

	mostBananas := 0
	for _, v := range changes {
		changesTotal := 0
		for _, b := range v {
			changesTotal += b
		}
		if changesTotal > mostBananas {
			mostBananas = changesTotal
		}
	}

	fmt.Println("part2", mostBananas)
}
