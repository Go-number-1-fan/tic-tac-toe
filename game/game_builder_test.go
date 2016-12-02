package game

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import . "github.com/go-number-1-fan/tic-tac-toe/player"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"
import "testing"

func TestGameBuilder_CanCreateAGameWithOneHumanPlayer(t *testing.T) {
	gameBuilder := GameBuilder{MockUI{1}}
	game := gameBuilder.BuildGame()
	_, ok := game.players[0].(HumanPlayer)
	assert.True(t, ok)
}

func TestGameBuilder_CanCreateAGameWithOneEasyComputerPlayer(t *testing.T) {

	gameBuilder := GameBuilder{MockUI{2}}
	game := gameBuilder.BuildGame()
	_, ok := game.players[0].(EasyComputerPlayer)
	assert.True(t, ok)
}

func TestGameBuilder_CanCreateAGameWithOneHardComputerPlayer(t *testing.T) {
	gameBuilder := GameBuilder{MockUI{3}}
	game := gameBuilder.BuildGame()
	_, ok := game.players[0].(HardComputerPlayer)
	assert.True(t, ok)
}

func TestGameBuilder_CanCreateAGameWithAHumanPlayerWithAGivenName(t *testing.T) {
	gameBuilder := GameBuilder{ConsoleUI{MockInput{1, "Tom"}, MockOutput{}, ConsoleInputValidator{}}}
	game := gameBuilder.BuildGame()
	_, ok := game.players[0].(HumanPlayer)
	assert.True(t, ok)
	assert.Equal(t, "Tom", game.players[0].GetName())
}

func TestGameBuilder_CanCreateAGameWithAComputerPlayerWithAGivenName(t *testing.T) {
	gameBuilder := GameBuilder{ConsoleUI{MockInput{2, "Eric"}, MockOutput{}, ConsoleInputValidator{}}}
	game := gameBuilder.BuildGame()
	_, ok := game.players[1].(EasyComputerPlayer)
	assert.True(t, ok)
	assert.Equal(t, "Eric", game.players[1].GetName())
}

func TestGameBuilder_CanCreateAGameWithAHumanPlayerWithAGivenMarker(t *testing.T) {
	gameBuilder := GameBuilder{ConsoleUI{MockInput{1, "T"}, MockOutput{}, ConsoleInputValidator{}}}
	game := gameBuilder.BuildGame()
	_, ok := game.players[0].(HumanPlayer)
	assert.True(t, ok)
	assert.Equal(t, "T", game.players[0].GetMarker())
}

func TestGameBuilder_CanCreateAGameWithAComputerPlayerWithAGivenMarker(t *testing.T) {
	gameBuilder := GameBuilder{ConsoleUI{MockInput{2, "E"}, MockOutput{}, ConsoleInputValidator{}}}
	game := gameBuilder.BuildGame()
	_, ok := game.players[1].(EasyComputerPlayer)
	assert.True(t, ok)
	assert.Equal(t, "E", game.players[1].GetMarker())
}

func TestGameBuilder_CanCreateAGameWithAStandardReferee(t *testing.T) {
	gameBuilder := GameBuilder{ConsoleUI{MockInput{1, ""}, MockOutput{}, ConsoleInputValidator{}}}
	game := gameBuilder.BuildGame()
	_, ok := game.referee.(StandardReferee)
	assert.True(t, ok)
}

func TestGameBuilder_CanCreateAGameWithACornerReferee(t *testing.T) {
	gameBuilder := GameBuilder{ConsoleUI{MockInput{2, ""}, MockOutput{}, ConsoleInputValidator{}}}
	game := gameBuilder.BuildGame()
	_, ok := game.referee.(CornerReferee)
	assert.True(t, ok)
}
