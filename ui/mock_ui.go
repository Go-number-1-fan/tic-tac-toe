package ui

type MockUI struct {
	MockUserInput int
}

func (ui MockUI) DisplayWelcomeMessage() {}

func (ui MockUI) DisplayComputerThinkingMessage(computerName string) {}

func (ui MockUI) GetValidMove(openSpots []int, marker string) int {
	return ui.MockUserInput
}

func (ui MockUI) DisplayBoard(board []string, player1Marker string, player2Marker string) {}

func (ui MockUI) DisplayTieMessage() {}

func (ui MockUI) DisplayWinMessage(winner string) {}

func (ui MockUI) GetRefereeSelection() RefereeTypeSelection {
	return RefereeTypeSelection(ui.MockUserInput)
}

func (ui MockUI) GetPlayerTypeSelection(playerName string, playerMarker string) PlayerTypeSelection {
	return PlayerTypeSelection(ui.MockUserInput)
}

func (ui MockUI) GetPlayerNameSelection(playerNumber string) string {
	return "Player" + playerNumber
}

func (ui MockUI) GetPlayerMarkerSelection(playerName string, playerNumber string) string {
	return playerName[:1]
}

func (ui MockUI) GetPlayAgainSelection() bool {
	return false
}

func (ui MockUI) Clear() {
}
