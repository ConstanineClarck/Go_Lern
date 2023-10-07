package main // Делаем файл исполняемым

import (
  "fmt" // Для работы с вводомвыводом
  "os"  // Для работы с ОС

  "github.com/eiannone/keyboard"  // Для работы с клавиатурой
  "github.com/fatih/color"        // Для работы с разнацветными символами
)

func main() {

  // Объявляем переменную для ввода выражения
  var expr string

  // Выводим приглашение для ввода
  fmt.Print("Введите выражение: ")

  // Считываем введенную строку  
  fmt.Scanln(&expr)

  // Цикл обработки выражений
  for {
    
    // Вызываем функцию расчета
    result, err := calc(expr)

    // Проверяем наличие ошибки
    if err != nil {
    
      // Выводим сообщение об ошибке  
      fmt.Println(err)

      // Просим повторить ввод
      fmt.Print("Повторите ввод: ")
      fmt.Scanln(&expr)

      // Переход к следующей итерации
      continue
      
    }
    
    // Выводим результат
    color.New(color.FgMagenta).Println(result)

    // Выходим из цикла
    break
    
  }

  // Приглашение для выхода
  fmt.Println("Нажмите Esc для выхода")

  // Инициализируем работу с клавиатурой
  err := keyboard.Open()
  if err != nil {
    panic(err)
  }

  // Закроем доступ по завершении
  defer keyboard.Close() 
  
  // Цикл ожидания нажатия Esc
  for {
  
    // Ждем нажатия клавиши
    char, key, err := keyboard.GetKey()

    // Обработка возможной ошибки
    if err != nil {
      panic(err)
    }

    // Проверяем нажатие Esc
    if key == keyboard.KeyEsc {
    
      // Выходим из цикла
      break
      
    }
  }

}

// Функция вычисления выражения  
func calc(input string) (int, error) {

  // Объявляем стек
  var stack []int

  // Переменные для операндов и оператора
  var operand1, operand2 int
  var operator rune

  // Разбор входной строки 
  for _, char := range input {

    // Если символ - цифра
    if char >= '0' && char <= '9' {

      // Преобразуем в число
      operand1 = operand1*10 + int(char-'0')

    // Если символ - оператор
    } else {
    
      switch char {
      
        // Обработка каждого оператора
        case '+':
          operator = char
          
          stack = append(stack, operand1)
          operand1 = 0
        
        case '-':
          operator = char
        
          stack = append(stack, operand1)
          operand1 = 0
        
        case '*':
          operator = char

          stack = append(stack, operand1)
          operand1 = 0
        
        case '/':
          operator = char
        
          stack = append(stack, operand1)
          operand1 = 0
          
        default:
          // Обработка ошибки  
          return 0, fmt.Errorf("Неизвестный оператор %c", char)
        
      }

      // Вывод оператора  
      color.New(color.FgGreen).Printf("%c ", char)

    }

  }
  
  // Добавляем последний операнд в стек
  stack = append(stack, operand1)

  // Вычисляем выражение
  for i := 0; i < len(stack); i++ {

    operand2 = stack[i+1]

    switch operator {
    
      case '+':
        stack[i] += operand2
      
      case '-':  
        stack[i] -= operand2
      
      case '*':
        stack[i] *= operand2
	  
      case '/':
        stack[i] /= operand2
      
    }

  }

  // Возвращаем результат
  return stack[0], nil

}
