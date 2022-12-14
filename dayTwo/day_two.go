package dayTwo

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/common"
	"strings"
)

type moveInformation struct {
	points     int
	equivalent string
	win        string
}

// CalculateRockPaperScissorsScore calculates the total score of a
// Rock, Paper, Scissor player's rounds based on their move and the
// result of the round
func CalculateRockPaperScissorsScore(filename string) (int, error) {
	input, err := common.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	moveMap := map[string]moveInformation{
		"X": {1, "A", "C"},
		"Y": {2, "B", "A"},
		"Z": {3, "C", "B"},
	}

	movePoints := 0
	resultPoints := 0
	for _, moves := range input {
		arr := strings.Split(moves, " ")
		key := arr[1]
		points := moveMap[key].points

		movePoints += points

		switch true {
		case arr[0] == moveMap[arr[1]].equivalent:
			resultPoints += 3
			break
		case arr[0] == moveMap[arr[1]].win:
			resultPoints += 6
			break
		default:
			resultPoints += 0
		}
	}

	return resultPoints + movePoints, nil
}

type fullMoveInformation struct {
	points int
	win    string
	loss   string
	draw   string
}

// CalculateFinalRockPaperScissorsScore determines the sum of a player's
// moves and the result of each round, based on the player's strategy
func CalculateFinalRockPaperScissorsScore(filename string) (int, error) {
	input, err := common.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	moveMap := map[string]fullMoveInformation{
		"A": {1, "C", "B", "A"},
		"B": {2, "A", "C", "B"},
		"C": {3, "B", "A", "C"},
	}

	movePoints := 0
	resultPoints := 0
	for _, moves := range input {
		arr := strings.Split(moves, " ")

		key := moveMap[arr[0]]
		switch true {
		case arr[1] == "X":
			movePoints += moveMap[key.win].points
			resultPoints += 0
			break
		case arr[1] == "Y":
			movePoints += moveMap[key.draw].points
			resultPoints += 3
			break
		case arr[1] == "Z":
			movePoints += moveMap[key.loss].points
			resultPoints += 6
			break
		default:
			fmt.Errorf("Unknown move detected: %v", arr[0])
		}
	}
	return resultPoints + movePoints, nil
}
