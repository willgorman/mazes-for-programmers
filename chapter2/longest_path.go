package maze

func LongestPath() {
	grid := NewDistanceGrid(NewGrid(25, 25), 0, 0)
	BinaryTree(grid.Grid)
	start := grid.CellAt(0, 0)
	distances := start.Distances()
	newStart, _ := distances.Max()
	newDistances := newStart.Distances()
	goal, _ := newDistances.Max()

	grid.Distances = newDistances.PathTo(goal)
	// fmt.Print(grid.String())

	err := grid.ToPNG().SavePNG("out.png")
	if err != nil {
		panic(err)
	}
}
