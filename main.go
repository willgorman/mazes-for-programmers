package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	maze "github.com/willgorman/mazes/chapter2"
)

var (
	cmd = flag.String("command", "", "")
)

func main() {
	flag.Parse()
	switch *cmd {
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
	maze.Sidewinder(grid)

	start := grid.CellAt(0, 0)
	d := start.Distances()
	maze.ColorGrid(grid, &d)
	png := grid.ToPNG()
	err := png.SavePNG("color.png")
	if err != nil {
		panic(err)
	}
}
