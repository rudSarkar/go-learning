package main

import (
	"fmt"
)

func main() {

	// three-component loop
	fmt.Println("Three-component loop")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for t := 1; t <= 4; t++ {
		fmt.Println(t)
	}

	// while loop
	fmt.Println("While loop")
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)

	// decrealing slice of string
	fmt.Println("For-each range loop")
	fruit_slice := []string{"Apple", "Banana", "Cherry"}

	for index, fruit := range fruit_slice {
		fmt.Printf("index: %d, value: %s\n", index, fruit)
	}

	// creating a map of string
	fruit_map := map[int]string{1: "Apple", 2: "Banana", 3: "Guava"}

	for indexName, fruitName := range fruit_map {
		fmt.Printf("Index id: %d, Fruit Name: %s\n", indexName, fruitName)
	}
}
