package main

import (
	"math/rand"
	"time"

	maze "github.com/willgorman/mazes/chapter2"
)

func main() {
	grid := maze.NewGrid(10, 10)
	maze.BinaryTree(grid)
	// fmt.Print(grid)
	png := grid.ToPNG()
	err := png.SavePNG("out.png")
	if err != nil {
		panic(err)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
