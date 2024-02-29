package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays in GOLANG")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "peach"

	fmt.Println("Fruit List is ,", fruitList)
	fmt.Println("length of the list is,", len(fruitList))

	var vegList = [3]string{"potato", "tomato", "mushroom"}
	fmt.Println("veggie list is ,", vegList)
	fmt.Println("length of the list is,", len(vegList))

}
