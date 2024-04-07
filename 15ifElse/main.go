package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to if-elseif-else Program")

	var score int = 30
	if score > 33 {
		fmt.Println("You are passed")
	} else if score == 33 {
		fmt.Println("you are barely passed")
	} else {
		fmt.Println("You are failed")
	}

	if num := 3; num < 10 { //intialisation + checking of condition in the same line pressing a "semicolon" in between
		fmt.Println("Number is lesser than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}

}
