package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Defer Learning")
	fmt.Println("Its easy")
	fmt.Println("printing the counting(1-5) with defer")
	defer fmt.Println("5")
	defer fmt.Println("4")
	defer fmt.Println("3")
	defer fmt.Println("2")
	fmt.Println("1") //defer stack[Lifo]

	defer myDefer()

}

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
