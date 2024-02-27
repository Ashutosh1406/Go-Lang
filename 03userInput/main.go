package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user Inpur"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the price:")

	// comma OK || err OK

	input, _ := reader.ReadString('\n') //if everything goes right "input" will contain it otherwise _ or error will be shown
	fmt.Println("Thanks for rating ,", input)
	fmt.Printf("type of rating is %T", input)

}
