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

func (g *garden) newRegion(row, col int) {
	r := &region{perimeter: 0}
	g.regions = append(g.regions, r)
	r.addPlot(g.plots[row][col])
	g.tryNeighbors(row, col)
}

func (g *garden) tryNeighbors(row, col int) {
	if row-1 >= 0 && g.plots[row][col].plant == g.plots[row-1][col].plant && g.plots[row-1][col].region == nil {
		g.plots[row][col].region.addPlot(g.plots[row-1][col])
		g.tryNeighbors(row-1, col)
	}
	if row+1 < len(g.plots) && g.plots[row][col].plant == g.plots[row+1][col].plant && g.plots[row+1][col].region == nil {
		g.plots[row][col].region.addPlot(g.plots[row+1][col])
		g.tryNeighbors(row+1, col)
	}
	if col-1 >= 0 && g.plots[row][col].plant == g.plots[row][col-1].plant && g.plots[row][col-1].region == nil {
		g.plots[row][col].region.addPlot(g.plots[row][col-1])
		g.tryNeighbors(row, col-1)
	}
	if col+1 < len(g.plots[0]) && g.plots[row][col].plant == g.plots[row][col+1].plant && g.plots[row][col+1].region == nil {
		g.plots[row][col].region.addPlot(g.plots[row][col+1])
		g.tryNeighbors(row, col+1)
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
}

func (r *region) addPlot(p *plot) {
	r.perimeter += r.addPerimeter(p)
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

	return count
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
}
