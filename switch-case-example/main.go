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

	// working with interface
	// In Go, the interface{} type is a special type that can hold any value. It is a bit similar to the Object type in Java or the dynamic type in C#. You can use interface{} when you want to write a function that can accept values of any type.

	whatType := func(ty interface{}) {
		switch t := ty.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Printf("The type of data is %T\n", t)
		}
	}

	whatType(true)
	whatType(1)
	whatType("hey")
}
