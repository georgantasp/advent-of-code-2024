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

type connectionsByComputerMap map[string]map[string]struct{}
type connectionsMap map[connection]struct{}

type connection struct {
	computer0 string
	computer1 string
}

type party map[string]struct{}

func NewConnection(computer0 string, computer1 string) connection {
	if computer0 < computer1 {
		return connection{computer0, computer1}
	}
	return connection{computer1, computer0}
}

func main() {

	connectionsInput := strings.Split(input, "\n")

	connectionsByComputer := connectionsByComputerMap{}
	connections := connectionsMap{}

	for _, conn := range connectionsInput {
		computers := strings.Split(conn, "-")

		c0 := computers[0]
		c1 := computers[1]
		connections[NewConnection(c0, c1)] = struct{}{}

		conn0, ok0 := connectionsByComputer[c0]
		conn1, ok1 := connectionsByComputer[c1]

		if ok0 {
			conn0[c1] = struct{}{}
		} else {
			connectionsByComputer[c0] = map[string]struct{}{c1: {}}
		}

		if ok1 {
			conn1[c0] = struct{}{}
		} else {
			connectionsByComputer[c1] = map[string]struct{}{c0: {}}
		}
	}

	total := map[string]struct{}{}
	for c1, conn := range connectionsByComputer {
		if strings.HasPrefix(c1, "t") {
			for c2, _ := range conn {
				i := intersectMaps(conn, connectionsByComputer[c2])
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

	parties := map[string][]party{}

	for computer0, conns := range connectionsByComputer {
		parties[computer0] = []party{}

		for computer1, _ := range conns {
			var found bool
			for _, testParty := range parties[computer0] {
				canBelongInParty := true
				for computer2, _ := range testParty {
					if _, ok := connections[NewConnection(computer1, computer2)]; !ok {
						canBelongInParty = false
						break
					}
				}
				if canBelongInParty {
					found = true
					testParty[computer1] = struct{}{}
				}
			}
			if !found {
				parties[computer0] = append(parties[computer0], party{computer0: struct{}{}, computer1: struct{}{}})
			}
		}
	}

	var biggestParty []string
	for _, computerPartires := range parties {
		for _, p := range computerPartires {
			if len(p) > len(biggestParty) {
				biggestParty = make([]string, 0, len(p))
				for c, _ := range p {
					biggestParty = append(biggestParty, c)
				}
				slices.Sort(biggestParty)
			}
		}
	}

	fmt.Println("part2", strings.Join(biggestParty, ","))
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

func all(fns ...func() bool) bool {
	for _, f := range fns {
		if !f() {
			return false
		}
	}
	return true
}
