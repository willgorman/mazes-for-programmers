package maze

type coordinate struct {
	row, column int
}

type Cell struct {
	coordinate
	links                    map[*Cell]bool
	north, south, east, west *Cell
}

func NewCell(row, column int) *Cell {
	return &Cell{
		coordinate: coordinate{row: row, column: column},
		links:      make(map[*Cell]bool),
	}
}

func (c *Cell) Link(cell *Cell) *Cell {
	c.links[cell] = true
	cell.links[c] = true
	return c
}

func (c *Cell) Unlink(cell *Cell) *Cell {
	delete(c.links, cell)
	delete(cell.links, c)
	return c
}

func (c *Cell) Links() []*Cell {
	keys := make([]*Cell, len(c.links))
	i := 0
	for l := range c.links {
		keys[i] = l
		i++
	}
	return keys
}

func (c *Cell) Linked(cell *Cell) bool {
	return c.links[cell]
}

func (c *Cell) Neighbors() []*Cell {
	ns := []*Cell{}
	if c.north != nil {
		ns = append(ns, c.north)
	}
	if c.east != nil {
		ns = append(ns, c.east)
	}
	if c.south != nil {
		ns = append(ns, c.south)
	}
	if c.west != nil {
		ns = append(ns, c.west)
	}
	return ns
}
