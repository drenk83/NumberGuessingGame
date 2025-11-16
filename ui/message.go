package ui

import (
	"fmt"

	"github.com/drenk83/NumberGuessingGame/config"
)

func FirstMessage(cfg config.GameConfig) {
	fmt.Println("Hello from UI")
	fmt.Println("Gamemode : ", cfg.Mode())
	fmt.Println("Attempt numbers :", cfg.MaxAttempt())
}

func ToManyArgs() {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("[WRN] Too many arguments. Invalid arguments will be ignored.")
	fmt.Println("------------------------------------------------------------")
}

func InvalidArg(arg, def string) {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("[ERR] Unknow arrgument:", arg, "— using default value:", def)
	fmt.Println("------------------------------------------------------------")
}
