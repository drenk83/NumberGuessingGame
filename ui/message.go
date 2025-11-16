package ui

import (
	"fmt"
)

const (
	eng = "eng"
	rus = "ru"
)

func WelcomeMessage(mode string, att int, lang string, min, max int) {
	switch lang {
	case eng:
		fmt.Println("Welcome to the \"Number Guessing Game\"!")
		fmt.Println("Game mode:", mode)
		fmt.Println("Number of attempts:", att)
		fmt.Println("Secret number range from", min, "to", max)
		fmt.Print("\n\n")
		engChoice()
	case rus:
		fmt.Println("Добро пожаловать в \"Number Guessing Game\"!")
		fmt.Println("Режим игры:", mode)
		fmt.Println("Количество попыток:", att)
		fmt.Println("Загадываемое число от", min, "до", max)
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

// For Setters
// ---------------------------
func SetLangMsg(l string) {
	if l == eng {
		fmt.Println("Language changed successfully!")
	} else {
		fmt.Println("Язык успешно изменён!")
	}
}

func SetModeMsg(l string) {
	if l == eng {
		fmt.Println("Mode changed successfully!")
	} else {
		fmt.Println("Режим успешно изменён!")
	}
}

func SetAttemptsMsg(l string, old int, new int) {
	if l == eng {
		fmt.Println("Number of attempts has been successfully changed:", old, "→", new)
	} else {
		fmt.Println("Количество попыток успешно изменено:", old, "→", new)
	}
}

func SetMinMaxMsg(l string, oldmin, newmin, oldmax, newmax int) {
	if l == eng {
		fmt.Println("The range of values has been successfully changed!")
		fmt.Println("Min:", oldmin, "→", newmin)
		fmt.Println("Max:", oldmax, "→", newmax)
	} else {
		fmt.Println("Диапазон значений успешно изменён!")
		fmt.Println("Мин:", oldmin, "→", newmin)
		fmt.Println("Макс:", oldmax, "→", newmax)
	}
}

// Errors
// ---------------------------
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
