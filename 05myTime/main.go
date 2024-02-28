package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study in Golang")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //only this value is allowed

	createdDate := time.Date(2024, time.June, 12, 23, 23, 0, 0, presentTime.Location())

	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
