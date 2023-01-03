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

// return values of a function can be named
func rectangle(length int, width int) (area int) {
	perimeter := 2 * (length + width)
	fmt.Println("Perimeter: ", perimeter)

	area = length * width

	return
}

// return values of a function can be named and it can be multiple
func rectangle_multiple_return(length int, width int) (area int, perimeter int) {

	perimeter = 2 * (length + width)
	area = length * width

	return
}

func main() {
	SimpleFunction()
	fmt.Println(add(10, 20))

	fmt.Println("Area: ", rectangle(20, 30))

	// multiple return value extract
	var num1, num2 int
	num1, num2 = rectangle_multiple_return(20, 40)

	fmt.Println("Area:", num1)
	fmt.Println("Perimeter:", num2)
}
