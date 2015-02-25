package main

import (
	"fmt"
)

func main() {
	b := NewPuzzle()
	rowcol := 0

	for i := 0; i < 9; i++ {
		rowcol = i
		row := indicesForRow(rowcol)
		for _, v := range row {
			b.board[v] = 1
		}

		column := indicesForColumn(rowcol)
		for _, v := range column {
			b.board[v] = 1
		}

		fmt.Println(b)
		b.reset()
	}
}

type Puzzle struct {
	board []int
}

func NewPuzzle() *Puzzle {
	return &Puzzle{board: make([]int, 81)}
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
		s += fmt.Sprintf("%v ", v)
	}
	s += "\n"
	return s
}

func (p *Puzzle) reset() {
	for i, _ := range p.board {
		p.board[i] = 0
	}
}

func rowForIndex(ind int) int {
	return ind / 9
}

func columnForIndex(ind int) int {
	return ind % 9
}

func houseForIndex(ind int) int {
	return ((rowForIndex(ind) / 3) * 3) + (columnForIndex(ind) / 3)
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
