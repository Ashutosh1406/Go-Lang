package main

import "fmt"

func main() {
	fmt.Println("structs in GO! Lang")

	//no inheritance in GoLang :, no "Super" or "Parent"

	structureParser := User{"Ashutosh", "ashutosh@gmail.com", false, 22}
	fmt.Println(structureParser)                                 //Println
	fmt.Printf("Ashtosh's details are : %+v\n", structureParser) // Printf with %+v full parsed
	fmt.Printf("Name is %v and email is %v.", structureParser.Name, structureParser.Email)
}

type User struct { //caps "U" in User
	Name   string
	Email  string
	Status bool
	Age    int
}
