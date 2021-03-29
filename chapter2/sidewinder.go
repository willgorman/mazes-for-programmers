package maze

import "math/rand"

func Sidewinder(g *Grid) *Grid {
	g.eachRow(func(r row) {
		run := []*Cell{}
		for _, c := range r {
			run = append(run, c)
			var east_edge, north_edge bool
			if c.east == nil {
				east_edge = true
			}
			if c.north == nil {
				north_edge = true
			}
			end_run := east_edge || (!north_edge && rand.Intn(2) == 0)

			if end_run {
				member := run[rand.Intn(len(run))]
				if member.north != nil {
					member.Link(member.north)
				}
				run = []*Cell{}
			} else {
				c.Link(c.east)
			}
		}
	})
	return g
}
