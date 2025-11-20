package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MenuInput() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите число: ")
		inputStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("[ERR] Ошибка чтения: %v\n", err)
			continue
		}

		inputStr = strings.TrimSpace(inputStr)
		input, err := strconv.Atoi(inputStr)
		if err == nil {
			return input
		}

		fmt.Printf("[ERR] '%s' не является числом. Попробуйте снова.\n", inputStr)
	}
}
