package common

import (
	"os"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	output := string(b)
	return strings.Split(output, "\n"), nil
}
