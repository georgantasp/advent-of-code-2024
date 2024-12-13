package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type garden struct {
	plots   [][]*plot
	regions []*region
}

func (g *garden) perimeterScore() int {
	total := 0

	for _, r := range g.regions {
		total += len(r.plots) * r.perimeter
	}

	return total
}

func (g *garden) sidesScore() int {
	total := 0

	for _, r := range g.regions {
		total += len(r.plots) * r.sides
	}

	return total
}

func (g *garden) newRegion(row, col int) {
	r := &region{perimeter: 0}
	g.regions = append(g.regions, r)
	r.addPlot(g.plots[row][col])
	g.tryNeighbors(row, col)
}

func (g *garden) tryNeighbors(row, col int) {

	// right
	if col+1 < len(g.plots[0]) && g.plots[row][col].plant == g.plots[row][col+1].plant && g.plots[row][col+1].region == nil {
		g.plots[row][col].region.addPlot(g.plots[row][col+1])
		g.tryNeighbors(row, col+1)
	}
	// below
	if row+1 < len(g.plots) && g.plots[row][col].plant == g.plots[row+1][col].plant && g.plots[row+1][col].region == nil {
		g.plots[row][col].region.addPlot(g.plots[row+1][col])
		g.tryNeighbors(row+1, col)
	}
	// left
	if col-1 >= 0 && g.plots[row][col].plant == g.plots[row][col-1].plant && g.plots[row][col-1].region == nil {
		g.plots[row][col].region.addPlot(g.plots[row][col-1])
		g.tryNeighbors(row, col-1)
	}
	// above
	if row-1 >= 0 && g.plots[row][col].plant == g.plots[row-1][col].plant && g.plots[row-1][col].region == nil {
		g.plots[row][col].region.addPlot(g.plots[row-1][col])
		g.tryNeighbors(row-1, col)
	}
}

type plot struct {
	plant string
	row   int
	col   int

	region *region
}

type region struct {
	plots     []*plot
	perimeter int
	sides     int
}

func (r *region) addPlot(p *plot) {
	r.perimeter += r.addPerimeter(p)
	r.sides += r.addSides(p)
	r.plots = append(r.plots, p)
	p.region = r
}

func (r *region) addPerimeter(p *plot) int {
	count := 0
	for _, np := range r.plots {
		if p.row-1 == np.row && p.col == np.col {
			count++
		}
		if p.row+1 == np.row && p.col == np.col {
			count++
		}
		if p.row == np.row && p.col-1 == np.col {
			count++
		}
		if p.row == np.row && p.col+1 == np.col {
			count++
		}
	}

	switch count {
	case 0:
		return 4
	case 1:
		return 2
	case 2:
		return 0
	case 3:
		return -2
	case 4:
		return -4
	}

	panic("should not reach here")
}

func (r *region) addSides(p *plot) int {

	fmt.Println("plant", p.plant, p.row, p.col, r.sides)

	var above *plot
	var aboveleft *plot
	var aboveright *plot
	var below *plot
	var belowleft *plot
	var belowright *plot
	var left *plot
	var right *plot

	for _, np := range r.plots {
		if p.row-1 == np.row && p.col == np.col {
			// above
			above = np
			for _, cn := range r.plots {
				if p.row-1 == cn.row && p.col-1 == cn.col {
					aboveleft = cn
				}
				if p.row-1 == cn.row && p.col+1 == cn.col {
					aboveright = cn
				}
			}
		}
		if p.row+1 == np.row && p.col == np.col {
			// below
			below = np
			for _, cn := range r.plots {
				if p.row+1 == cn.row && p.col-1 == cn.col {
					belowleft = cn
				}
				if p.row+1 == cn.row && p.col+1 == cn.col {
					belowright = cn
				}
			}
		}
		if p.row == np.row && p.col-1 == np.col {
			// left
			left = np
			for _, cn := range r.plots {
				if p.row-1 == cn.row && p.col-1 == cn.col {
					aboveleft = cn
				}
				if p.row+1 == cn.row && p.col-1 == cn.col {
					belowleft = cn
				}
			}
		}
		if p.row == np.row && p.col+1 == np.col {
			// right
			right = np
			for _, cn := range r.plots {
				if p.row-1 == cn.row && p.col+1 == cn.col {
					aboveright = cn
				}
				if p.row+1 == cn.row && p.col+1 == cn.col {
					belowright = cn
				}
			}
		}
	}

	sumPlots := func(b ...*plot) int {
		total := 0
		for _, v := range b {
			if v != nil {
				total++
			}
		}
		return total
	}

	switch sumPlots(left, above, right, below) {
	case 0:
		return 4
	case 4:
		return -4
	case 1:
		switch sumPlots(aboveleft, aboveright, belowleft, belowright) {
		case 0:
			return 0
		case 1:
			return 2
		case 2:
			return 4
		default:
			panic("should not reach here")
		}
	case 2:
		switch sumPlots(aboveleft, aboveright, belowleft, belowright) {
		case 0:
			//if sumPlots(above, below) == 2 || sumPlots(left, right) == 2 {
			//	return -4
			//}
			//return -2
			panic("doesn't reach here")
		case 1:
			switch {
			case sumPlots(above, below) == 2,
				sumPlots(left, right) == 2,
				sumPlots(left, above, aboveleft) == 3,
				sumPlots(above, right, aboveright) == 3,
				sumPlots(right, below, belowright) == 3,
				sumPlots(below, left, belowleft) == 3:
				return -2
			default:
				return 0
			}
		case 2:
			switch {
			case sumPlots(above, below) == 2,
				sumPlots(left, right) == 2,
				sumPlots(left, above, aboveleft) == 3,
				sumPlots(above, right, aboveright) == 3,
				sumPlots(right, below, belowright) == 3,
				sumPlots(below, left, belowleft) == 3:
				return 0
			default:
				//return 2
				panic("doesn't reach here")
			}
		case 3:
			return 2
		case 4:
			return 4
		default:
			panic("should not reach here")
		}
	case 3:
		switch sumPlots(aboveleft, aboveright, belowleft, belowright) {
		case 0, 1:
			panic("doesn't reach here")
		case 2:
			switch {
			case sumPlots(left, aboveleft, belowleft) == 0,
				sumPlots(above, aboveleft, aboveright) == 0,
				sumPlots(right, aboveright, belowright) == 0,
				sumPlots(below, belowleft, belowright) == 0:
				return -4
			case sumPlots(left, aboveright, belowright) == 0,
				sumPlots(above, belowleft, belowright) == 0,
				sumPlots(right, aboveleft, belowleft) == 0,
				sumPlots(below, aboveleft, aboveright) == 0:
				return -4
			default:
				return -2
			}
		case 3:
			switch {
			case sumPlots(above, aboveleft) == 0,
				sumPlots(above, aboveright) == 0,
				sumPlots(left, aboveleft) == 0,
				sumPlots(left, belowleft) == 0,
				sumPlots(right, aboveright) == 0,
				sumPlots(right, belowright) == 0,
				sumPlots(below, belowleft) == 0,
				sumPlots(below, belowright) == 0:
				return -2
			default:
				return 0
			}
		case 4:
			return 0
		default:
			panic("should not reach here")
		}
	default:
		panic("should not reach here")
	}
}

func main() {

	lines := strings.Split(input, "\n")

	g := garden{
		plots: make([][]*plot, len(lines)),
	}
	for i, l := range lines {
		plotCols := strings.Split(l, "")
		g.plots[i] = make([]*plot, len(plotCols))
		for j, p := range plotCols {
			g.plots[i][j] = &plot{plant: p, row: i, col: j}
		}
	}

	for i := range g.plots {
		for j := range g.plots[i] {
			if g.plots[i][j].region == nil {
				g.newRegion(i, j)
			}
		}
	}

	fmt.Println("Part 1:", g.perimeterScore())
	fmt.Println("Part 2:", g.sidesScore())
}
