package main

import "fmt"

const LoginToken string = "asnjdhahd" // "CAPITAL FIRST LETTER MEANS PUBLIC"  // public
func main() {
	var username string = "ashutosh"
	fmt.Println(username)
	fmt.Printf("variables is of type: %T \n", username)

	var isVerified bool = true
	fmt.Println(isVerified)
	fmt.Printf("variables is of type: %T \n", isVerified)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variables is of type: %T \n", smallVal)

	var smallFloat float32 = 255.323232328372
	fmt.Println(smallFloat)
	fmt.Printf("variables is of type: %T \n", smallFloat)

	//default and some aliases

	var anotherVariable string = "*"
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	//implicit

	var website = "https://secure.com"
	fmt.Println(website)

	//no var style

	numberofUser := 30000.0 //( =>   " := ") declaration inside main
	fmt.Println(numberofUser)
	fmt.Printf("type of %T \n", numberofUser) //prints float

	fmt.Println(LoginToken)
}
