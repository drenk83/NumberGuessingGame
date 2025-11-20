package main

import (
	e "errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/drenk83/NumberGuessingGame/config"
	"github.com/drenk83/NumberGuessingGame/game"
	"github.com/drenk83/NumberGuessingGame/ui"
)

func main() {
	modeFlag := flag.String("mode", config.ModeDefault, "game mode")
	attemptsFlag := flag.Int("att", config.DefAttempt, "attempts")
	langFlag := flag.String("lang", config.Eng, "language eng")

	flag.Parse()

	if err := validFlags(modeFlag, attemptsFlag, langFlag); err != nil {
		fmt.Println(err)
		return
	}
	cfg := config.MakeCfg(modeFlag, attemptsFlag, langFlag)

	ui.WelcomeMessage(cfg)
	ui.Menu(cfg.Lang())
	validInput()
	game.Game(cfg)
}

func validFlags(mode *string, att *int, lang *string) error {

	if mode == nil || att == nil || lang == nil {
		return e.New("nil pointer")
	}

	switch *mode {
	case "default", "def", "d":
		*mode = config.ModeDefault
	case "botmode", "bot", "b":
		*mode = config.ModeBot
	default:
		ui.InvalidArg(*mode, config.ModeDefault)
		*mode = config.ModeDefault
	}

	switch {
	case *att > 1:
	default:
		ui.InvalidArg(strconv.Itoa(*att), strconv.Itoa(config.DefAttempt))
		*att = config.DefAttempt
	}

	switch *lang {
	case "english", "eng", "en", "e":
		*lang = config.Eng
	case "russian", "rus", "ru", "r":
		*lang = config.Rus
	default:
		ui.InvalidArg(*lang, config.Eng)
		*lang = config.Eng
	}
	return nil
}

// нужно будет подумать над указателями чтоб структуру передать
func validInput() {
	whileInput := true
	for whileInput {
		switch ui.MenuInput() {
		case 1:
			whileInput = false
		case 2:

		case 3:
		case 4:
		case 5:
		case 6:
			ExitGame()
		default:

		}
	}
}

func ExitGame() {
	os.Exit(1)
}
