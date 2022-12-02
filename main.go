package main

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/dayOne"
)

func main() {
	maxCalories, err := dayOne.FindMaxElfCalories("puzzle_input_1.txt")
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}
	fmt.Println(maxCalories)
}
