package dayOne

import (
	"os"
	"sort"
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

func sumCaloriesCarried(input string) ([]int, error) {
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
				return nil, err
			}
			count += i
			elfCaloricSums = append(elfCaloricSums, count)
			continue
		}
		i, err := strconv.Atoi(ch)
		if err != nil {
			return nil, err
		}
		count += i
	}
	return elfCaloricSums, nil
}

func FindMaxElfCalories(filename string) (int, error) {
	input, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}

	elfCaloricSums, err := sumCaloriesCarried(input)
	if err != nil {
		return 0, nil
	}

	mostCalories := findTopElf(elfCaloricSums)
	return mostCalories, nil
}

func findTopElf(elfCaloricSums []int) int {
	maxCaloriesCarried := 0
	for _, i := range elfCaloricSums {
		if i > maxCaloriesCarried {
			maxCaloriesCarried = i
		}
	}
	return maxCaloriesCarried
}

func FindSumTopThreeElfTotalCalories(filename string) (int, error) {
	input, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}

	elfCaloricSums, err := sumCaloriesCarried(input)
	if err != nil {
		return 0, nil
	}

	topThree := findTopThreeElves(elfCaloricSums)
	topThreeSum := 0
	for _, i := range topThree {
		topThreeSum += i
	}
	return topThreeSum, nil
}

func findTopThreeElves(elfCaloricSums []int) []int {
	sort.Ints(elfCaloricSums)

	return elfCaloricSums[len(elfCaloricSums)-3:]
}
