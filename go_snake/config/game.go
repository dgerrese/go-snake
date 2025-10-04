package config

type gameConfig struct {
	WorldWidth  int
	WorldHeight int
}

var GameConfig = gameConfig{
	WorldWidth:  16,
	WorldHeight: 16,
}
