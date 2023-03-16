package main

import (
	"testing"

	"github.com/notnil/chess"
)

func TestFlipSide(t *testing.T) {
	loaded, err := chess.FEN("r2q1rk1/pp1n1ppp/2pbpn2/3p3b/8/1P1PPNPP/PBPN1PB1/R2Q1RK1 b - - 0 10")
	if err != nil {
		t.Error(err)
	}
	pos := chess.NewGame(loaded).Position()
	pre := pos.Turn()
	post := FlipSide(pos).Turn()

	if pre == post {
		t.Errorf("Did not flip: %s == %s", pre.String(), post.String())
	}
}
