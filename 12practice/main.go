package main

import (
	"fmt"
)

func main() {
	a := 21
	b := 20.999999999999999 //at around 15th concurrent place it round up

	fmt.Printf("A: %6T %[1]v\n", a)
	fmt.Printf("B: %6T %[1]v\n", b) //where %v is value and %T is type
	a = int(b)

	fmt.Printf("A: %6T %[1]v\n", a)

}
