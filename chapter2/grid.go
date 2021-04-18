package maze

import (
	"context"
	"image/color"
	"math/rand"
	"strconv"

	"github.com/fogleman/gg"
)

type field [][]*Cell

type row []*Cell

type rowVisitor func(row)

type cellVisitor func(*Cell)

type rowVisitorUntil func(context.Context, row)

type cellVisitorUntil func(context.Context, *Cell)

type Grid struct {
	field
	rows, columns int
	cellPrinter   CellContents
	colorizer     CellColorizer
}

type CellColorizer interface {
	Background(*Cell) color.Color
}

type defaultColorizer struct{}

func (c defaultColorizer) Background(*Cell) color.Color {
	return color.White
}

func NewGrid(rows, columns int) *Grid {
	g := Grid{rows: rows, columns: columns}
	g.field = g.prepare()
	g.configure()
	g.cellPrinter = g.defaultContents
	g.colorizer = defaultColorizer{}
	return &g
}

func (g *Grid) Rows() int {
	return g.rows
}

func (g *Grid) Columns() int {
	return g.columns
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
func (g *Grid) CellAt(row, col int) *Cell {
	if row < 0 || row > g.rows-1 {
		return nil
	}
	if col < 0 || col > g.columns-1 {
		return nil
	}

	return g.field[row][col]
}

func (g Grid) RandomCell() *Cell {
	row := rand.Intn(g.rows)
	col := rand.Intn(len(g.field[row]))

	return g.field[row][col]
}

func (g Grid) size() int {
	return g.rows * g.columns
}

func (g *Grid) eachRow(visit rowVisitor) {
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

func (g *Grid) eachRowUntil(ctx context.Context, visit rowVisitor) {
	for _, row := range g.field {
		visit(row)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func (g *Grid) eachCellUntil(ctx context.Context, visit cellVisitor) {
	g.eachRowUntil(ctx, func(r row) {
		for _, cell := range r {
			visit(cell)
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	})
}

type CellContents func(*Cell) string

func (g *Grid) defaultContents(cell *Cell) string {
	return " "
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
			body := " " + g.cellPrinter(cell) + " "
			if cell.Linked(cell.east) {
				top += body + " "
			} else {
				top += body + "|"
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
	for _, mode := range []string{"background", "wall"} {
		dc.SetColor(color.Black)
		g.eachCell(func(c *Cell) {
			x1 := c.column * cell_size
			y1 := c.row * cell_size
			x2 := (c.column + 1) * cell_size
			y2 := (c.row + 1) * cell_size
			if mode == "background" {
				cellColor := g.colorizer.Background(c)
				dc.SetColor(cellColor)
				dc.DrawRectangle(float64(x1), float64(y1), float64(x2), float64(y2))
				dc.Fill()
			} else {
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
			}
		})
	}
	dc.Fill()
	return dc
}

// TODO: return new grid rather than modifying
func NewDistanceGrid(g *Grid, startRow, startCol int) *DistanceGrid {
	distances := g.CellAt(startRow, startCol).Distances()

	dgrid := &DistanceGrid{
		Grid:      g,
		Distances: &distances,
	}
	g.cellPrinter = dgrid.cellPrinter // TODO: follow through
	// fmt.Println(distances)
	// TODO: interface or something less scope-juggly
	// g.cellPrinter = func(c *Cell) string {
	// 	if dgrid.Distances.GetDistance(c) != Unvisited {
	// 		return strconv.FormatInt(int64(dgrid.Distances.GetDistance(c)), 36)
	// 	} else {
	// 		return g.defaultContents(c)
	// 	}
	// }

	return dgrid
}

type DistanceGrid struct {
	*Grid
	Distances *Distances
}

func (d *DistanceGrid) String() string {
	// TODO: still feels like there's an interface here
	d.cellPrinter = func(c *Cell) string {
		if d.Distances.GetDistance(c) != Unvisited {
			return strconv.FormatInt(int64(d.Distances.GetDistance(c)), 36)
		} else {
			return d.Grid.defaultContents(c)
		}
	}

	return d.Grid.String()
}

func (d *DistanceGrid) ToPNG() *gg.Context {
	d.colorizer = d
	return d.Grid.ToPNG()
}

func (d *DistanceGrid) Background(c *Cell) color.Color {
	if d.Distances.GetDistance(c) != Unvisited {
		return color.RGBA{255, 100, 100, 255}
	} else {
		return defaultColorizer{}.Background(c)
	}
}
