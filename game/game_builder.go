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
	player1 := builder.getPlayer("X", "1")
	player2 := builder.getPlayer("O", "2")
	return CreateGame(builder.UI, EmptyBoard(), player1, player2, builder.Ref)
}

func (builder GameBuilder) getPlayer(stringMarker string, playerNumber string) Player {
	playerNameSelection := builder.UI.GetPlayerNameSelection(playerNumber)
	playerTypeSelection := builder.UI.GetPlayerTypeSelection(playerNameSelection)
	switch playerTypeSelection {
	case HumanSelected:
		return HumanPlayer{stringMarker, playerNameSelection}
	case EasyComputerSelected:
		return EasyComputerPlayer{stringMarker, playerNameSelection}
	default:
		return HardComputerPlayer{stringMarker, playerNameSelection, builder.Ref}
	}
}
