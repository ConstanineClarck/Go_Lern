package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Добро пожаловать в программу Calculator! Если хотите выйти - ведите 'exit' для выхода.")

	for {
		num1 := getUserInput("Первое число:")
		operator := getUserInput("Оператор (+, -, *, /):")
		num2 := getUserInput("Второе число:")

		if operator == "exit" {
			fmt.Println("До свидания!")
			os.Exit(0)
		}

		result, err := calculateResult(num1, operator, num2)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func getUserInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt + " ")
	fmt.Scanln(&input)
	return input
}

func calculateResult(num1 float64, operator string, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 != 0 {
			return num1 / num2, nil
		} else {
			return 0, fmt.Errorf("Ошибка: деление на ноль!")
		}
	default:
		return 0, fmt.Errorf("Ошибка: неверный оператор!")
	}
}
