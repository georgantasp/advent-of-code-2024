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
	diskMapS := strings.Split(input, "")
	diskMap := make([]int, len(diskMapS))
	for i := range len(diskMapS) {
		diskMap[i], _ = strconv.Atoi(diskMapS[i])
	}

	pdm := printDiskMap(diskMap)

	part1(pdm)
	part2(diskMap)
}

func part1(pdm []*int) {
	for i := range len(pdm) {
		if pdm[i] == nil {
			for j := len(pdm) - 1; j > i; j-- {
				if pdm[j] != nil {
					pdm[i] = pdm[j]
					pdm[j] = nil
					break
				}
			}
		}
	}
	part1Total := calcSum(pdm)

	fmt.Println("par1", part1Total)
}

type diskUnit struct {
	isFile bool
	fileID int
	size   int
	moved  bool
}

func part2(diskMap []int) {
	ids := make([]diskUnit, len(diskMap))
	for i := range len(diskMap) {
		if i%2 == 0 {
			ids[i] = diskUnit{isFile: true, fileID: i / 2, size: diskMap[i]}
		} else {
			ids[i] = diskUnit{isFile: false, fileID: -1, size: diskMap[i]}
		}
	}

	for j := len(ids) - 1; j >= 0; j-- {
		fmt.Println(j)
		if !ids[j].isFile || ids[j].moved {
			continue
		}

		for i := range len(ids) {
			if ids[i].isFile {
				// not a space
				continue
			}
			if ids[i].size < ids[j].size {
				// space too small
				continue
			}
			if i > j {
				// don't move backwards
				break
			}

			// it's moving
			ids[j].moved = true
			// increase the size of the space before
			ids[j-1].size += ids[j].size

			start := ids[:i]
			// replace the space with empty space, the moved file, and remaining space
			mid := []diskUnit{
				{isFile: false, fileID: -1, size: 0},
				ids[j],
				{isFile: false, fileID: -1, size: ids[i].size - ids[j].size},
			}
			end1 := ids[i+1 : j]
			var end2 []diskUnit
			if j+1 < len(ids) {
				end2 = ids[j+1:]
			}

			var newIds []diskUnit
			newIds = append(newIds, start...)
			newIds = append(newIds, mid...)
			newIds = append(newIds, end1...)
			newIds = append(newIds, end2...)

			ids = newIds

			// there's a new "end"
			j += 2

			break
		}
	}

	part2Total := printDiskUnits(ids)

	fmt.Println("part2", part2Total)

	fmt.Println()
}

func printDiskMap(diskMap []int) []*int {
	printedDiskMap := []*int{}
	for i := range len(diskMap) {
		if i%2 == 0 {
			for range diskMap[i] {
				val := i / 2
				printedDiskMap = append(printedDiskMap, &val)
				fmt.Print(val)
			}
		} else {
			for range diskMap[i] {
				printedDiskMap = append(printedDiskMap, nil)
				fmt.Print(".")
			}
		}
	}
	fmt.Print("\n")
	return printedDiskMap
}

func printDiskUnits(ids []diskUnit) int {
	total := 0
	i := 0
	for _, du := range ids {
		for range du.size {
			if du.isFile {
				fmt.Print(du.fileID)
				total += i * du.fileID
			} else {
				fmt.Print(".")
			}
			i++
		}
	}
	fmt.Println()
	return total
}

func calcSum(pdm []*int) int {
	total := 0
	for i := range len(pdm) {
		if pdm[i] != nil {
			total += i * *pdm[i]
		}
	}
	return total
}
