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

}

func panic_div(x, y int) int {
	if y == 0 {
		panic("division by zero")
	}
	return x / y
}
