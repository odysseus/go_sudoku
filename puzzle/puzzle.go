package puzzle

import (
	"errors"
	"fmt"
)

/// Puzzle Struct ///
type Puzzle struct {
	board []Cell
}

func NewPuzzle() *Puzzle {
	p := Puzzle{board: make([]Cell, 81)}
	for i, _ := range p.board {
		p.board[i] = NewCell()
	}
	return &p
}

func (p *Puzzle) String() string {
	s := ""
	for i, v := range p.board {
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
		s += fmt.Sprintf("%v ", v.String())
	}
	s += "\n"
	return s
}

func (p *Puzzle) validPlacement(ind, val int) bool {
	// Ensures that the cell is not filled and tha val is a valid candidate
	cell := &p.board[ind]
	if cell.Filled() {
		return false
	} else {
		return cell.Candidates()[val]
	}
}

func (p *Puzzle) Place(ind, val int) error {
	// Places the
	if p.validPlacement(ind, val) {
		p.board[ind].Set(val)

		checkIndices := checkIndicesForIndex(ind)
		for i, _ := range checkIndices {
			p.board[i].RemoveCandidate(val)
		}
		return nil
	} else {
		return errors.New(
			fmt.Sprintf("%d is not a valid placement for %d", val, ind))
	}
}

/// Debug Helpers ///

func (p *Puzzle) reset() {
	for i, _ := range p.board {
		p.board[i].Set(0)
	}
}

func (p *Puzzle) showIndices() {
	for i, _ := range p.board {
		p.board[i].Set(i)
	}
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

func checkIndicesForIndex(ind int) []int {
	checkIndices := indicesForRow(rowForIndex(ind))
	checkIndices = append(checkIndices, indicesForColumn(columnForIndex(ind))...)
	checkIndices = append(checkIndices, indicesForHouse(houseForIndex(ind))...)

	return checkIndices
}
