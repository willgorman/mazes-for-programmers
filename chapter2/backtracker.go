package maze

import "math/rand"

type cellStack []*Cell

func (c *cellStack) push(cell *Cell) {
	*c = append(*c, cell)
}

func (c *cellStack) pop() *Cell {
	if len(*c) == 0 {
		return nil
	}
	poppee := len(*c) - 1
	retval := (*c)[poppee]
	*c = (*c)[:poppee]
	return retval
}

func (c *cellStack) peek() *Cell {
	if len(*c) == 0 {
		return nil
	}
	poppee := len(*c) - 1
	return (*c)[poppee]
}

func RecursiveBacktracker(g *Grid) *Grid {
	startAt := g.RandomCell()
	stack := cellStack{startAt}

	for len(stack) > 0 {
		current := stack.peek()
		neighbors := current.CollectNeighbors(withoutLinks)

		if len(neighbors) == 0 {
			stack.pop()
		} else {
			neighbor := neighbors[rand.Intn(len(neighbors))]
			current.Link(neighbor)
			stack.push(neighbor)
		}
	}
	return g
}
