package main

import (
	"math"

	"github.com/notnil/chess"
)

type SearchTree struct {
	position *chess.Position
	move     *chess.Move
	eval     int
	cutoff   bool
	children []*SearchTree
}

func (tree *SearchTree) AddChild(child *SearchTree) {
	tree.children = append(tree.children, child)
}

func (tree *SearchTree) Expand(depth int) {
	if depth == 0 {
		return
	}
	for _, move := range tree.position.ValidMoves() {
		child := &SearchTree{position: tree.position.Update(move), move: move}
		child.Expand(depth - 1)
		tree.AddChild(child)
	}
}

func (tree *SearchTree) AlphaBetaExpand(evalf func(*chess.Position) int, depth int) {
	alpha := math.MinInt + 2
	beta := math.MaxInt - 2
	if tree.position.Turn() == chess.Black {
		alpha, beta = -beta, -alpha
	}

	tree.AlphaBetaStep(evalf, alpha, beta, depth)
}

func (tree *SearchTree) AlphaBetaStep(evalf func(*chess.Position) int, alpha, beta, depth int) int {
	if depth == 0 {
		return evalf(tree.position)
	}

	tree.Expand(1)
	for _, child := range tree.children {
		child.eval = -child.AlphaBetaStep(evalf, -beta, -alpha, depth-1)

		if child.eval >= beta {
			child.cutoff = true
			return beta
		}
		if child.eval > alpha {
			alpha = child.eval
		}
	}
	return alpha
}

func (tree *SearchTree) BestChild() *SearchTree {
	best := tree.children[0]

	for _, child := range tree.children {
		if child.eval > best.eval {
			best = child
		}
	}

	return best
}
