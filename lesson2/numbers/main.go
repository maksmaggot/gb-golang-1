package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	   С клавиатуры вводится трехзначное число.
	   Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
	*/

	fmt.Println("Программа отображающая разряды числа")

	fmt.Println("Введите число")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil || number < 100 || number > 999 {
		fmt.Println("Ошибка ввода числа")
		os.Exit(1)
	}

	/* вариант без цикла */
	hundreds := number / 100
	decimals := (number - hundreds*100) / 10
	units := number - hundreds*100 - decimals*10

	fmt.Printf("Сотни: %d, Десятки: %d, Единицы: %d \n", hundreds, decimals, units)

	/* вариант в цикле */
	var digits [3]int

	divider := 100
	for i := 0; i < len(digits); i++ {

		digits[i] = number / divider
		number = number - (number/divider)*divider
		divider = divider / 10
	}

	fmt.Printf("Сотни: %d, Десятки: %d, Единицы: %d \n", digits[0], digits[1], digits[2])

}
