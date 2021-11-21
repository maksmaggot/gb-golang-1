package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите последовательность чисел для сортировки, через запятую")

	var inputString string
	fmt.Scan(&inputString)

	numbers := parseToSlice(inputString)
	sorted := insertionSort(numbers)

	fmt.Printf("%d\n", sorted)
}

func parseToSlice(input string) []int {
	parsedString := strings.Split(input, ",")
	converted := make([]int, 0, len(parsedString))
	for _, parsed := range parsedString {
		convertedItem, err := strconv.Atoi(parsed)
		if err != nil {
			fmt.Println(err)
		}
		converted = append(converted, convertedItem)
	}

	return converted
}

func insertionSort(sortable []int) []int {
	for idx, value := range sortable {
		position := idx
		for ; position >= 1 && sortable[position-1] > value; position-- {
			sortable[position] = sortable[position-1]
		}
		sortable[position] = value
	}
	return sortable
}
