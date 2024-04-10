package main

import (
	"fmt"
)

func main() {

	fmt.Println("Functions in GoLang")

	greeter()

	result := add(3, 5)
	fmt.Println("Addition of numbers is", result)

	proresult := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Total sum is", proresult)
}
func greeter() {
	fmt.Println("Jai Jasjappan")
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func add(value1 int, value2 int) int {
	return value1 + value2
}
