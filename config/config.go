package config

import (
	"fmt"
)

type GameMode int
type Language int

// Consts
// ---------------------------------
const (
	ClassicMode GameMode = iota
	TimedMode
)

const (
	LangEng Language = iota
	LangRus
)

const (
	DefaultMode        GameMode = ClassicMode
	DefaultMaxAttempts int      = 7
	DefaultLanguage    Language = LangEng
	DefaultMaxNumber   int      = 100
)

// Struct
// ---------------------------------
type GameConfig struct {
	mode        GameMode
	maxAttempts int
	language    Language
	maxNumber   int
}

// Constructors
// ---------------------------------
func NewGameConfig() *GameConfig {
	return &GameConfig{
		mode:        DefaultMode,
		maxAttempts: DefaultMaxAttempts,
		language:    DefaultLanguage,
		maxNumber:   DefaultMaxNumber,
	}
}

// Getters
// ---------------------------------
func (g *GameConfig) Mode() GameMode {
	return g.mode
}

func (g *GameConfig) MaxAttempts() int {
	return g.maxAttempts
}

func (g *GameConfig) Language() Language {
	return g.language
}

func (g *GameConfig) MaxNumber() int {
	return g.maxNumber
}

// Setters
// --------------------------------
func (g *GameConfig) SetMode(m GameMode) error {
	switch m {
	case ClassicMode, TimedMode:
		g.mode = m
		return nil

	default:
		return fmt.Errorf("[ERROR] in changing the mode. Got: %v - %d", m, m)
	}
}

func (g *GameConfig) SetMaxAttempts(m int) error {
	if m <= 0 {
		return fmt.Errorf("[ERROR] Zero or negative number of attempts. Got: %d", m)
	}
	g.maxAttempts = m
	return nil
}

func (g *GameConfig) SetLanguage(l Language) error {
	switch l {
	case LangEng, LangRus:
		g.language = l
		return nil

	default:
		return fmt.Errorf("[ERROR] in changing the language. Got: %v - %d", l, l)
	}
}

func (g *GameConfig) SetMaxNumber(mx int) error {
	if mx <= 1 {
		return fmt.Errorf("[ERROR] in changing the max number. Got: %d", mx)
	}
	g.maxNumber = mx
	return nil
}

// Strings
// --------------------------------
func (g GameMode) String() string {
	switch g {
	case ClassicMode:
		return "Classic"
	case TimedMode:
		return "Timed"
	default:
		return "Unknown"
	}
}

func (l Language) String() string {
	switch l {
	case LangEng:
		return "English"
	case LangRus:
		return "Russian"
	default:
		return "Unknown"
	}
}
