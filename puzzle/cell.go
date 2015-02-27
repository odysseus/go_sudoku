package puzzle

import (
	"fmt"
)

/// Cell Struct ///
type Cell struct {
	value, candcount int
	candidates       []bool
	filled           bool
}

func NewCell() Cell {
	c := Cell{
		value:      0,
		candcount:  9,
		candidates: make([]bool, 10),
		filled:     false,
	}
	for i, _ := range c.candidates {
		c.candidates[i] = true
	}
	return c
}

func (c *Cell) String() string {
	if c.value == 0 {
		return "."
	} else {
		return fmt.Sprintf("%v", c.value)
	}
}

func (c *Cell) Value() int {
	return c.value
}

func (c *Cell) CandidateCount() int {
	if c.filled {
		return 0
	} else {
		return c.candcount
	}
}

func (c *Cell) Filled() bool {
	return c.filled
}

func (c *Cell) Candidates() []bool {
	// Returns a copy of the current candidates array
	r := make([]bool, 9)
	_ = copy(r, c.candidates)
	return r
}

func (c *Cell) Set(n int) {
	// "Fills in" the cell
	c.value = n
	c.filled = true
}

func (c *Cell) RemoveCandidate(n int) {
	// Removes a candidate from the array and decrements the count
	// of valid candidates
	if !c.filled {
		c.candidates[n] = false
		c.candcount--
	}
}

func (c *Cell) TestRemoval(n int) int {
	// Returns the number of candidates remaining after n is removed
	if c.candidates[n] {
		return c.candcount - 1
	} else {
		return c.candcount
	}
}

func (c *Cell) IntCandidates() []int {
	fin := make([]int, 0, 8)
	for i := 1; i < len(c.candidates); i++ {
		if c.candidates[i] {
			fin = append(fin, i)
		}
	}
	return fin
}
