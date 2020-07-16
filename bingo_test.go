package gobingo

import "testing"


func TestGrid_String(t *testing.T) {
	if baseGrid().String() != expectedGrid {
		t.Error("basegrid and expected grid are not equal")
		t.Error("basegrid")
		t.Error(baseGrid().String())
		t.Error("expected grid")
		t.Error(expectedGrid)
	}
}

const bound = 5
func baseGrid() Grid {
	grid := make([][]int, bound)
	c := 1
	for i := 0; i< bound; i++ {
		grid[i] = make([]int, bound)
		for j := 0; j < bound; j++ {
			grid[i][j] = c
			c++
		}
	}
	return Grid(grid)
}

var expectedGrid = ` 1  2  3  4  5
 6  7  8  9 10
11 12  F 14 15
16 17 18 19 20
21 22 23 24 25`