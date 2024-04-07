package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops Learning in goLang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {      "Commonly Used"
	// 	if days[i] == "Thursday" {
	// 		fmt.Println("You got a Lucky Day")
	// 	}
	// 	fmt.Println(days[i], i)
	// }

	// for it := range days {                "Range Based as i mostly use in Cpp"
	// 	fmt.Println(days[it], &it) //Interesting
	// }

	//							         "Range Based with index and Pointing to the content"

	// for index, day := range days {  "{ day,index}=>{index,anything inside the slice(content)}"
	// 	fmt.Printf("The day is %v and the number is %v\n", day, index)
	// }

	arbitraryValue := 1

	for arbitraryValue < 10 {

		if arbitraryValue == 7 {
			goto ashutosh
		}
		if arbitraryValue == 6 {
			arbitraryValue++
			continue
		}
		fmt.Println("Value is ", arbitraryValue)
		arbitraryValue++
	}

	/* go label */
ashutosh:
	fmt.Println("THALA !!!!!")
}
