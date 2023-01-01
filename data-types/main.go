package main

import (
	"fmt"
)

func main() {

	/*
		"string" type represents a sequence of Unicode characters.
	*/
	var xs string = "A string"
	fmt.Println(xs)

	/*
		int, int8, int16, int32, int64: These types represent signed integers of different sizes. A signed integer is a whole number that can be positive, negative, or zero.
	*/

	var x int = 10
	var y int8 = -128
	var z int32 = -2147483648

	fmt.Println("This is signed numbers (int)")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	/*

		1. uint: The uint type represents an unsigned integer that is at least 32 bits in size. The size of the uint type may vary depending on the platform, but it is guaranteed to be at least 32 bits. The range of values that the uint type can represent is from 0 to 2^32-1 on a 32-bit platform, or from 0 to 2^64-1 on a 64-bit platform.
		2. uint8: The uint8 type represents an unsigned 8-bit integer. The range of values that the uint8 type can represent is from 0 to 255.
		3. uint16: The uint16 type represents an unsigned 16-bit integer. The range of values that the uint16 type can represent is from 0 to 65535.
		4. uint32: The uint32 type represents an unsigned 32-bit integer. The range of values that the `uint
	*/

	var xu uint = 10
	var yu uint8 = 255
	var zu uint32 = 4294967295
	fmt.Println("This is unsigned numbers (uint)")
	fmt.Println(xu)
	fmt.Println(yu)
	fmt.Println(zu)

	/*
		You should use the float32 type if you need to store numbers with a fractional component and you want to use less memory.
		You should use the float64 type if you need greater precision or if you are working with larger values.
	*/

	var xf float32 = 3.5
	var yf float64 = 1.234e-10

	fmt.Println("This is float numbers (float)")
	fmt.Println(xf)
	fmt.Println(yf)

	fmt.Println("This is rune")
	/*
		Example of rune types
	*/
	xru := "Hello, 世界"

	fmt.Printf("Type of the variable is: %T\n", xru)
	for _, i := range xru {
		fmt.Printf("%q", i)
	}

	fmt.Println("\nThis is array with modify and short hand")
	/*
		Array example
	*/
	var aint [4]int = [4]int{1, 2, 3, 4}
	var astr [4]string = [4]string{"a", "b", "c", "d"}

	fmt.Println(aint)
	fmt.Println(astr)

	// modify element
	astr[3] = "e"

	fmt.Println(astr)

	// Declare and initialize an array using short syntax
	bint := [3]int{5, 6, 7}

	// Print the array
	fmt.Println(bint)

	fmt.Println("Byte example")
	// Declare and initialize a byte slice
	b := []byte("hello")

	// Print the byte slice
	fmt.Println(b)

	// Modify an element of the byte slice
	b[0] = 'H'

	// Print the modified byte slice
	fmt.Println(b)

	// Convert the byte slice to a string
	s := string(b)

	// Print the string
	fmt.Println(s)

	fmt.Println("Complex Datatype example")
	var7 := complex(9, 15)

	fmt.Println(var7)
}
