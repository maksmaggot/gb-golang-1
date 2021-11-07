package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Программа для расчёта площади прямоугольника")

	fmt.Println("Введите высоту:")
	var height uint
	_, err := fmt.Scan(&height)
	if err != nil {
		fmt.Println("Ошибка ввода высоты: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Введите ширину:")
	var width uint
	_, err = fmt.Scan(&width)
	if err != nil {
		fmt.Println("Ошибка ввода ширины: ", err.Error())
		os.Exit(1)
	}

	result, err := calculateSquare(height, width)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Площадь:", result)

}

func calculateSquare(height, width uint) (uint, error) {
	var result uint

	if height == 0 || width == 0 {
		return result, errors.New("cтороны не могут иметь нулевую длину")
	}

	result = height * width
	return result, nil
}
