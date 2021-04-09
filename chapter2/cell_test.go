package maze_test

import (
	"fmt"
	"testing"

	maze "github.com/willgorman/mazes/chapter2"
)

func TestDistances(t *testing.T) {
	grid := maze.NewGrid(4, 4)
	maze.BinaryTree(grid)
	fmt.Println(grid.RandomCell().Distances())
}
