package main

import (
	"fmt"

	"github.com/notnil/chess"
)

func ShowBoard(game *chess.Game, eval int) {
	if game.Position().Turn() == chess.Black {
		eval = -eval // negamax
	}
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf(game.String())
	fmt.Printf("eval: %d\n", eval)
}

func main() {
	game := chess.NewGame()
	pause := true

	// generate moves until game is over
	for game.Outcome() == chess.NoOutcome {
		move, score := Search(game.Position(), 4)

		game.Move(move)
		if pause {
			ShowBoard(game, score)
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
