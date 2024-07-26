package main

import (
	"fmt"
	"time"
)

func main() {
	//consume favorite fruits

	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	fruits := []string{"mango", "strawberry", "orange", "kiwi", "muskmelon"}
	count := 0
	for _, bestFruit := range fruits { // k ,v pattern to loop as key(k) represents index and v(value) represents the value at that index

		// consumeFruit(bestFruit, count) //without goRoutine
		go consumeFruit(bestFruit, count) //using goRoutine
		count += 1
	}

	time.Sleep(time.Second * 2)
}

func consumeFruit(targetFruit string, count int) {
	fmt.Println("Consuming Fruit", targetFruit)

	fmt.Println(count)
	//core logic
	time.Sleep(time.Second)
}
