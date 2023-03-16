package main

import (
	"math"

	"github.com/notnil/chess"
)

func iterSearch(position *chess.Position, depth int) *SearchTree {
	tree := &SearchTree{}
	for _, move := range position.ValidMoves() {
		tree = append(tree, &MoveInfo{move, 0})
	}

	for i := 1; i <= depth; i++ {
		tree = alphaBeta(position, tree, i)
	}

	return tree
}

func alphaBeta(position *chess.Position, moves *SearchTree, depth int) *SearchTree {
	alpha := math.MinInt + 2
	beta := math.MaxInt - 2
	tree := &SearchTree{}

	if position.Turn() == chess.Black {
		alpha, beta = -beta, -alpha
	}

	moves.Inorder(func(move *MoveInfo) {
		score := -alphaBetaStep(position.Update(move.move), -beta, -alpha, depth-1)
		tree = append(tree, &MoveInfo{move.move, score})
		if score >= beta {
			return
		}
		if score > alpha {
			alpha = score
		}
	})

	return tree
}

func alphaBetaStep(position *chess.Position, alpha int, beta int, depthleft int) int {
	if depthleft == 0 {
		return evaluate(position)
	}
	for _, move := range position.ValidMoves() {
		score := -alphaBetaStep(position.Update(move), -beta, -alpha, depthleft-1)
		if score >= beta {
			return beta //  fail hard beta-cutoff
		}
		if score > alpha {
			alpha = score // alpha acts like max in MiniMax
		}
	}
	return alpha
}
