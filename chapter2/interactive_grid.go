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
