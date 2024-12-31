package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input
var input string

//go:embed inputtest
var inputtest string

type connectionsMap map[string]map[string]struct{}

func main() {

	connectionsInput := strings.Split(input, "\n")

	connections := connectionsMap{}

	for _, conn := range connectionsInput {
		computers := strings.Split(conn, "-")

		c0 := computers[0]
		c1 := computers[1]

		conn0, ok0 := connections[c0]
		conn1, ok1 := connections[c1]

		if ok0 {
			conn0[c1] = struct{}{}
		} else {
			connections[c0] = map[string]struct{}{c1: {}}
		}

		if ok1 {
			conn1[c0] = struct{}{}
		} else {
			connections[c1] = map[string]struct{}{c0: {}}
		}
	}

	total := map[string]struct{}{}
	for c1, conn := range connections {
		if strings.HasPrefix(c1, "t") {
			for c2, _ := range conn {
				i := intersectMaps(conn, connections[c2])
				for c3, _ := range i {
					if c3 != c2 && c3 != c1 {
						possibleSet := []string{c1, c2, c3}
						slices.Sort(possibleSet)
						total[strings.Join(possibleSet, ",")] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Println("part1", len(total))

	biggestParty := []string{}
	for c1, _ := range connections {
		party := connections.findBiggestParty(map[string]struct{}{c1: {}})

		if len(party) > len(biggestParty) {
			biggestParty = []string{}
			for p, _ := range party {
				biggestParty = append(biggestParty, p)
			}
			slices.Sort(biggestParty)
		}
	}

	fmt.Println("part1", strings.Join(biggestParty, ","))
}

func (c connectionsMap) findBiggestParty(party map[string]struct{}) map[string]struct{} {
	biggestParty := party
	//for computer, _ := range c {
	//	if
	//}

	return biggestParty
}

func intersectMaps(map1, map2 map[string]struct{}) map[string]struct{} {
	result := map[string]struct{}{}

	for key, _ := range map1 {
		if _, ok := map2[key]; ok {
			result[key] = struct{}{}
		}
	}

	return result
}
