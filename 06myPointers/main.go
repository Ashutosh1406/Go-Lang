package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	// var ptr *int
	// fmt.Println("value of pointer is ,", ptr) //no value assigns nil to it

	myNumber := 23

	var pointer2 = &myNumber                               //refrencing of myNumber
	fmt.Println("value of actual pointer is ,", pointer2)  //prints "address" ` before ` dereferencing
	fmt.Println("value of actual pointer is ,", *pointer2) //prints "value"  ` after `  dereferencing

	*pointer2 = *pointer2 * 2
	fmt.Println("my number", myNumber)

	myString := "ashutosh"

	var pointer3 = &myString
	fmt.Println("value of pointer is", *pointer3)
	fmt.Println("address of pointer is", pointer3)
	// ans := *pointer3 for copying before append / + operation
	*pointer3 = *pointer3 + "ashutosh"
	ans := *pointer3 //after appending and copying inside a new string type ans variable
	fmt.Println("new string is", myString)
	fmt.Println("length of new String is ,", len(myString))
	fmt.Println(ans)

}
