package game

import (
	"fmt"
	"math/rand"

	"github.com/drenk83/NumberGuessingGame/config"
)

type GameRule struct {
	randNumber int
	attempts   []int
	lenAtt     int
	maxAtt     int
}

func Game(cfg config.GameConfig) {
	var randNum int

	switch cfg.Mode() {
	case config.ModeDefault:
		for randNum = rand.Intn(cfg.Max()); randNum == 0; {
			randNum = rand.Intn(cfg.Max())
		}
	case config.ModeBot:
		randNum = UserInput()
	}

	g := MakeGame(cfg, randNum)

	fmt.Println(*g)
}

func MakeGame(c config.GameConfig, r int) *GameRule {
	return &GameRule{
		randNumber: r,
		attempts:   make([]int, 0, c.MaxAtt()),
		lenAtt:     0,
		maxAtt:     c.MaxAtt(),
	}
}

func UserInput() int {
	return 123
}
