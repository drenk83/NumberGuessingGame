package game

import (
	"math/rand"

	"github.com/drenk83/NumberGuessingGame/config"
)

type GameRule struct {
	randNumber      int
	attempts        []int
	counterAttempts int
}

func GameInit(cfg *config.GameConfig) {
	random := rand.Intn(cfg.MaxNumber())
	gr := NewGameRule(random, cfg.MaxAttempts())

	switch cfg.Mode() {
	case config.ClassicMode:
		ClassicGame(gr)
	case config.TimedMode:
		TimeGame(gr)
	}
}

func ClassicGame(g *GameRule) {

}

func TimeGame(g *GameRule) {

}

func NewGameRule(r, m int) *GameRule {
	return &GameRule{
		randNumber:      r,
		attempts:        make([]int, 0, m),
		counterAttempts: 0,
	}
}
