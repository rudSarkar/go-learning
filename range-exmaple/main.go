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

	// creating array exmple with no index number in range
	num_list := []int{1, 2, 3, 4, 5, 6, 7}

	for _, num := range num_list {
		fmt.Printf("%d", num)
	}
}
