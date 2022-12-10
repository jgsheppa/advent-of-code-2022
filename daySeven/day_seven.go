package daySeven

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/common"
	"math/rand"
	"strconv"
	"strings"
)

type Directory struct {
	memory    int
	traversed []int
}

func FindDirectoryWithMostMemory(filename string) (int, error) {
	commands, err := common.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	paths := make(map[string]int)

	directories, _ := walkDir(commands, paths)

	var totalMemory int

	for _, mem := range directories {
		if mem <= 100000 {
			totalMemory += mem
		}
	}

	return totalMemory, nil
}

func walkDir(commands []string, dirStructure map[string]int) (map[string]int, error) {
	paths := []string{"/"}
	rand.Seed(2324930403)
	for index := 1; index < len(commands); index++ {
		current := strings.Split(commands[index], " ")

		if current[1] == "cd" && current[2] == ".." {
			paths = paths[:len(paths)-1]
			continue
		}

		if current[0] == "$" {
			if current[1] == "cd" {
				dir := current[2]
				if _, ok := dirStructure[dir]; ok {
					dir = current[2] + strconv.Itoa(rand.Int())
				}
				paths = append(paths, dir)
			}
			continue
		}

		if current[0] != "dir" {
			size, err := strconv.Atoi(current[0])
			if err != nil {
				return nil, err
			}
			for _, p := range paths {
				dirStructure[p] += size
			}
		}
	}
	return dirStructure, nil
}

func DeleteDirectory(filename string) (int, error) {
	totalStorage := 70000000
	neededStorage := 30000000

	commands, err := common.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	paths := make(map[string]int)

	directories, _ := walkDir(commands, paths)

	currentStorage := directories["/"]
	remainingStorage := totalStorage - currentStorage

	var potentialDirectories []int
	for _, dir := range directories {
		if dir+remainingStorage >= neededStorage {
			fmt.Printf("directory to delete: %v \n", dir)
			potentialDirectories = append(potentialDirectories, dir)
		}
	}

	min := 0
	for i, size := range potentialDirectories {
		if i == 0 || size < min {
			min = size
		}
	}

	return min, nil
}
