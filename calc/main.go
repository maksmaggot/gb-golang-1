package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Калькулятор")

	var a, b, result float64
	var operator string

	fmt.Println("Введите первое число: ")
	_, err := fmt.Scanln(&a)

	if err != nil {
		fmt.Println("Ввод не корректен!")
		os.Exit(1)
	}

	fmt.Println("Введите второе число: ")
	_, err = fmt.Scan(&b)

	if err != nil {
		fmt.Println("Ввод не корректен!")
		os.Exit(1)
	}

	fmt.Println("Введите арифметический оператор (+, -, *, /, ^, V): ")
	_, err = fmt.Scan(&operator)

	if err != nil {
		fmt.Println("Некорректно введён оператор")
		os.Exit(1)
	}

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка деления на 0")
			os.Exit(1)
		}
		result = a / b
	case "^":
		result = math.Pow(a, b)
	default:
		fmt.Println("Не верно выбран оператор!")
		os.Exit(1)
	}

	fmt.Printf("Результат %.2f \n", result)

}
