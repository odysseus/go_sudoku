package main

import (
	"fmt"
	"github.com/odysseus/stopwatch"
	"github.com/odysseus/sudoku/puzzle"
	"log"
	"math/rand"
	"time"
)

func main() {
	sw := stopwatch.New()

	rand.Seed(int64(time.Now().Nanosecond()))
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
	i := 0
	for i < len(p.Board) {
		cands := p.CandidatesForIndex(i)
		l := len(cands)
		if l == 0 {
			log.Fatal(fmt.Sprintf("%v", p))
		}
		r := rand.Intn(l)
		n := cands[r]
		p.Board[i] = n
		i++
	}
	return p
}
