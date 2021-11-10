package main

import (
	"fmt"
	"math"
	"os"
)

const pi = math.Pi

func main() {
	/*
		Напишите программу, вычисляющую диаметр
		и длину окружности по заданной площади круга.
		Площадь круга должна вводиться пользователем с клавиатуры.
	*/

	fmt.Println("Программа для вычисления диаметра и длины окружности")

	fmt.Println("Введите площадь круга:")
	var square float64
	_, err := fmt.Scan(&square)

	if err != nil {
		fmt.Println("Ошибка ввода площади")
		os.Exit(1)
	}

	raduis := math.Sqrt(square / pi)

	diameter := raduis * 2
	length := 2 * pi * raduis

	fmt.Printf("Диаметр окружности: %.2f Длина окружности: %.2f \n", diameter, length)
}
