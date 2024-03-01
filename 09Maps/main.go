package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps in GO!")
	languages := make(map[string]string)

	languages["Js"] = "javacript"
	languages["rb"] = "ruby"
	languages["cpp"] = "cplusplus"
	languages["py"] = "python"

	fmt.Println("Languages are - ", languages)
	fmt.Println("cpp stands for - ", languages["cpp"])

	//deletion
	delete(languages, "rb")
	fmt.Println("languages are", languages)

	//loops are interesting
	for _, value := range languages {
		fmt.Printf("for key v,value is %v\n", value)
	}
}
