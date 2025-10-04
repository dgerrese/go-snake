package gamestate

type GameState byte

const (
	Starting GameState = iota
	Running
	Paused
	Ended
	Menu
)

var gameStateNames = map[GameState]string{
	Starting: "Starting",
	Running:  "Running",
	Paused:   "Paused",
	Ended:    "Ended",
	Menu:     "Menu",
}

func (gs GameState) String() string {
	return gameStateNames[gs]
}
