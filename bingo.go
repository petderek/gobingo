package gobingo

import (
	"strconv"
	"strings"
)

const (
	defaultWidth  = 5
	defaultHeight = 5
	sep           = '\n'
	space         = ' '
	f             = 'F'
)

type Grid [][]int

func (g Grid) String() string {
	sb := strings.Builder{}
	for i, r := range g {
		if i > 0 {
			sb.WriteByte(sep)
		}
		for j, c := range r {
			if j > 0 {
				sb.WriteByte(space)
			}

			// if middlemost, its a free space. so write F
			if i == 2 && j == 2 {
				sb.WriteByte(space)
				sb.WriteByte(f)
				continue
			}

			// write the actual character. prefix with a space if sub-10
			if c < 10 {
				sb.WriteByte(space)
			}
			sb.WriteString(strconv.Itoa(c))
		}
	}

	return sb.String()
}

// Bingo contains methods to generate bingo cards
type Bingo struct {
}

func (b *Bingo) GetGrid(guid string) Grid {
	return nil
}

// Returns a "random"-ish identifier
func (b *Bingo) NewIdentifier() string {

	return ""
}
