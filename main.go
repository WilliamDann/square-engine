package main

import (
	"fmt"
	"time"

	"github.com/notnil/chess"
)

func main() {
	game := chess.NewGame()
	pause := true
	depth := 4

	// generate moves until game is over
	for game.Outcome() == chess.NoOutcome {
		tree := &SearchTree{position: game.Position()}

		start := time.Now()
		tree.AlphaBetaExpand(evaluate, depth)
		best := tree.BestChild()
		end := time.Now()

		game.Move(best.move)
		if pause {
			eval := best.eval / 100
			if game.Position().Turn() == chess.Black {
				eval = -eval // negamax
			}

			fmt.Println(game.Position().Board().Draw())
			fmt.Print(game.String())
			fmt.Printf("\ndepth: %d", depth)
			fmt.Printf("\neval: %d\n", eval)
			fmt.Printf("time: %.3fs\n", end.Sub(start).Seconds())

			var val string
			fmt.Scanln(&val)

			if val == "exit" {
				return
			}
			if val == "continue" {
				pause = false
			}
		} else {
			fmt.Print(game.String())
			fmt.Printf("\ntime: %.3fs", end.Sub(start).Seconds())
		}
	}

	// print outcome and game PGN
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
	fmt.Println(game.String())
}
