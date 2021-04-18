package maze

import (
	"math/rand"
)

// FIXME: infinite loops
func Wilsons(g *Grid) *Grid {
	unvisited := []*Cell{}
	g.eachCell(func(c *Cell) { unvisited = append(unvisited, c) })

	first := unvisited[rand.Intn(len(unvisited))]
	unvisited = filter(first, unvisited)

	for len(unvisited) > 0 {
		cell := unvisited[rand.Intn(len(unvisited))]
		path := []*Cell{cell}
		for includes(cell, unvisited) {
			cell = cell.Neighbors()[rand.Intn(len(cell.Neighbors()))]
			position := index(cell, path)
			if position >= 0 {
				path = path[:position]
			} else {
				path = append(path, cell)
			}
		}
		for i := 0; i <= len(path)-2; i++ {
			path[i].Link(path[i+1])
			unvisited = filter(path[i], unvisited)
		}
	}

	return g
}

func filter(cell *Cell, cells []*Cell) []*Cell {
	n := 0
	for _, x := range cells {
		if x != cell {
			cells[n] = x
			n++
		}
	}
	return cells[:n]
}

func index(cell *Cell, cells []*Cell) int {
	for i, v := range cells {
		if v == cell {
			return i
		}
	}
	return -1
}

func includes(cell *Cell, cells []*Cell) bool {
	for _, v := range cells {
		if v == cell {
			return true
		}
	}
	return false
}
