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

func (r *registers) run(opcode, operand int) (*int, *int) {
	switch opcode {
	case 0:
		//adv
		denom := 1
		for range r.comboOperand(operand) {
			denom *= 2
		}
		r.A = r.A / denom
		return nil, nil
	case 1:
		//bxl
		r.B = r.B ^ operand
		return nil, nil
	case 2:
		//bst
		r.B = r.comboOperand(operand) % 8
		return nil, nil
	case 3:
		//jnz
		if r.A != 0 {
			return &operand, nil
		}
		return nil, nil
	case 4:
		//bxc
		r.B = r.B ^ r.C
		return nil, nil
	case 5:
		//out
		out := r.comboOperand(operand) % 8
		return nil, &out
	case 6:
		//bdv
		denom := 1
		for range r.comboOperand(operand) {
			denom *= 2
		}
		r.B = r.A / denom
		return nil, nil
	case 7:
		//cdv
		denom := 1
		for range r.comboOperand(operand) {
			denom *= 2
		}
		r.C = r.A / denom
		return nil, nil
	default:
		panic("should not reach here")
	}
}

func (r *registers) comboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
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

	r := &registers{}
	r.A, _ = strconv.Atoi(strings.Split(registersInput[0], ": ")[1])
	r.B, _ = strconv.Atoi(strings.Split(registersInput[1], ": ")[1])
	r.C, _ = strconv.Atoi(strings.Split(registersInput[2], ": ")[1])

	programInput := strings.Split(strings.Split(splitinput[1], " ")[1], ",")
	program := make([]int, len(programInput))
	for i := range programInput {
		program[i], _ = strconv.Atoi(programInput[i])
	}

	part1(r, program)
	part2(program)
}

func run(r *registers, program []int) []int {
	instructionPointer := 0
	var allOut []int
	for {
		if instructionPointer >= len(program) {
			break
		}
		opcode := program[instructionPointer]
		operand := program[instructionPointer+1]
		jump, out := r.run(opcode, operand)
		if jump == nil {
			instructionPointer += 2
		} else {
			instructionPointer = *jump
		}
		if out != nil {
			allOut = append(allOut, *out)
		}
	}
	return allOut
}

func part1(r *registers, program []int) {
	allOut := run(r, program)
	printProgram(allOut)
}

func part2(program []int) {
	printProgram(program)
	var r *registers

	m := make(map[int]int)
	j := 0
	for i := 15; i >= 0; i-- {
		for ; j <= 7; j++ {
			//fmt.Println(i, j)
			regA := 0
			regA = regA ^ (j << (3 * i))
			for k, v := range m {
				regA = regA ^ (v << (3 * k))
			}
			if regA == 0 {
				continue
			}

			r = &registers{
				A: regA,
			}

			out := run(r, program)
			if out[i] == program[i] {
				m[i] = j
				break
			}

			//r = &registers{
			//	A: 5<<(3*15) ^ // 0
			//		6<<(3*14) ^ //3
			//		0<<(3*13) ^ //5
			//		0<<(3*12) ^ //5
			//		0<<(3*11) ^ //5
			//		0<<(3*10) ^ //4
			//		0<<(3*9) ^ //4
			//		0<<(3*8) ^ //1
			//		0<<(3*7) ^ //3
			//		0<<(3*6) ^ //0
			//		0<<(3*5) ^ //5
			//		0<<(3*4) ^ //7
			//		0<<(3*3) ^ //1
			//		0<<(3*2) ^ //1
			//		0<<(3*1) ^ //4
			//		0<<(3*0), //2
			//}
		}
		if _, ok := m[i]; !ok {
			j = m[i+1] + 1
			delete(m, i+1)
			i += 2
		} else {
			j = 0
		}
	}
	regA := 0
	for k, v := range m {
		regA = regA ^ (v << (3 * k))
	}
	fmt.Println(regA)
	printProgram(run(&registers{A: regA}, program))
}

func programsEqual(p1, p2 []int) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i, v := range p1 {
		if v != p2[i] {
			return false
		}
	}
	return true
}

func printProgram(program []int) {
	for i := range program {
		fmt.Print(program[i])
		if i == len(program)-1 {
			fmt.Println()
		} else {
			fmt.Print(",")
		}
	}
}

func part2Brute(program []int) {
	regA := 0
regLoop:
	for {
		r := &registers{A: regA}
		instructionPointer := 0
		var allOut []int
		for {
			if instructionPointer >= len(program) {
				break
			}
			opcode := program[instructionPointer]
			operand := program[instructionPointer+1]
			jump, out := r.run(opcode, operand)
			if jump == nil {
				instructionPointer += 2
			} else {
				instructionPointer = *jump
			}
			if out != nil {
				allOut = append(allOut, *out)
				if len(allOut) > len(program) || allOut[len(allOut)-1] != program[len(allOut)-1] {
					if len(allOut) > 7 {
						fmt.Println(regA, allOut[:8])
					}
					break
				}
			}
		}

		if len(allOut) != len(program) {
			regA++
			continue regLoop
		}
		for i, v := range program {
			if v != allOut[i] {
				regA++
				continue regLoop
			}
		}

		break
	}

	fmt.Println("part2", regA)
}
