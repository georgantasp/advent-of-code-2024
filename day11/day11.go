package main

import (
	_ "embed"
	"fmt"
	//_ "net/http/pprof"
	"strconv"
	"strings"
)

//go:embed inputtest
var inputtest string

//go:embed input
var input string

func main() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	intest := strings.Split(inputtest, " ")

	test := itterrate(intest, 6)
	fmt.Println("test answer is", test)

	in := strings.Split(input, " ")

	part1 := itterrate(in, 25)
	fmt.Println("Part 1 answer is", part1)
	part2 := itterrate(in, 75)
	fmt.Println("Part 2 answer is", part2)
}

func itterrate(in []string, times int) int {
	stones := make([]int, len(in))
	for i := range len(stones) {
		stones[i], _ = strconv.Atoi(in[i])
	}

	total := 0
	for _, s := range stones {
		n := getOrNewNode(s)
		total += n.blink(times)
	}
	return total
}

type node struct {
	stone int
	memos map[int]int
}

var memos = map[int]*node{}

func getOrNewNode(stone int) *node {
	if n, ok := memos[stone]; ok {
		return n
	}

	n := &node{stone: stone, memos: map[int]int{}}
	memos[stone] = n
	return n
}

func (n *node) blink(times int) int {
	if times == 0 {
		return 1
	}

	if res, ok := n.memos[times]; ok {
		return res
	}

	if n.stone == 0 {
		nextNode := getOrNewNode(1)
		val := nextNode.blink(times - 1)
		n.memos[times] = val
		return val
	}

	str := strconv.Itoa(n.stone)

	if len(str)%2 == 0 {
		half := len(str) / 2
		s1, _ := strconv.Atoi(str[:half])
		s2, _ := strconv.Atoi(str[half:])

		nextNode1 := getOrNewNode(s1)
		nextNode2 := getOrNewNode(s2)
		val := nextNode1.blink(times-1) + nextNode2.blink(times-1)
		n.memos[times] = val
		return val
	}

	nextNode := getOrNewNode(n.stone * 2024)
	val := nextNode.blink(times - 1)
	n.memos[times] = val
	return val
}
