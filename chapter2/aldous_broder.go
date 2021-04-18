package maze

import "math/rand"

func AldousBroder(g *Grid) *Grid {
	cell := g.RandomCell()
	unvisited := g.size() - 1

	for unvisited > 0 {
		neighbors := cell.Neighbors()
		neighbor := neighbors[rand.Intn(len(neighbors))]

		if len(neighbor.links) == 0 {
			cell.Link(neighbor)
			unvisited--
		}

		cell = neighbor
	}

	return g
}
