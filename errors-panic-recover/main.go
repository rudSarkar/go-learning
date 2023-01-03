package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// In Go, errors are represented using the error type, which is an interface type defined as follows:

	/*
		type error interface {
			Error() string
		}
	*/

	// error example with file operatinon
	file_name, err := os.Open("hello.txt")

	// error handling
	if err != nil {
		log.Fatalln(err)
	}

	// used defer to close the file at the end
	defer file_name.Close()

	fmt.Println(file_name.Name()) // print file name just for example

	// defer example
	fmt.Println("Hello")
	defer fmt.Println("Goodbye") // this line will be executed last
	fmt.Println("World")

	// panic example
	// The panic function in Go is used to halt the normal execution of a program and print a message to the console.

	result := panic_div(10, 0)
	fmt.Println(result)

	result2 := recover_div(10, 0)
	fmt.Println(result2) // this line will not be reached

	/*
		In this example, the div function divides x by y. If y is zero, it panics with a message saying "division by zero". If y is not zero, it returns the result of the division.

		The defer statement at the beginning of the div function calls a function that calls recover. If a panic occurs within the div function, the recover function will catch it and print a message saying "Recovered from panic: division by zero".

		If you run this code, it will print "Recovered from panic: division by zero" and continue execution. The line fmt.Println(result) will not be reached, because the div function returns before that line is executed.

		Reference: recover_div()
	*/

}

func panic_div(x, y int) int {
	if y == 0 {
		panic("division by zero")
	}
	return x / y
}

func recover_div(x, y int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	if y == 0 {
		panic("division by zero")
	}
	return x / y
}
