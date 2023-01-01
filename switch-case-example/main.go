package main

import (
	"fmt"
	"time"
)

func main() {
	// switch case example
	var x int = 2

	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	// working with time.Now()
	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("Today is weekend")
	default:
		fmt.Println("Today is workday")
	}
}
