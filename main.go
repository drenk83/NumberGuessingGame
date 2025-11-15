package main

import (
	"fmt"
	"os"

	"github.com/drenk83/NumberGuessingGame/game"
	"github.com/drenk83/NumberGuessingGame/ui"
)

const (
	defMod   = "defmod"
	vsBotMod = "botmod"
)

type GameConfig struct {
	GameMode    string
	MaxAttempts uint
}

func main() {
	cfg := makeConfig()
	args := os.Args

	if len(args) > 1 {
		switch args[1] {
		case defMod:
			cfg.GameMode = defMod
		case vsBotMod:
			cfg.GameMode = vsBotMod
		default:
			fmt.Println("Unown arguments")
		}
	}

	game.Game()
	ui.FirstMessage(cfg.GameMode, cfg.MaxAttempts)

	fmt.Println("input args: ", args)
}

func makeConfig() *GameConfig {
	return &GameConfig{
		GameMode:    defMod,
		MaxAttempts: 10,
	}
}
