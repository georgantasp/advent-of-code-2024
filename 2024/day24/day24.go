package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

type gate struct {
	in1       string
	operation string
	in2       string
	out       string
}

func main() {
	splitInput := strings.Split(input, "\n\n")
	startInputs := strings.Split(splitInput[0], "\n")

	computedWires := map[string]int{}
	for _, startInput := range startInputs {
		splitStartInput := strings.Split(startInput, ": ")
		bit, _ := strconv.Atoi(splitStartInput[1])
		computedWires[splitStartInput[0]] = bit
	}

	gateInputs := strings.Split(splitInput[1], "\n")
	gates := make([]gate, len(gateInputs))
	for i, gateInput := range gateInputs {
		splitGateInput := strings.Split(gateInput, " -> ")
		doubleSplitGateInput := strings.Split(splitGateInput[0], " ")

		if strings.HasPrefix(doubleSplitGateInput[0], "y") && strings.HasPrefix(doubleSplitGateInput[2], "x") {
			gates[i] = gate{
				in1:       doubleSplitGateInput[2],
				operation: doubleSplitGateInput[1],
				in2:       doubleSplitGateInput[0],
				out:       splitGateInput[1],
			}
		} else {
			gates[i] = gate{
				in1:       doubleSplitGateInput[0],
				operation: doubleSplitGateInput[1],
				in2:       doubleSplitGateInput[2],
				out:       splitGateInput[1],
			}
		}

	}
	zOutput := computeOut(gates, computedWires)
	fmt.Println("part1", zOutput)

	// https://www.electronics-tutorials.ws/combination/comb_7.html

	gatesMapByOut := map[string]gate{}
	for _, g := range gates {
		gatesMapByOut[g.out] = g
	}
	// jss <> rds
	tmp := gatesMapByOut["jss"]
	gatesMapByOut["jss"] = gate{
		in1:       gatesMapByOut["rds"].in1,
		operation: gatesMapByOut["rds"].operation,
		in2:       gatesMapByOut["rds"].in2,
		out:       "jss",
	}
	gatesMapByOut["rds"] = gate{
		in1:       tmp.in1,
		operation: tmp.operation,
		in2:       tmp.in2,
		out:       "rds",
	}
	// wss <> z18
	tmp = gatesMapByOut["wss"]
	gatesMapByOut["wss"] = gate{
		in1:       gatesMapByOut["z18"].in1,
		operation: gatesMapByOut["z18"].operation,
		in2:       gatesMapByOut["z18"].in2,
		out:       "wss",
	}
	gatesMapByOut["z18"] = gate{
		in1:       tmp.in1,
		operation: tmp.operation,
		in2:       tmp.in2,
		out:       "z18",
	}
	// bmn <> z23
	tmp = gatesMapByOut["bmn"]
	gatesMapByOut["bmn"] = gate{
		in1:       gatesMapByOut["z23"].in1,
		operation: gatesMapByOut["z23"].operation,
		in2:       gatesMapByOut["z23"].in2,
		out:       "bmn",
	}
	gatesMapByOut["z23"] = gate{
		in1:       tmp.in1,
		operation: tmp.operation,
		in2:       tmp.in2,
		out:       "z23",
	}
	// mvb <> z08
	tmp = gatesMapByOut["mvb"]
	gatesMapByOut["mvb"] = gate{
		in1:       gatesMapByOut["z08"].in1,
		operation: gatesMapByOut["z08"].operation,
		in2:       gatesMapByOut["z08"].in2,
		out:       "mvb",
	}
	gatesMapByOut["z08"] = gate{
		in1:       tmp.in1,
		operation: tmp.operation,
		in2:       tmp.in2,
		out:       "z08",
	}

	gatesMapByIn := map[string][]gate{}
	for _, g := range gatesMapByOut {
		gatesMapByIn[g.in1] = append(gatesMapByIn[g.in1], g)
		gatesMapByIn[g.in2] = append(gatesMapByIn[g.in2], g)

		slices.SortFunc(gatesMapByIn[g.in1], func(a, b gate) int {
			return cmp.Compare(a.operation, b.operation)
		})
		slices.SortFunc(gatesMapByIn[g.in2], func(a, b gate) int {
			return cmp.Compare(a.operation, b.operation)
		})
	}

	zOutput = 0
	var carry gate
	for i := range 45 {
		sum := gatesMapByOut[fmt.Sprintf("z%02d", i)]
		x := gatesMapByIn[fmt.Sprintf("x%02d", i)]

		var xor *gate
		var and *gate
		for _, g2 := range x {
			if g2.operation == "XOR" {
				xor = &g2
			}
			if g2.operation == "AND" {
				and = &g2
			}
		}
		if xor == nil || and == nil {
			fmt.Println("ISSUE x and y must both be XORed and ANDed")
		}
		if !(i == 0 || (sum.in1 == xor.out && sum.in2 == carry.out) || (sum.in2 == xor.out && sum.in1 == carry.out)) {
			fmt.Println("ISSUE sum inputs incorrect")
		}

		if sum.operation != "XOR" {
			fmt.Println("ISSUE sum operation should be XOR")
		}

		if i == 0 {
			carry = *and
		} else if len(gatesMapByIn[and.out]) != 1 {
			fmt.Println("ISSUE AND out should be input once")
			carry = gate{}
		} else {
			carry = gatesMapByIn[and.out][0]
		}
	}

	part2Out := []string{"jss", "rds", "wss", "z18", "bmn", "z23", "mvb", "z08"}
	slices.Sort(part2Out)

	fmt.Println("part2", strings.Join(part2Out, ","))
}

func zeroWires() map[string]int {
	m := map[string]int{}
	for i := range 45 {
		m[fmt.Sprintf("x%02d", i)] = 0
		m[fmt.Sprintf("y%02d", i)] = 0
	}
	return m
}

func computeOut(gates []gate, computedWires map[string]int) int {
	remaininggates := gates
	for {
		if len(remaininggates) < 1 {
			break
		}
		newRemainingGates := []gate{}
		for _, g := range gates {
			w1, ok1 := computedWires[g.in1]
			w2, ok2 := computedWires[g.in2]

			if !ok1 || !ok2 {
				newRemainingGates = append(remaininggates, g)
				continue
			}

			computedWires[g.out] = compute(w1, w2, g.operation)
		}
		remaininggates = newRemainingGates
	}

	zOutput := 0
	for w, o := range computedWires {
		if strings.HasPrefix(w, "z") {
			n, _ := strconv.Atoi(w[1:])
			if o == 1 {
				zOutput += powInt(2, n)
			}
		}
	}
	return zOutput
}

func compute(w1 int, w2 int, operation string) int {
	switch operation {
	case "AND":
		if w1+w2 == 2 {
			return 1
		} else {
			return 0
		}
	case "OR":
		if w1+w2 > 0 {
			return 1
		} else {
			return 0
		}
	case "XOR":
		if w1+w2 == 1 {
			return 1
		} else {
			return 0
		}
	default:
		panic("unknown operation")
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
