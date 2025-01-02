package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func getInput(day int) (string, error) {
	const inputDir string = "inputs"
	inputFile := filepath.Join(inputDir, fmt.Sprintf("day-%d", day))
	if _, err := os.Stat(inputFile); err == nil {
		content, _ := os.ReadFile(inputFile)
		return string(content), nil
	}

	return "", errors.New(fmt.Sprintf("Problem reading input file: %s", inputFile))
}

func main() {
	day := 1
	input, err := getInput(day)

	if err != nil {
		panic(err)
	}

	day1_1(input)
	day1_2(input)
}
