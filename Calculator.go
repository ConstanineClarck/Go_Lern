package main

import "fmt"

func main() {
	// Выводим приветствие и принцип использования калькулятора
	fmt.Println("Добро пожаловать ")
  fmt.Println(" ")
	fmt.Println("Введите первое число, оператор (+, -, *, /), затем второе число:")

	// Получаем ввод пользователя
	var num1, num2 float64
	var operator string
	fmt.Print("Первое число: ")
	fmt.Scanln(&num1)
	fmt.Print("Оператор: ")
	fmt.Scanln(&operator)
	fmt.Print("Второе число: ")
	fmt.Scanln(&num2)

	// Вычисляем результат на основе оператора
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Ошибка: деление на ноль!")
			return
		}
	default:
		fmt.Println("Ошибка: введён не допустимый оператор!")
		return
	}

	// Выводим результат
	fmt.Printf("Результат: %.2f\n", result)
}
