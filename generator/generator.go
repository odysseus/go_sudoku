package main

import (
	"fmt"
	"github.com/odysseus/stopwatch"
	"github.com/odysseus/sudoku/puzzle"
	"log"
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
		cands := v.Candidates()
		l := len(cands)
		if l == 0 {
			log.Fatal(fmt.Sprintf("%v", p))
		}
		r := rand.Intn(l)
		n := cands[r]
		p.Place(i, n)
	}
	return p
}
