package main

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/dayFour"
	"github.com/jgsheppa/advent-of-code-2022/dayOne"
)

func main() {
	_, err := dayOne.FindSumTopThreeElfTotalCalories("puzzle_input_1.txt")
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}

	sum, err := dayFour.CalculateOverlappingSections("puzzle_input_4.txt")
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}
	fmt.Printf("CalculateOverlappingSections: %v \n", sum)
}
