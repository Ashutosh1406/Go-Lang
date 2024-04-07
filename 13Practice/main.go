package main

import (
	"fmt"
	"os"
)

func main() {

	var sum float64
	var numbers int

	for {
		var value float64
		_, err := fmt.Fscanln(os.Stdin, &value)

		if err != nil {
			break //if error points to Nil that means no error and if it "doesn't" point towards nil that means there is an error
		}
		sum += value
		numbers++
	}

	if numbers == 0 {
		fmt.Fprintln(os.Stderr, "No values input")
		os.Exit(-1)
	}

	fmt.Println("The average is ", sum/float64(numbers))

}
