package main

import (
	"math"

	"github.com/notnil/chess"
)

func Search(position *chess.Position, depth int) (*chess.Move, int) {
	alpha := math.MinInt + 2
	beta := math.MaxInt - 2

	if position.Turn() == chess.Black {
		alpha, beta = -beta, -alpha
	}

	var bestMove *chess.Move
	for _, move := range position.ValidMoves() {
		score := -alphaBeta(position.Update(move), -beta, -alpha, depth-1)
		if score >= beta {
			return move, beta //  fail hard beta-cutoff
		}
		if score > alpha {
			alpha = score // alpha acts like max in MiniMax
			bestMove = move
		}
	}

	return bestMove, alpha
}

func alphaBeta(position *chess.Position, alpha int, beta int, depthleft int) int {
	if depthleft == 0 {
		return evaluate(position)
	}
	for _, move := range position.ValidMoves() {
		score := -alphaBeta(position.Update(move), -beta, -alpha, depthleft-1)
		if score >= beta {
			return beta //  fail hard beta-cutoff
		}
		if score > alpha {
			alpha = score // alpha acts like max in MiniMax
		}
	}
	return alpha
}
