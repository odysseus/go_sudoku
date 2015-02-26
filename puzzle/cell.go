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
	return c.candidates
}

func (c *Cell) Set(n int) {
	c.value = n
	c.filled = true
}

func (c *Cell) RemoveCandidate(n int) {
	if !c.filled {
		c.candidates[n] = false
		c.candcount--
	}
}
