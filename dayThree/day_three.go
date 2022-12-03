package dayThree

import (
	"github.com/jgsheppa/advent-of-code-2022/common"
	"golang.org/x/exp/maps"
)

func createAlphabetToPriorityMap() map[string]int {
	lowercaseLetters := make(map[string]int)
	uppercaseLetters := make(map[string]int)

	count := 1
	for i := 'a'; i <= 'z'; i++ {
		lowercaseLetters[string(i)] = count
		count++
	}

	for i := 'A'; i <= 'Z'; i++ {
		uppercaseLetters[string(i)] = count
		count++
	}

	maps.Copy(lowercaseLetters, uppercaseLetters)

	return lowercaseLetters
}

func CalculateSumOfDuplicateItemPriority(filename string) (int, error) {
	rucksacks, err := common.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	letterPriorityMap := createAlphabetToPriorityMap()
	score := 0

	for _, ch := range rucksacks {
		m := make(map[string]bool)

		firstHalf := ch[:len(ch)/2]
		secondHalf := ch[len(ch)/2:]

		for _, item := range firstHalf {
			m[string(item)] = true
		}

		for _, item := range secondHalf {
			if _, ok := m[string(item)]; ok {
				score += letterPriorityMap[string(item)]
				break
			}
		}
	}

	return score, nil
}

func CalculateSumOfDuplicateElfBadges(filename string) (int, error) {
	return 0, nil
}
