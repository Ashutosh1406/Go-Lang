package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...) //spread operator works on slices too in Golang
	result = append(result, right[j:]...)

	return result
}

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	midpoint := len(array) / 2
	left := mergeSort(array[:midpoint])  //till midPoint
	right := mergeSort(array[midpoint:]) //after midPoint

	return merge(left, right)
}

func printArray(array []int) {
	for _, value := range array {
		fmt.Println(value)
	}
	fmt.Println()
}

func main() {

	fmt.Println(`MERGE SORT IMPLEMENTATION in Golang`)

	for i := 0; i < 5; i++ {
		fmt.Println()
	}

	//[ for directly passing slice/array uncomment line no 53 ] and [comment all the lines from line no 55->97].
	//array := []int{12, 48, 36, 96, 120, 84, 24, 60, 72, 108}

	reader := bufio.NewReader(os.Stdin)

	//read the size
	fmt.Print("Enter the Size of Array: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(`Error Reading the Input`, err)
		return
	}

	input = strings.TrimSpace(input)
	size, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid Input.Please Enter an Integer")
		return
	}

	array := make([]int, size)
	fmt.Printf("Enter %d elements (space-separated):\n", size)

	inputElements, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(`Error Reading input`, err)
		return
	}

	inputElements = strings.TrimSpace(inputElements)
	elements := strings.Split(inputElements, " ")

	if len(elements) != size {
		fmt.Println("Number of elements does not match the array size provided by you")
		return
	}

	for i, element := range elements {
		array[i], err = strconv.Atoi(element)
		if err != nil {
			fmt.Println("Invalid Input.Please enter an integer", err)
			return
		}
	}

	fmt.Println("Given Array is")
	printArray(array)

	sortedArray := mergeSort(array)

	fmt.Println("Sorted Array is")
	printArray(sortedArray)
}
