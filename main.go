package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	maze "github.com/willgorman/mazes/chapter2"
)

var (
	cmd = flag.String("command", "", "")
)

func main() {
	flag.Parse()
	switch *cmd {
	case "interactive":
		grid := maze.NewGrid(25, 25)
		maze.RecursiveBacktracker(grid)
		model := maze.TeaGrid{
			Grid: grid,
		}
		p := tea.NewProgram(&model)
		if err := p.Start(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	case "backtracker":
		grid := maze.NewGrid(25, 25)
		maze.RecursiveBacktracker(grid)
		png := grid.ToPNG()
		err := png.SavePNG("out.png")
		if err != nil {
			panic(err)
		}
	case "hunt-and-kill":
		grid := maze.NewGrid(25, 25)
		maze.HuntAndKill(grid)
		png := grid.ToPNG()
		err := png.SavePNG("out.png")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d dead-ends", len(grid.Deadends()))
	case "wilsons":
		grid := maze.NewGrid(25, 25)
		maze.Wilsons(grid)
		png := grid.ToPNG()
		err := png.SavePNG("out.png")
		if err != nil {
			panic(err)
		}
	case "aldousBroder":
		grid := maze.NewGrid(25, 25)
		maze.AldousBroder(grid)
		png := grid.ToPNG()
		err := png.SavePNG("out.png")
		if err != nil {
			panic(err)
		}
	case "longestPath":
		maze.LongestPath()
	case "coloredMaze":
		coloredMaze()
	case "distance":
		distance()
	case "topng":
		grid := maze.NewGrid(25, 25)
		maze.BinaryTree(grid)
		png := grid.ToPNG()
		err := png.SavePNG("out.png")
		if err != nil {
			panic(err)
		}
	default:
		grid := maze.NewGrid(25, 25)
		maze.BinaryTree(grid)
		fmt.Print(grid)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func distance() {
	grid := maze.NewGrid(10, 10)
	maze.Sidewinder(grid)
	dgrid := maze.NewDistanceGrid(grid, 0, 0)
	// fmt.Println(grid.CellAt(0, 0).Distances())
	// fmt.Println(dgrid)

	path := dgrid.Distances.PathTo(dgrid.CellAt(dgrid.Rows()-1, 0))
	dgrid.Distances = path
	fmt.Println(dgrid)
}

func coloredMaze() {
	grid := maze.NewGrid(25, 25)
	maze.AldousBroder(grid)

	start := grid.CellAt(0, 0)
	d := start.Distances()
	maze.ColorGrid(grid, &d)
	png := grid.ToPNG()
	err := png.SavePNG("color.png")
	if err != nil {
		panic(err)
	}
}
