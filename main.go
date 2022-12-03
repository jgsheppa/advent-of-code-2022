package main

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/dayOne"
	"github.com/jgsheppa/advent-of-code-2022/dayThree"
)

func main() {
	_, err := dayOne.FindSumTopThreeElfTotalCalories("puzzle_input_1.txt")
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}

	sum, err := dayThree.CalculateSumOfDuplicateElfBadges("puzzle_input_3.txt")
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}
	fmt.Printf("CalculateSumOfDuplicateItemPriority: %v", sum)
}
