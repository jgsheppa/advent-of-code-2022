package dayTwo

import (
	"os"
	"strings"
)

func readFile(filename string) ([]string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	output := strings.TrimSpace(string(b))
	return strings.Split(output, "\n"), nil
}

type MoveInformation struct {
	points     int
	equivalent string
	win        string
}

func CalculateRockPaperScissorsScore(filename string) (int, error) {
	input, err := readFile(filename)
	if err != nil {
		return 0, err
	}

	moveMap := map[string]MoveInformation{
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
