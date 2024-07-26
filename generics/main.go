package main

import "fmt"

type Number interface {
	int64 | float64
}

func addIntValues(mpp map[string]int64) int64 { //this contains a mpp named map that has {string(key) , int64(value)}
	var s int64

	for _, value := range mpp {
		s += value
	}

	return s
}

func addFloatValue(mpp map[string]float64) float64 {
	var s float64

	for _, value := range mpp {
		s += value
	}

	return s
}

func sumIntorFloats[Key comparable, Value int64 | float64](mpp map[Key]Value) Value { //return type of Value to be decided in compile time
	var s Value // as 's' is to be returned it will Infer the return type as that of Value that is decided on fn call
	for _, value := range mpp {
		s += value
	}
	return s
}

func main() {

	// Initialise map of int64 as value

	ints := map[string]int64{
		"first":  50,
		"second": 100,
	}

	// Initialise map of float64 as value

	floating := map[string]float64{
		"first":  50.10,
		"second": 100.20,
	}

	// Initialise map according to the generic function

	fmt.Printf("Generic sum %v and %v :\n", sumIntorFloats[string, int64](ints), sumIntorFloats[string, float64](floating))

	fmt.Printf("Non-Generic Sums %v and %v \n", addIntValues(ints), addFloatValue(floating))
}
