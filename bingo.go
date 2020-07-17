package gobingo

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
)

const (
	sep   = '\n'
	space = ' '
	f     = 'F'

	EIGHT_MASK = 0b11111111
	LOW_MASK   = 0b00001111
)

// Grid stores bingo state in 128 bits
// bits 0-95: numbers (as nibbles)
// bits 96-121: state
type Grid [2]uint64

func (g Grid) ToNibbles() (n [16]Nibbles) {
	for i := 1; i >= 0; i-- {
		half := g[i]
		for j := 7; j >= 0; j-- {
			n[j+i*8] = Nibbles(half & EIGHT_MASK)
			half >>= 8
		}
	}
	return
}

// String plops out a displayable version of the board and its status
func (g Grid) String() string {

	nibs := g.ToNibbles()
	sb := strings.Builder{}

	// grab the 'state' and put it in an array of bools
	isSet := [24]bool{}
	for i := 0; i < 3; i++ {
		data := nibs[12+i]
		for j := 0; j < 8; j++ {
			isSet[j+i*8] = (data & (1 << j)) != 0
		}
	}

	skip := 0
	inc := 1
	for i := 0; i < 12; i++ {
		for j := 0; j < 2; j++ {
			var data uint8
			switch j {
			case 0:
				data = nibs[i].Left()
			case 1:
				data = nibs[i].Right()
			}

			addRow := ((inc + skip - 1) / 5) * 15

			number := int(data) + addRow
			numberString := strconv.Itoa(number)

			if number < 10 {
				sb.WriteByte(space)
			}

			sb.WriteString(numberString)

			if inc == 12 {
				sb.WriteByte(space)
				sb.WriteByte(space)
				sb.WriteByte(f)
				skip = 1
			}

			if inc > 0 && (inc+skip)%5 == 0 {
				if isSet[inc-1] {
					sb.WriteByte('*')
				}
				sb.WriteByte(sep)
			} else if isSet[inc-1] {
				sb.WriteByte('*')
			} else {
				sb.WriteByte(space)
			}
			inc++
		}
	}
	return sb.String()
}

// Nibbles is basically circumventing golang not having a uint4 by treating a uint8 as a tuple. There is probably a
// better way to do this. A Nibbles object has a left and a right nibble, which work the way you expect them too.
type Nibbles uint8

func (n Nibbles) Left() uint8 {
	return uint8(n) >> 4
}

func (n Nibbles) Right() uint8 {
	return uint8(n) & LOW_MASK
}

func ToGrid(guid string) (grid Grid, err error) {
	sanitized := strings.Replace(guid, "-", "", 4)
	data, err := hex.DecodeString(sanitized)
	if err != nil {
		return
	}

	if len(data) < 16 {
		err = errors.New("not enough data on string")
		return
	}

	grid[0] = binary.LittleEndian.Uint64(data[0:8])
	grid[1] = binary.LittleEndian.Uint64(data[8:])
	return
}

func FromGrid(grid Grid) (guid string) {
	data := make([]byte, 16)
	binary.LittleEndian.PutUint64(data[:8], grid[0])
	binary.LittleEndian.PutUint64(data[8:], grid[1])

	b := strings.Builder{}

	for i, c := range hex.EncodeToString(data) {
		switch i {
		case 8, 12, 16, 20:
			b.WriteByte('-')
			fallthrough
		default:
			b.WriteRune(c)
		}
	}

	return b.String()
}
