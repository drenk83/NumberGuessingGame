package config

const (
	defaultMode = "defmod"

	eng = "eng"
)

type GameConfig struct {
	GameMode    string
	MaxAttempts uint
	Language    string
}

func MakeCfg() *GameConfig {
	return &GameConfig{
		GameMode:    defaultMode,
		MaxAttempts: 10,
		Language:    eng,
	}
}
