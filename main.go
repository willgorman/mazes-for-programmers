package main

import (
	"fmt"
	"math/rand"
	"time"

	maze "github.com/willgorman/mazes/chapter2"
)

func main() {
	maze.LongestPath()
	// grid := maze.NewGrid(10, 10)
	// maze.BinaryTree(grid)
	// fmt.Print(grid)
	// png := grid.ToPNG()
	// err := png.SavePNG("out.png")
	// if err != nil {
	// 	panic(err)
	// }
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func distance() {
	grid := maze.NewGrid(5, 5)
	maze.BinaryTree(grid)
	dgrid := maze.NewDistanceGrid(grid, 0, 0)
	// fmt.Println(grid.CellAt(0, 0).Distances())
	fmt.Println(dgrid)

	path := dgrid.Distances.PathTo(dgrid.CellAt(dgrid.Rows()-1, 0))
	dgrid.Distances = path
	fmt.Println(dgrid)
}
