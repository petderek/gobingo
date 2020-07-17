package gobingo

import (
	"testing"
)

func TestGrid_String(t *testing.T) {
	if baseGrid().String() != expectedGrid {
		t.Error("basegrid and expected grid are not equal")
		t.Error("basegrid")
		t.Error(baseGrid().String())
		t.Error("expected grid")
		t.Error(expectedGrid)
	}
}

func baseGrid() (grid Grid) {
	inc := uint64(1)
	for g := 0; g < 2; g++ {
		for i := 0; i < 16; i++ {
			if g == 0 && i == 12 {
				inc++
			}
			grid[g] <<= 4
			grid[g] |= inc
			inc++
			if inc > 5 {
				inc = 1
			}
		}
	}

	return grid
}

const expectedGrid = ` 1  2* 3  4  5*
16 17 18 19 20
31*32  F 34*35*
46 47 48*49 50
61 62*63 64*65
`

const (
	Fifteen = 0b1111
)

func TestNibbles_Left(t *testing.T) {
	nib := Nibbles(uint8(Fifteen << 4))

	if nib.Left() != Fifteen {
		t.Error(nib.Left())
	}
}

func TestNibbles_Right(t *testing.T) {
	nib := Nibbles(uint8(Fifteen))

	if nib.Right() != Fifteen {
		t.Error(nib.Right())
	}
}

const SampleGuid = "a7f7bf4c-a2eb-463f-8a3b-111d599ed869"

func TestGuidSymmetry(t *testing.T) {
	grid, err := ToGrid(SampleGuid)
	if err != nil {
		t.Fatal("Error running ToGrid: ", err)
	}

	guid := FromGrid(grid)

	if guid != SampleGuid {
		t.Errorf("expected equality:\n%s\n%s", guid, SampleGuid)
	}
}
