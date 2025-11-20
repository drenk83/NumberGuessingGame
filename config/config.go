package config

import (
	"errors"
)

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

// Getters
// ---------------------------------
func (g *GameConfig) Mode() string {
	return g.gameMode
}

func (g *GameConfig) MaxAtt() int {
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

// Setters
// --------------------------------
func (g *GameConfig) SetMode(m string) error {
	switch m {
	case ModeDefault:
		g.gameMode = ModeDefault
		return nil

	case ModeBot:
		g.gameMode = ModeBot
		return nil

	default:
		return errors.New("[ERROR] in changing the mode")
	}
}

func (g *GameConfig) SwapMode() {
	if g.gameMode == ModeDefault {
		g.gameMode = ModeBot
	} else {
		g.gameMode = ModeDefault
	}
}

func (g *GameConfig) SetLang(l string) error {
	switch l {
	case Eng:
		g.language = Eng
		return nil

	case Rus:
		g.language = Rus
		return nil

	default:
		return errors.New("[ERROR] in changing the language")
	}
}

func (g *GameConfig) SetAttempts(a int) error {
	if a <= 0 {
		return errors.New("[ERROR] Zero or negative number of attempts")
	}
	g.maxAttempts = a
	return nil
}

func (g *GameConfig) SetMinMax(mn, mx int) error {
	if mn > mx {
		return errors.New("[ERROR] Min is more than max")
	}
	if mn == mx {
		return errors.New("[ERROR] Same values")
	}
	g.min = mn
	g.max = mx
	return nil
}
