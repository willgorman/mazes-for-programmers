package main

import (
	"testing"

	maze "github.com/willgorman/mazes/chapter2"
)

func Benchmark_Aldous(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := maze.NewGrid(25, 25)
		maze.AldousBroder(grid)
	}

}

func Benchmark_Wilsons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := maze.NewGrid(25, 25)
		maze.Wilsons(grid)
	}

}

func Benchmark_HuntAndKill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := maze.NewGrid(25, 25)
		maze.HuntAndKill(grid)
	}

}
