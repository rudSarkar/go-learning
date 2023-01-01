package main

import (
	"fmt"
)

func main() {

	if 5/2 == 0 {
		fmt.Println("Yes the output is 0")
	} else {
		fmt.Println("The output isn't 0")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
