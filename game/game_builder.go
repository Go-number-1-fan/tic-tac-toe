package game

import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import . "github.com/go-number-1-fan/tic-tac-toe/player"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"

type GameBuilder struct {
	UI  UI
	Ref Referee
}

func (builder GameBuilder) BuildGame() Game {
	player1 := builder.getPlayer("X", "Player 1")
	player2 := builder.getPlayer("O", "Player 2")
	return CreateGame(builder.UI, EmptyBoard(), player1, player2, builder.Ref)
}

func (builder GameBuilder) getPlayer(stringMarker string, playerName string) Player {
	playerTypeSelection := builder.UI.GetPlayerTypeSelection(playerName)

	switch playerTypeSelection {
	case HumanSelected:
		return HumanPlayer{stringMarker}
	case EasyComputerSelected:
		return EasyComputerPlayer{stringMarker}
	default:
		return HardComputerPlayer{stringMarker, builder.Ref}
	}
}
