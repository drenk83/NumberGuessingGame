package config

const (
	ModeDefault = "default"
	ModeBot     = "botmode"

	DefAttempt = 7

	Eng = "eng"
	Rus = "ru"
)

type GameConfig struct {
	gameMode    string
	maxAttempts int
	language    string
}

func MakeCfg() GameConfig {
	return GameConfig{
		gameMode:    ModeDefault,
		maxAttempts: DefAttempt,
		language:    Eng,
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

func (g *GameConfig) ModeDefault() {
	g.gameMode = ModeDefault
}

func (g *GameConfig) ModeBot() {
	g.gameMode = ModeBot
}
