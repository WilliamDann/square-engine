package main

import (
	"fmt"
	"time"

	"github.com/notnil/chess"
)

func main() {
	game := chess.NewGame()
	pause := true

	// generate moves until game is over
	for game.Outcome() == chess.NoOutcome {
		start := time.Now()
		info := iterSearch(game.Position(), 5)
		end := time.Now()

		game.Move(info.value.move)
		if pause {
			if game.Position().Turn() == chess.Black {
				info.value.eval = -info.value.eval // negamax
			}
			fmt.Println(game.Position().Board().Draw())
			fmt.Print(game.String())
			fmt.Printf("\neval: %d\n", info.value.eval)
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
		}
	}

	// print outcome and game PGN
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
	fmt.Println(game.String())
}
