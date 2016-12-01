package referee

type GameStatus string

const (
	WinP1    GameStatus = "winp1"
	WinP2    GameStatus = "winp2"
	Tie      GameStatus = "tie"
	Continue GameStatus = "continue"
)
