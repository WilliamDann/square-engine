package main

import (
	"testing"

	"github.com/notnil/chess"
)

func TestAddChildStartingMoves(t *testing.T) {
	tree := &SearchTree{position: chess.StartingPosition()}
	added := 0
	for _, move := range tree.position.ValidMoves() {
		tree.AddChild(&SearchTree{position: tree.position.Update(move), move: move})
		added++
	}
	n := len(tree.children)
	if n != added {
		t.Errorf("%d != %d : Not all children found", n, added)
	}
}
