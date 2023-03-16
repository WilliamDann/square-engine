package main

import (
	"github.com/notnil/chess"
)

var PIECE_VALUE = map[chess.PieceType]int{
	chess.Pawn:   100,
	chess.Bishop: 300,
	chess.Knight: 300,
	chess.Rook:   500,
	chess.Queen:  900,
	chess.King:   1000,
}

var PIECE_MAX_MOVES = map[chess.PieceType]int{
	chess.Pawn:   2,
	chess.Bishop: 13,
	chess.Knight: 8,
	chess.Rook:   14,
	chess.Queen:  27,
	chess.King:   8,
}

func mobilityWeight(position *chess.Position, square *chess.Square, piece *chess.Piece) int {

}

func material(position *chess.Position) int {
	score := 0

	for square, piece := range position.Board().Rotate().SquareMap() {
		value := PIECE_VALUE[piece.Type()] * mobilityWeight(position, &square, &piece)
		if piece.Color() == chess.Black {
			value = -value
		}
		score += value
	}

	return score
}

func evaluate(position *chess.Position) int {
	score := material(position)

	if position.Turn() == chess.Black {
		score = -score
	}
	return score
}
