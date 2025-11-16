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
	min         int
	max         int
}

func MakeCfg(m *string, a *int, l *string) GameConfig {
	return GameConfig{
		gameMode:    *m,
		maxAttempts: *a,
		language:    *l,
		min:         1,
		max:         100,
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

func (g *GameConfig) Min() int {
	return g.min
}

func (g *GameConfig) Max() int {
	return g.max
}

func (g *GameConfig) ModeDefault() {
	g.gameMode = ModeDefault
}

func (g *GameConfig) ModeBot() {
	g.gameMode = ModeBot
}
