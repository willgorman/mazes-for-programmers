package maze

import (
	"testing"

	gocmp "github.com/google/go-cmp/cmp"

	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
)

var (
	cell1 = &Cell{coordinate: coordinate{row: 1, column: 1}}
	cell2 = &Cell{coordinate: coordinate{row: 2, column: 2}}
)

func Test_cellStack_push(t *testing.T) {
	type args struct {
		cell *Cell
	}
	tests := []struct {
		name string
		c    cellStack
		args args
		want cellStack
	}{
		{
			name: "push 1",
			c:    nil,
			args: args{cell1},
			want: cellStack{cell1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.push(tt.args.cell)
			assert.Assert(t, cmp.DeepEqual(tt.c, tt.want, gocmp.AllowUnexported(Cell{}, coordinate{})))
		})
	}
}

func Test_cellStack_pop(t *testing.T) {
	tests := []struct {
		name string
		c    cellStack
		want *Cell
		post cellStack
	}{
		{
			name: "1 of 2",
			c:    cellStack{cell1, cell2},
			want: cell2,
			post: cellStack{cell1},
		},
		{
			name: "1 of 1",
			c:    cellStack{cell1},
			want: cell1,
			post: cellStack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.pop()
			assert.Assert(t, cmp.DeepEqual(got, tt.want, gocmp.AllowUnexported(Cell{}, coordinate{})))
			assert.Assert(t, cmp.DeepEqual(tt.c, tt.post, gocmp.AllowUnexported(Cell{}, coordinate{})))
		})
	}
}

func TestRecursiveBacktracker(t *testing.T) {
	type args struct {
		g *Grid
	}
	tests := []struct {
		name string
		args args
		want *Grid
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RecursiveBacktracker(tt.args.g)
			assert.Assert(t, cmp.DeepEqual(got, tt.want))
		})
	}
}
