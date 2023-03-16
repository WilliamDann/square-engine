package main

import "github.com/notnil/chess"

func material(position *chess.Position) int {
	score := 0
	for _, p := range position.Board().SquareMap() {
		value := 0
		switch p.Type() {
		case chess.Pawn:
			value += 100
		case chess.Bishop:
			value += 300
		case chess.Knight:
			value += 300
		case chess.Rook:
			value += 500
		case chess.Queen:
			value += 900
		case chess.King:
			value += 200
		}
		if p.Color() == chess.Black {
			score -= value
			continue
		}
		score += value
	}
	return score
}

// TODO per piece
func mobility(position *chess.Position) int {
	score := 0
	mod := 1

	if position.Turn() == chess.Black {
		mod = -mod
	}

	score += len(position.ValidMoves()) * mod
	flip := FlipSide(position)
	mod = -mod
	score += len(flip.ValidMoves()) * mod

	return score
}

func evaluate(position *chess.Position) int {
	score := material(position) + mobility(position)

	if position.Turn() == chess.Black {
		score = -score
	}
	return score
}
