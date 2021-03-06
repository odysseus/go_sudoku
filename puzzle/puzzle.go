package puzzle

import (
	"fmt"
)

/// Puzzle Struct ///
type Puzzle struct {
	Board []int
}

func NewPuzzle() *Puzzle {
	p := Puzzle{Board: make([]int, 81)}
	return &p
}

func (p *Puzzle) String() string {
	s := ""
	for i, v := range p.Board {
		row := rowForIndex(i)
		col := columnForIndex(i)
		if i%9 == 0 {
			s += "\n"
			if row != 0 && row%3 == 0 {
				s += "----------------------\n"
			}
		}
		if col != 0 && col%3 == 0 {
			s += "| "
		}
		s += fmt.Sprintf("%v ", v)
	}
	s += "\n"
	return s
}

func (p *Puzzle) CandidatesForIndex(ind int) []int {
	bools := make([]bool, 10)
	ci := CheckIndices(ind)
	for _, v := range ci {
		bools[p.Board[v]] = true
	}
	cands := make([]int, 0, 9)
	for i := 1; i < len(bools); i++ {
		if !bools[i] {
			cands = append(cands, i)
		}
	}
	return cands
}

/// Index finding methods ///

func rowForIndex(ind int) int {
	return ind / 9
}

func columnForIndex(ind int) int {
	return ind % 9
}

func houseForIndex(ind int) int {
	return ((rowForIndex(ind) / 3) * 3) + (columnForIndex(ind) / 3)
}

func indexForRowAndCol(row, col int) int {
	return row*9 + col
}

func indicesForRow(row int) []int {
	indices := make([]int, 9)
	start := row * 9
	for i := 0; i < 9; i++ {
		indices[i] = start + i
	}
	return indices
}

func indicesForColumn(col int) []int {
	indices := make([]int, 9)
	start := col
	for i := 0; i < 9; i++ {
		indices[i] = start
		start += 9
	}
	return indices
}

func indicesForHouse(house int) []int {
	indices := make([]int, 0, 9)
	houseCorners := []int{0, 3, 6, 27, 30, 33, 54, 57, 60}
	corner := houseCorners[house]
	for i := 0; i < 3; i++ {
		current := corner + (9 * i)
		for j := 0; j < 3; j++ {
			indices = append(indices, (current + j))
		}
	}
	return indices
}

func CheckIndices(ind int) []int {
	checkIndices := indicesForRow(rowForIndex(ind))
	checkIndices = append(checkIndices, indicesForColumn(columnForIndex(ind))...)
	checkIndices = append(checkIndices, indicesForHouse(houseForIndex(ind))...)

	return checkIndices
}

/// Debug Helpers ///

func (p *Puzzle) reset() {
	for i, _ := range p.Board {
		p.Board[i] = 0
	}
}

func (p *Puzzle) showIndices() {
	for i, _ := range p.Board {
		p.Board[i] = i
	}
}
