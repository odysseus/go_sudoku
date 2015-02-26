package main

import (
	"fmt"
	"github.com/odysseus/stopwatch"
	"github.com/odysseus/sudoku/puzzle"
)

func main() {
	sw := stopwatch.New()

	b := puzzle.NewPuzzle()
	b.Place(5, 5)
	b.Place(7, 4)
	b.Place(3, 4)

	fmt.Println(b)

	fmt.Println(sw)
}
