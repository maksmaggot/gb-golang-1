package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Println("Введите число: ")
	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Println("Ошибка ввода числа!")
	}

	fmt.Println("Простые числа: ")
	for i := 2; i <= number; i++ {
		if isSimple(i) {
			fmt.Printf("%d \n", i)
		}
	}
}

func isSimple(number int) bool {
	for j := 2; j < number; j++ {
		if number%j == 0 {
			return false
		}
	}

	return true
}
