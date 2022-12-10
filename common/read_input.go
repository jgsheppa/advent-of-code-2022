package common

import (
	"os"
	"strings"
)

// ReadFile reads a file, splits the data into
// an array of strings by line break, and returns
// this array
func ReadFile(filename string) ([]string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	output := string(b)
	return strings.Split(output, "\n"), nil
}

func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}
