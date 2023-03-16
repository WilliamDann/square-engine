package main

import (
	"math/rand"
	"testing"

	"github.com/notnil/chess"
)

func TestAppendEmptyTree(t *testing.T) {
	move := chess.StartingPosition().ValidMoves()[0]
	tree := &SearchTree{}
	tree = append(tree, &MoveInfo{move, 050})

	if tree.value.move != move || tree.value.eval != 050 {
		t.Errorf(
			"Correct element not adde (%s, %d)",
			tree.value.move.String(),
			tree.value.eval,
		)
	}
}

func TestAppendStartingMoves(t *testing.T) {
	tree := &SearchTree{}
	added := 0
	for _, move := range chess.StartingPosition().ValidMoves() {
		tree = append(tree, &MoveInfo{move, rand.Intn(200) - 100})
		added++
	}

	found := 0
	tree.Inorder(func(mi *MoveInfo) { found++ })

	if added != found {
		t.Errorf("Added != found: %d != %d", added, found)
	}
}

// TODO preorder test
// TODO postorder test

func TestInorderTraverse(t *testing.T) {
	tree := &SearchTree{}
	i := 0
	for _, move := range chess.StartingPosition().ValidMoves() {
		tree = append(tree, &MoveInfo{move, i})
		i++
	}

	last := -1
	tree.Inorder(func(mi *MoveInfo) {
		if mi.eval <= last {
			t.Errorf("Wrong order: %d", mi.eval)
		}
		last = mi.eval
	})
}
