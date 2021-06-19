package maze

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TeaGrid struct {
	*Grid
	active *Cell
}

func (g *TeaGrid) Init() tea.Cmd {
	g.active = g.RandomCell()
	return nil
}

func (g *TeaGrid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return g, tea.Quit
		case "w", "up":
			if g.active.Linked(g.active.north) {
				g.active = g.active.north
			}
			return g, nil
		case "a", "left":
			if g.active.Linked(g.active.west) {
				g.active = g.active.west
			}
			return g, nil
		case "s", "down":
			if g.active.Linked(g.active.south) {
				g.active = g.active.south
			}
			return g, nil
		case "d", "right":
			if g.active.Linked(g.active.east) {
				g.active = g.active.east
			}
			return g, nil
		default:
			g.active = g.RandomCell()
		}
	}
	return g, nil
}

func (g *TeaGrid) View() string {
	// TODO: render as something other than basic ascii
	g.cellPrinter = func(c *Cell) string {
		if c == g.active {
			return "@"
		} else {
			return g.defaultContents(c)
		}
	}
	return g.String()
}
