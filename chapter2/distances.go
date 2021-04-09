package maze

import "fmt"

const (
	Unvisited = -1
)

type Distances struct {
	root  *Cell
	cells map[*Cell]int
}

func NewDistances(root *Cell) Distances {
	d := Distances{}
	d.root = root
	d.cells = map[*Cell]int{}
	d.cells[root] = 0
	return d
}

func (d Distances) GetDistance(c *Cell) int {
	distance, ok := d.cells[c]
	if !ok {
		return Unvisited
	}
	return distance
}

func (d Distances) SetDistance(c *Cell, dist int) {
	d.cells[c] = dist
}

func (d Distances) Cells() []*Cell {
	cells := make([]*Cell, len(d.cells))
	for cell := range d.cells {
		cells = append(cells, cell)
	}
	return cells
}

func (d Distances) String() string {
	output := ""
	for cell, distance := range d.cells {
		output += fmt.Sprintf("[%d, %d]: %d", cell.row, cell.column, distance)
		if cell.row == d.root.row && cell.column == d.root.column {
			output += " (root)"
		}
		output += "\n"
	}
	return output
}

func (d Distances) PathTo(goal *Cell) *Distances {
	current := goal

	breadcrumbs := NewDistances(d.root)
	breadcrumbs.SetDistance(current, d.cells[current])
	for current != d.root {
		for _, neighbor := range current.Links() {
			if d.cells[neighbor] < d.cells[current] {
				breadcrumbs.cells[neighbor] = d.cells[neighbor]
				current = neighbor
			}
		}
	}
	return &breadcrumbs
}

func (d Distances) Max() (maxCell *Cell, maxDistance int) {
	maxCell = d.root
	for cell, dist := range d.cells {
		if dist > maxDistance {
			maxCell = cell
			maxDistance = dist
		}
	}
	return
}
