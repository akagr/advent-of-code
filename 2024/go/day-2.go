package main

import (
	"strconv"
	"strings"
)

func parseDay2Input(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][]int, len(lines))
	for i, line := range lines {
		levels := strings.Split(strings.TrimSpace(line), " ")
		result[i] = make([]int, len(levels))
		for j, level := range levels {
			result[i][j], _ = strconv.Atoi(level)
		}
	}
	return result
}
