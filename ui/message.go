package ui

import "fmt"

func FirstMessage(gamemode string, attemptNums int) {
	fmt.Println("Hello from UI")
	fmt.Println("Gamemode : ", gamemode)
	fmt.Println("Attempt numbers :", attemptNums)
}
