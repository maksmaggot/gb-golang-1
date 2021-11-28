package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const addition = "+"
const subtraction = "-"
const multiplication = "*"
const division = "/"
const exponentiation = "^"
const calc = "calc"
const history = "history"

func main() {
	fmt.Println("Калькулятор")

	var a, b, result float64
	var operand string
	var historyRepository []float64

	for {
		fmt.Println("Введите первое число: ")
		_, error := fmt.Scanln(&a)
		if error != nil {
			fmt.Println("Не верный ввод числа")
			os.Exit(1)
		}

		fmt.Println("Введите оператор, (*, /, +, -, ^)")
		_, error = fmt.Scanln(&operand)
		if error != nil {
			fmt.Println("Не верный ввод оператора")
			os.Exit(1)
		}

		fmt.Println("Введите второе число: ")
		_, error = fmt.Scanln(&b)
		if error != nil {
			fmt.Println("Не верный ввод числа")
			os.Exit(1)
		}

		result, error = calculate(a, b, operand)
		if error != nil {
			fmt.Printf("Ошибка: %s\n", error)
			os.Exit(1)
		}

		historyRepository = append(historyRepository, result)
		fmt.Printf("Результат: %f\n", result)

		var operation string
		fmt.Println("Чтобы повторить ввод - введите calc, показать историю - history")
		_, error = fmt.Scan(&operation)
		if error != nil {
			fmt.Println("Ошибка ввода")
			os.Exit(1)
		}

		switch operation {
		case calc:
			continue
		case history:
			historyString := splitHistory(historyRepository)
			fmt.Printf("История: %s", historyString)
			os.Exit(1)
		default:
			fmt.Println("Операция введена не верно")
			os.Exit(1)
		}

	}
}

func splitHistory(historyRepository []float64) string {

	var stringBuilder strings.Builder
	for _, historyItem := range historyRepository {
		stringBuilder.WriteString(strconv.FormatFloat(historyItem, 'E', -1, 64) + ", ")
	}

	return stringBuilder.String()
}

func calculate(a, b float64, operand string) (float64, error) {
	switch operand {
	case addition:
		return a + b, nil
	case subtraction:
		return a - b, nil
	case multiplication:
		return a * b, nil
	case division:
		if b == 0 {
			return 0, errors.New("ошибка деления на 0")
		}
		return a / b, nil
	case exponentiation:
		return math.Pow(a, b), nil
	default:
		return 0, errors.New("указан не верный оператор")
	}

}
