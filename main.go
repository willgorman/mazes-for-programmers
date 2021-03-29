package main

import (
	"fmt"
	"math/rand"
	"time"

	maze "github.com/willgorman/mazes/chapter2"
)

func main() {
	grid := maze.NewGrid(10, 10)
	maze.Sidewinder(grid)
	fmt.Print(grid)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
