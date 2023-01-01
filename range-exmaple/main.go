package main

import (
	"fmt"
)

func main() {
	// creating map
	fruit_lists := map[int]string{1: "Apple", 2: "Banana", 3: "Cherry", 4: "Mango"}

	for index, fruit := range fruit_lists {
		fmt.Printf("Index ID: %d, Fruit Name: %s\n", index, fruit)
	}
}
