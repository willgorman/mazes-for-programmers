package maze

import (
	"image/color"
	"math"
)

func ColorGrid(grid *Grid, distances *Distances) {
	_, max := distances.Max()
	grid.colorizer = &DistanceColorizer{distances, max}
}

type DistanceColorizer struct {
	distances *Distances
	max       int
}

func (dc *DistanceColorizer) SetDistances(d *Distances) {
	dc.distances = d
	_, dc.max = d.Max()
}

func (dc *DistanceColorizer) Background(cell *Cell) color.Color {
	distance, ok := dc.distances.cells[cell]
	if !ok {
		return nil
	}
	intensity := float64(dc.max-distance) / float64(dc.max)
	dark := math.Round(255 * intensity)
	bright := 128 + math.Round(127*intensity)
	return color.RGBA{uint8(bright), uint8(dark), uint8(dark), 255}
}
