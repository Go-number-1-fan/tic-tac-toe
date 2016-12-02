package ui

func ExampleWelcomeMessage() {
	ui := ConsoleUI{MockInput{0, ""}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.DisplayWelcomeMessage()
	// Output:
	//Welcome To Tic Tac Toe!
	//Follow the instructions and enter requested information when prompted!
	//
	//Most Importantly... Have fun!
}

func ExampleGetRefereeSelection() {
	ui := ConsoleUI{MockInput{2, ""}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetRefereeSelection()
	// Output:
	//Please Select a Ruleset:
	//	1. Standard Ruleset - 3 in a row, Horizontal/Vertical/Diagonal
	//	2. Corner Ruleset - Place a Marker in 3 of the 4 corners

}

func ExampleGetPlayerTypeSelection() {
	ui := ConsoleUI{MockInput{2, ""}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetPlayerTypeSelection("Player1", "X")
	// Output:
	//	Player1 | X
	//Please Select a Player Type:
	//	1. Human Player
	//	2. Easy Computer Player
	//	3. Hard Computer Player
}

func ExampleGetPlayerNameSelection() {
	ui := ConsoleUI{MockInput{2, "Player1"}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetPlayerNameSelection("1")
	// Output:
	//Please Select a Name for Player1:
}

func ExampleGetPlayerMarkerSelection() {
	ui := ConsoleUI{MockInput{2, "O"}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetPlayerMarkerSelection("Tom", "2")
	// Output:
	//Please Select a Marker for Tom:
}

func ExampleGetPlayAgainSelection() {
	ui := ConsoleUI{MockInput{2, "O"}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetPlayAgainSelection()
	// Output:
	//Would you like to play again? (y | n):
}

func ExampleCanClearTheBoard() {
	ui := ConsoleUI{MockInput{2, "O"}, MockOutput{}, ConsoleInputValidator{}}
	ui.Clear()
	// Output:
	//CLEARED
}
