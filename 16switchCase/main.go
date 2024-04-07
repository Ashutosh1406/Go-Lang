package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("SWITCH CASE In Golang")
	//count := 20
	//for i := 0; i < count; i++ {
	var number int = rand.Intn(6) + 1
	fmt.Println("Value of number is", number)
	//}
	switch number {
	case 1:
		fmt.Println("Dice Value is 1 and you can START")
	case 2:
		fmt.Println("Move by 2 Places")
	case 3:
		fmt.Println("Move by 3 Places")
		fallthrough
	case 4:
		fmt.Println("Move by 4 Places")
		fallthrough
	case 5:
		fmt.Println("Move by 5 Places")
		fallthrough
	case 6:
		fmt.Println("Move by 6 places and you got another chance Rollagain()")
	default:
		fmt.Println("Ouch! That wasn't expected")

	}

}
