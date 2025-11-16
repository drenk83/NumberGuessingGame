package config

const (
	defaultMode = "default"
	botMode     = "bot mode"

	eng = "eng"
)

type GameConfig struct {
	gameMode    string
	maxAttempts int
	language    string
}

func MakeCfg() *GameConfig {
	return &GameConfig{
		gameMode:    defaultMode,
		maxAttempts: 10,
		language:    eng,
	}
}

func (g *GameConfig) Mode() string {
	return g.gameMode
}

func (g *GameConfig) MaxAttempt() int {
	return g.maxAttempts
}

func (g *GameConfig) Lang() string {
	return g.language
}

func (g *GameConfig) DefaultMode() {
	g.gameMode = defaultMode
}

func (g *GameConfig) BotMode() {
	g.gameMode = botMode
}
