package main

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/daySeven"
)

func main() {

	memory, err := daySeven.DeleteDirectory("puzzle_input_day_7.txt")
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}
	fmt.Printf("FindDirectoryWithMostMemory: %v \n", memory)
}
