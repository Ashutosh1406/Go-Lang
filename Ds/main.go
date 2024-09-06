package main

import (
	"fmt"
)

type CustomSet struct {
	data map[int]bool
}

func NewCustomSet() *CustomSet {
	return &CustomSet{
		data: make(map[int]bool),
	}
}

func (s *CustomSet) Insert(x int) {
	s.data[x] = true
}

func (s *CustomSet) Find(x int) bool { //return type is bool as if found then return true; otherwise false
	_, doesExist := s.data[x]
	return doesExist
}

func (s *CustomSet) Remove(x int) {
	delete(s.data, x)
}

func (s *CustomSet) View() []int {
	elements := []int{} //slice of type int

	for key := range s.data {
		elements = append(elements, key)
	}
	return elements
}

func main() {
	mySet := NewCustomSet()

	//insert query

	mySet.Insert(7)
	mySet.Insert(18)

	for i := 0; i < 5; i++ {
		mySet.Insert(i)
	}

	fmt.Println("Set contains 1:", mySet.Find(7))  // true
	fmt.Println("Set contains 3:", mySet.Find(20)) // false

	//find

	mySet.Find(7)
	mySet.Find(20)

	//delete

	mySet.Remove(18)
	mySet.Remove(20)

	fmt.Println("Set after removal:", mySet.View())

}
