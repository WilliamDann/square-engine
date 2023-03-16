package main

import (
	"strings"

	"github.com/notnil/chess"
)

func FlipSide(position *chess.Position) *chess.Position {
	a, b := "w", "b"
	if position.Turn() == chess.Black {
		b, a = a, b
	}
	loaded, _ := chess.FEN(strings.Replace(position.String(), a, b, -1))
	return chess.NewGame(loaded).Position()
}
