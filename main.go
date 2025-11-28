package main

import (
	"github.com/drenk83/NumberGuessingGame/config"
	"github.com/drenk83/NumberGuessingGame/game"
	"github.com/drenk83/NumberGuessingGame/ui"
)

func main() {
	cfg := config.NewGameConfig()

	ui.WelcomeMessage(cfg)
	//ui.Menu(cfg.Lang())

	for {
		switch ui.MenuInput() {
		case 1:
			game.GameInit(cfg)
		case 2:
		case 3:
		case 4:
		case 5:
		case 6:
		default:
		}
	}
}
