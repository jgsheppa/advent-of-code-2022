package main

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/dayOne"
	"github.com/jgsheppa/advent-of-code-2022/dayTwo"
)

func main() {
	_, err := dayOne.FindSumTopThreeElfTotalCalories("puzzle_input_1.txt")
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}

	totalPoints, err := dayTwo.CalculateRockPaperScissorsScore("puzzle_input_2.txt")
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}
	fmt.Println(totalPoints)
}
