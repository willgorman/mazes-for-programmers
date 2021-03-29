package maze

import "math/rand"

func BinaryTree(g *Grid) *Grid {
	g.eachCell(func(cell *Cell) {
		neighbors := []*Cell{}
		if cell.north != nil {
			neighbors = append(neighbors, cell.north)
		}
		if cell.east != nil {
			neighbors = append(neighbors, cell.east)
		}
		if len(neighbors) > 0 {
			n := neighbors[rand.Intn(len(neighbors))]
			if n != nil {
				cell.Link(n)
			}
		}
	})
	return g
}
