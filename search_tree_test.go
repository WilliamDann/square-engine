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

func TestInorderInOrder(t *testing.T) {
	tree := &SearchTree{}
	i := 0
	for _, move := range chess.StartingPosition().ValidMoves() {
		tree = append(tree, &MoveInfo{move, i})
		i++
	}

	last := 101
	n := 0
	tree.Inorder(func(mi *MoveInfo) {
		if mi.eval > last {
			t.Errorf("Wrong order: %d", mi.eval)
		}
		last = mi.eval
		n++
	})

	if n != i {
		t.Errorf("Added != found: %d != %d", i, n)
	}
}

func TestInOrderRandomOrder(t *testing.T) {
	tree := &SearchTree{}
	added := 0
	for _, move := range chess.StartingPosition().ValidMoves() {
		tree = append(tree, &MoveInfo{move, rand.Intn(200) - 100})
		added++
	}

	last := 101
	n := 0
	tree.Inorder(func(mi *MoveInfo) {
		if mi.eval > last {
			t.Errorf("Wrong order: %d", mi.eval)
		}
		last = mi.eval
		n++
	})

	if n != added {
		t.Errorf("Added != found: %d != %d", added, n)
	}
}

func TestInorderReverseOrder(t *testing.T) {
	tree := &SearchTree{}
	added := 0
	total := len(chess.StartingPosition().ValidMoves())
	for _, move := range chess.StartingPosition().ValidMoves() {
		tree = append(tree, &MoveInfo{move, total - added})
		added++
	}

	last := 101
	n := 0
	tree.Inorder(func(mi *MoveInfo) {
		if mi.eval > last {
			t.Errorf("Wrong order: %d", mi.eval)
		}
		last = mi.eval
		n++
	})

	if n != added {
		t.Errorf("Added != found: %d != %d", added, n)
	}
}

// TODO preorder test
// TODO postorder test
