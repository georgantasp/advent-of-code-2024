package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputtest
var inputtest string

//go:embed input
var input string

type registers struct {
	A int
	B int
	C int
}

type instruction struct {
	opcode  int
	operand int
}

func (i instruction) run(r *registers) (*int, *int) {
	switch i.opcode {
	case 0:
		//adv
		denom := 1
		for range i.comboOperand(r) {
			denom *= 2
		}
		r.A = r.A / denom
		return nil, nil
	case 1:
		//bxl
		r.B = r.B ^ i.operand
		return nil, nil
	case 2:
		//bst
		r.B = i.comboOperand(r) % 8
		return nil, nil
	case 3:
		//jnz
		if r.A != 0 {
			return &i.operand, nil
		}
		return nil, nil
	case 4:
		//bxc
		r.B = r.B ^ r.C
		return nil, nil
	case 5:
		//out
		out := i.comboOperand(r) % 8
		return nil, &out
	case 6:
		//bdv
		denom := 1
		for range i.comboOperand(r) {
			denom *= 2
		}
		r.B = r.A / denom
		return nil, nil
	case 7:
		//cdv
		denom := 1
		for range i.comboOperand(r) {
			denom *= 2
		}
		r.C = r.A / denom
		return nil, nil
	default:
		panic("should not reach here")
	}
}

func (i instruction) comboOperand(r *registers) int {
	switch i.operand {
	case 0, 1, 2, 3:
		return i.operand
	case 4:
		return r.A
	case 5:
		return r.B
	case 6:
		return r.C
	case 7:
		panic("reserved")
	default:
		panic("should not reach here")
	}
}

func main() {
	splitinput := strings.Split(input, "\n\n")
	registersInput := strings.Split(splitinput[0], "\n")

	r := registers{}
	r.A, _ = strconv.Atoi(strings.Split(registersInput[0], ": ")[1])
	r.B, _ = strconv.Atoi(strings.Split(registersInput[1], ": ")[1])
	r.C, _ = strconv.Atoi(strings.Split(registersInput[2], ": ")[1])

	programInput := strings.Split(strings.Split(splitinput[1], " ")[1], ",")
	program := make([]int, len(programInput))
	for i := range programInput {
		program[i], _ = strconv.Atoi(programInput[i])
	}

	instructionPointer := 0
	var allOut []int
	for {
		if instructionPointer >= len(program) {
			break
		}
		opcode := program[instructionPointer]
		operand := program[instructionPointer+1]
		jump, out := instruction{opcode, operand}.run(&r)
		if jump == nil {
			instructionPointer += 2
		} else {
			instructionPointer = *jump
		}
		if out != nil {
			allOut = append(allOut, *out)
		}
	}

	allOutS := make([]string, len(allOut))
	for i := range allOut {
		allOutS[i] = strconv.Itoa(allOut[i])
	}

	fmt.Println(strings.Join(allOutS, ","))
}
