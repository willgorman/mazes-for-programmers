package maze

import (
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type field [][]*Cell

type row []*Cell

type rowVistor func(row)

type cellVisitor func(*Cell)

type Grid struct {
	field
	rows, columns int
}

func NewGrid(rows, columns int) *Grid {
	g := Grid{rows: rows, columns: columns}
	g.field = g.prepare()
	g.configure()
	return &g
}

// will need to override later
func (g *Grid) configure() {
	for row, r := range g.field {
		for column, cell := range r {
			if row > 0 {
				cell.north = g.field[row-1][column]
			}
			if row < g.rows-1 {
				cell.south = g.field[row+1][column]
			}
			if column > 0 {
				cell.west = g.field[row][column-1]
			}
			if column < g.columns-1 {
				cell.east = g.field[row][column+1]
			}
		}
	}
}

// will need to override later
func (g *Grid) prepare() field {
	field := make(field, g.rows)
	for r := 0; r < g.rows; r++ {
		field[r] = make([]*Cell, g.columns)
		for c := 0; c < g.columns; c++ {
			field[r][c] = NewCell(r, c)
		}
	}
	return field
}

// will need to override
func (g *Grid) cellAt(row, col int) *Cell {
	if row < 0 || row > g.rows-1 {
		return nil
	}
	if col < 0 || col > g.columns-1 {
		return nil
	}

	return g.field[row][col]
}

func (g Grid) randomCell() *Cell {
	row := rand.Intn(g.rows)
	col := rand.Intn(len(g.field[row]))

	return g.field[row][col]
}

func (g Grid) size() int {
	return g.rows * g.columns
}

func (g *Grid) eachRow(visit rowVistor) {
	for _, row := range g.field {
		visit(row)
	}
}

func (g *Grid) eachCell(visit cellVisitor) {
	g.eachRow(func(r row) {
		for _, cell := range r {
			visit(cell)
		}
	})
}

func (g *Grid) String() string {
	output := "+"
	for i := 0; i < g.columns; i++ {
		output += "---+"
	}
	output += "\n"
	g.eachRow(func(r row) {
		top := "|"
		bottom := "+"
		for _, c := range r {
			cell := NewCell(-1, -1)
			if c != nil {
				cell = c
			}
			if cell.Linked(cell.east) {
				top += "    "
			} else {
				top += "   |"
			}
			if cell.Linked(cell.south) {
				bottom += "   +"
			} else {
				bottom += "---+"
			}
		}
		output = output + top + "\n"
		output = output + bottom + "\n"
	})
	return output
}

func (g *Grid) ToPNG() *gg.Context {
	cell_size := 100
	dc := gg.NewContext(cell_size*g.rows, cell_size*g.columns)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetColor(color.Black)
	g.eachCell(func(c *Cell) {
		x1 := c.column * cell_size
		y1 := c.row * cell_size
		x2 := (c.column + 1) * cell_size
		y2 := (c.row + 1) * cell_size

		if c.north == nil {
			dc.DrawLine(float64(x1), float64(y1), float64(x2), float64(y1))
			dc.Stroke()
		}
		if c.west == nil {
			dc.DrawLine(float64(x1), float64(y1), float64(x1), float64(y2))
			dc.Stroke()
		}

		if !c.Linked(c.east) {
			dc.DrawLine(float64(x2), float64(y1), float64(x2), float64(y2))
			dc.Stroke()
		}
		if !c.Linked(c.south) {
			dc.DrawLine(float64(x1), float64(y2), float64(x2), float64(y2))
			dc.Stroke()
		}

	})
	dc.Fill()
	return dc
}