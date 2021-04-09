package maze

import "fmt"

func LongestPath() {
	grid := NewDistanceGrid(NewGrid(5, 5), 0, 0)
	BinaryTree(grid.Grid)
	start := grid.CellAt(0, 0)
	distances := start.Distances()
	newStart, _ := distances.Max()
	newDistances := newStart.Distances()
	goal, _ := newDistances.Max()

	grid.Distances = newDistances.PathTo(goal)
	fmt.Print(grid.String())
}
