package maze

import (
	"context"
	"math/rand"
)

var withoutLinks = func(n *Cell) bool { return len(n.Links()) == 0 }
var withLinks = func(n *Cell) bool { return len(n.Links()) > 0 }

func HuntAndKill(g *Grid) *Grid {
	current := g.RandomCell()
	for current != nil {
		unvisitedNeighbors := current.CollectNeighbors(withoutLinks)
		if len(unvisitedNeighbors) > 0 {
			neighbor := unvisitedNeighbors[rand.Intn(len(unvisitedNeighbors))]
			current.Link(neighbor)
			current = neighbor
		} else {
			current = nil
			ctx, cancel := context.WithCancel(context.Background())
			g.eachCellUntil(ctx, func(c *Cell) {
				visitedNeighbors := c.CollectNeighbors(withLinks)
				if len(c.Links()) == 0 && len(visitedNeighbors) > 0 {
					current = c

					neighbor := visitedNeighbors[rand.Intn(len(visitedNeighbors))]
					current.Link(neighbor)
					cancel()
				}
			})
		}
	}
	return g
}
