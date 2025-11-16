package ui

import (
	"fmt"

	c "github.com/drenk83/NumberGuessingGame/config"
)

func WelcomeMessage(cfg c.GameConfig) {
	switch cfg.Lang() {
	case c.Eng:
		fmt.Println("Welcome to the \"Number Guessing Game\"!")
		fmt.Println("Game mode:", cfg.Mode())
		fmt.Println("Number of attempts:", cfg.MaxAttempt())
		fmt.Println("Secret number range from", cfg.Min(), "to", cfg.Max())
		fmt.Print("\n\n")
		engChoice()
	case c.Rus:
		fmt.Println("Добро пожаловать в \"Number Guessing Game\"!")
		fmt.Println("Режим игры:", cfg.Mode())
		fmt.Println("Количество попыток:", cfg.MaxAttempt())
		fmt.Println("Загадываемое число от", cfg.Min(), "до", cfg.Max())
		fmt.Print("\n\n")
		rusChoice()
	}
}

func rusChoice() {
	fmt.Println("Выберите:")
	fmt.Println("1 - Начать игру")
	fmt.Println("2 - Сменить режим")
	fmt.Println("3 - Сменить язык")
	fmt.Println("4 - Сменить количество попыток")
	fmt.Println("5 - Сменить диапазон")
	fmt.Println("6 - Выход")
}

func engChoice() {
	fmt.Println("Choose:")
	fmt.Println("1 - Start game")
	fmt.Println("2 - Change mode")
	fmt.Println("3 - Change language")
	fmt.Println("4 - Change number of attempts")
	fmt.Println("5 - Change range")
	fmt.Println("6 - Exit")
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
