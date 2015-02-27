package main

import (
	"fmt"
	"github.com/odysseus/stopwatch"
	"github.com/odysseus/sudoku/puzzle"
	"math/rand"
)

func main() {
	sw := stopwatch.New()

	p := Generate()
	fmt.Println(p)

	fmt.Println(sw)
}

func RandSudokuNum() int {
	return rand.Intn(9) + 1
}

func Generate() *puzzle.Puzzle {
	p := puzzle.NewPuzzle()
	// Place a value in every single index
	for i, v := range p.Board {
		placed := false
		// Get a slice of all the valid candidates
		cands := v.IntCandidates()
		ci := puzzle.CheckIndices(i)
		n := rand.Intn(len(cands))
		for !placed {
			// Loop until we successfully place a value
			// Grab a random element from the slice
			placed = true
			for i, _ := range ci {
				if p.Board[i].TestRemoval(cands[n]) == 0 {
					placed = false
					n = n + 1%len(cands)
					break
				}
			}
			if placed {
				v.Set(cands[n])
				for i, _ := range ci {
					p.Board[i].RemoveCandidate(cands[n])
				}
			}
		}
		fmt.Print(i)
	}

	return p
}
