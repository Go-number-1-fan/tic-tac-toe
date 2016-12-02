package ui

const (
	DefaultPlayerName    string = "Player"
	DefaultMaxNameLength int    = 15
	DefaultPlayer1Marker string = "X"
	DefaultPlayer2Marker string = "O"
)
const (
	WelcomeMessage          string = "\n\n\nWelcome To Tic Tac Toe!\nFollow the instructions and enter requested information when prompted!\n\nMost Importantly... Have fun!"
	ComputerThinkingMessage string = " is Thinking...\n"
	TieMessage              string = "Game Over!\nIt's a tie!!\n"
	WinMessage              string = " Wins!!\n"
)

type PlayerTypeSelection int

const (
	HumanSelected        PlayerTypeSelection = 1
	EasyComputerSelected PlayerTypeSelection = 2
	HardComputerSelected PlayerTypeSelection = 3
)

const (
	SelectASpotMessage string = ", Please select an open spot of the board by the index:\n"
	NotValidMessage    string = "Your selection is Invalid!!\n"
	NotANumberMessage  string = "You must enter a NUMBER: "

	NameSelectMessage       string = "Please Select a Name for Player"
	MarkerSelectMessage     string = "Please Select a Marker for "
	PlayAgainSelectMessage  string = "Would you like to play again? (y | n):"
	PlayerTypeSelectMessage string = "Please Select a Player Type:\n" +
		"\t1. Human Player\n" +
		"\t2. Easy Computer Player\n" +
		"\t3. Hard Computer Player\n"
)
const (
	NewLineString           string = "\n"
	HorizontalDividerString string = "-"
	VerticalDividerString   string = " | "
)
