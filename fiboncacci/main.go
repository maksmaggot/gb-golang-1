package main

func main() {
	fibMap := make(map[int]int)
	fib(100, fibMap)
}

func fib(number int, fibMap map[int]int) int {

	if value, exists := fibMap[number]; exists {
		return value
	}

	if number <= 1 {
		return number
	}

	result := fib(number-1, fibMap) + fib(number-2, fibMap)
	fibMap[number] = result
	return result
}
