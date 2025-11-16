package main

import (
	"fmt"
	"os"

	"github.com/drenk83/NumberGuessingGame/config"
	"github.com/drenk83/NumberGuessingGame/game"
	"github.com/drenk83/NumberGuessingGame/ui"
)

func main() {
	cfg := config.MakeCfg()
	args := os.Args

	game.Game()
	ui.FirstMessage(cfg.Mode(), cfg.MaxAttempt())

	fmt.Println("input args: ", args)
}
