package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"apple", "2", "tomato", "peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Default fruitList is ,", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("New fruitList is {_append} ,", fruitList) //append

	fruitList = append(fruitList[1:5]) //start from 1 and end on 3 range isnt't inclusive
	fmt.Println(fruitList)

	fruitList = append(fruitList[:5]) //start from 1 and end on 3 range isnt't inclusive
	fmt.Println(fruitList)

	highScore := make([]int, 4) //just like vectors with .append() functionality

	highScore[0] = 234
	highScore[1] = 293
	highScore[2] = 833
	highScore[3] = 111
	// highScore[4] = 444
	highScore = append(highScore, 555, 873, 1000)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove value from slice based on index
	var courses = []string{"java", "javascript", "go", "python"}

	fmt.Println(courses)

	//now removing the value by index
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //part by part
	fmt.Println(courses)

}
