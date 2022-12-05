package dayFive

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/common"
	"sort"
	"strconv"
	"unicode"
)

type CrateColumn struct {
	crates []string
	index  int
}

type CrateDirection struct {
	count       int
	origin      string
	destination string
}

func MoveCrate(origin, dest []string, count int) ([]string, []string) {
	var ret []string

	if len(origin)+1 == count {
		ret = append(ret, origin[:]...)
	} else if len(origin) > 0 {
		fmt.Printf("origin length: %v count: %v \n", len(origin), count)
		ret = append(ret, origin[count:]...)
	}

	var ter []string
	if len(origin) > 0 {
		for i := count - 1; i > -1; i-- {
			ter = append(ter, origin[i])
		}
	}
	ter = append(ter, dest[:]...)
	return ret, ter
}

func FindTopCrates(filename, crateMap string) (string, error) {
	moves, err := common.ReadFile(filename)
	if err != nil {
		return "", err
	}
	crateDirections, err := transformCrateMovementDirections(moves)
	if err != nil {
		return "", err
	}

	crates, err := common.ReadFile(crateMap)
	if err != nil {
		return "", err
	}

	columnCount := 0
	var columnIds []string

	crateTracker, columnIdsFull := transformInitialCrateMap(crates, columnIds, columnCount)

	// Let's get a slice of all keys
	keySlice := make([]int, 0)
	for key, _ := range crateDirections {
		keySlice = append(keySlice, key)
	}

	// Now sort the slice
	sort.Ints(keySlice)

	// Iterate over all keys in a sorted order
	for _, key := range keySlice {
		direction := crateDirections[key]

		origin, destination := MoveCrate(crateTracker[direction.origin].crates, crateTracker[direction.destination].crates, direction.count)

		crateTracker[direction.origin] = CrateColumn{origin, crateTracker[direction.origin].index}
		crateTracker[direction.destination] = CrateColumn{destination, crateTracker[direction.destination].index}

	}

	var result string

	for i := 0; i < len(columnIdsFull); i++ {
		key := columnIdsFull[i]
		if len(crateTracker[key].crates) > 0 {
			result += crateTracker[key].crates[0]
		}
	}

	return result, nil
}

func transformCrateMovementDirections(moves []string) (map[int]CrateDirection, error) {
	directions := make(map[int]CrateDirection)

	for index, move := range moves {
		var movement []string
		for _, ch := range move {
			if unicode.IsNumber(ch) {
				movement = append(movement, string(ch))
			}
		}
		count, err := strconv.Atoi(movement[0])
		if err != nil {
			return nil, err
		}
		origin := movement[1]
		if err != nil {
			return nil, err
		}
		dest := movement[2]
		if err != nil {
			return nil, err
		}
		directions[index] = CrateDirection{count, origin, dest}
	}

	return directions, nil
}

func transformInitialCrateMap(crates, columns []string, columnNumber int) (map[string]CrateColumn, []string) {
	crateTracker := make(map[string]CrateColumn)

	for index, ch := range crates[len(crates)-1] {
		if unicode.IsNumber(ch) {
			var initialColumnState []string
			crateTracker[string(ch)] = CrateColumn{initialColumnState, index}
			columns = append(columns, string(ch))
			columnNumber++
		}
	}

	for i := 0; i < len(crates)-1; i++ {
		tempSlice := make([]string, len(crates)-1)
		for j := 1; j <= columnNumber; j++ {
			indexAsString := strconv.Itoa(j)
			currentColumn := crateTracker[indexAsString]

			if len(crates[i]) >= currentColumn.index && unicode.IsLetter(rune(crates[i][currentColumn.index])) {
				tempSlice = append(currentColumn.crates, string(crates[i][currentColumn.index]))
				crateTracker[indexAsString] = CrateColumn{tempSlice, currentColumn.index}

			}

		}
	}

	return crateTracker, columns
}
