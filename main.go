package main

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/dayFive"
)

func main() {
	topCrates, err := dayFive.FindTopCrates("puzzle_input_5.txt", "puzzle_input_5_map.txt")
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}
	fmt.Printf("FindTopCrates: %v \n", topCrates)
}
