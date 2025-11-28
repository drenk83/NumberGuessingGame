package ui

import (
	"fmt"

	"github.com/drenk83/NumberGuessingGame/config"
)

func WelcomeMessage(c *config.GameConfig) {
	switch c.Language() {
	case config.LangEng:
		fmt.Println("Welcome to the \"Number Guessing Game\"!")
		fmt.Println("Game mode:", c.Mode())
		fmt.Println("Number of attempts:", c.MaxAttempts())
		fmt.Println("Secret number range from", 1, "to", c.MaxNumber())
		fmt.Print("\n\n")
		engChoice()
	case config.LangRus:
		fmt.Println("Добро пожаловать в \"Number Guessing Game\"!")
		fmt.Println("Режим игры:", c.Mode())
		fmt.Println("Количество попыток:", c.MaxAttempts())
		fmt.Println("Загадываемое число от", 1, "до", c.MaxNumber())
		fmt.Print("\n\n")
		rusChoice()
	}
}

func Menu(l config.Language) {
	switch l {
	case config.LangEng:
		engChoice()
	case config.LangRus:
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

// For Setters
// ---------------------------
// func SetLangMsg(l string) {
// 	if l == config.Eng {
// 		fmt.Println("Language changed successfully!")
// 	} else {
// 		fmt.Println("Язык успешно изменён!")
// 	}
// }

// func SetModeMsg(l string) {
// 	if l == config.Eng {
// 		fmt.Println("Mode changed successfully!")
// 	} else {
// 		fmt.Println("Режим успешно изменён!")
// 	}
// }

// func SetAttemptsMsg(l string, old int, new int) {
// 	if l == config.Eng {
// 		fmt.Println("Number of attempts has been successfully changed:", old, "→", new)
// 	} else {
// 		fmt.Println("Количество попыток успешно изменено:", old, "→", new)
// 	}
// }

// func SetMinMaxMsg(l string, oldmin, newmin, oldmax, newmax int) {
// 	if l == config.Eng {
// 		fmt.Println("The range of values has been successfully changed!")
// 		fmt.Println("Min:", oldmin, "→", newmin)
// 		fmt.Println("Max:", oldmax, "→", newmax)
// 	} else {
// 		fmt.Println("Диапазон значений успешно изменён!")
// 		fmt.Println("Мин:", oldmin, "→", newmin)
// 		fmt.Println("Макс:", oldmax, "→", newmax)
// 	}
// }
