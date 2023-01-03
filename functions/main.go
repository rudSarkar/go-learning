package main

import "fmt"

// SimpleFunction print "Hello World" message
func SimpleFunction() {
	fmt.Println("Hello World")
}

/*
Simple with parameters and return type
add takes two integer parameter a and b
*/
func add(a int, b int) int /* return type */ {
	return a + b
}

func main() {
	SimpleFunction()
	fmt.Println(add(10, 20))
}
