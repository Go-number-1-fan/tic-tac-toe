package game

import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import . "github.com/go-number-1-fan/tic-tac-toe/player"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"

type GameBuilder struct {
	UI UI
}

func (builder GameBuilder) BuildGame() Game {
	ref := builder.getReferee()
	player1 := builder.getPlayer("1", ref)
	player2 := builder.getPlayer("2", ref)
	return CreateGame(builder.UI, EmptyBoard(), player1, player2, ref)
}
func (builder GameBuilder) getReferee() Referee {
	refSelection := builder.UI.GetRefereeSelection()
	switch refSelection {
	case StandardRefSelected:
		return StandardReferee{}
	case CornerRefSelected:
		return CornerReferee{}
	}
	return StandardReferee{}
}

func (builder GameBuilder) getPlayer(playerNumber string, ref Referee) Player {
	playerNameSelection := builder.UI.GetPlayerNameSelection(playerNumber)
	playerMarkerSelection := builder.UI.GetPlayerMarkerSelection(playerNameSelection, playerNumber)
	playerTypeSelection := builder.UI.GetPlayerTypeSelection(playerNameSelection, playerMarkerSelection)
	switch playerTypeSelection {
	case HumanSelected:
		return HumanPlayer{playerMarkerSelection, playerNameSelection}
	case EasyComputerSelected:
		return EasyComputerPlayer{playerMarkerSelection, playerNameSelection}
	default:
		return HardComputerPlayer{playerMarkerSelection, playerNameSelection, ref}
	}
}
