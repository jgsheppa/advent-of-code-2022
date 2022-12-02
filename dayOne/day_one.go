package dayOne

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) (string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	output := string(b)
	return output, nil
}

func FindMaxElfCalories(filename string) (int, error) {
	input, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}

	caloriesArray := strings.Split(input, "\n")

	var elfCaloricSums []int
	count := 0

	for index, ch := range caloriesArray {
		if len(ch) == 0 {
			elfCaloricSums = append(elfCaloricSums, count)
			count = 0
			continue
		} else if index == len(caloriesArray)-1 {
			i, err := strconv.Atoi(ch)
			if err != nil {
				return 0, err
			}
			count += i
			elfCaloricSums = append(elfCaloricSums, count)
			continue
		}
		i, err := strconv.Atoi(ch)
		if err != nil {
			return 0, err
		}
		count += i
	}

	maxCaloriesCarried := 0
	for _, i := range elfCaloricSums {
		if i > maxCaloriesCarried {
			maxCaloriesCarried = i
		}
	}
	return maxCaloriesCarried, nil
}
