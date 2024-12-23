// https://adventofcode.com/2024/day/1

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseDay1Input(input string) (left []int, right []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		parts := strings.Fields(line)

		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	return left, right
}

func day1_1(input string) {
	left, right := parseDay1Input(input)

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		if distance < 0 {
			distance = distance * -1
		}
		totalDistance += distance
	}
	fmt.Println(totalDistance)
}

func day1_2(input string) {
	left, right := parseDay1Input(input)

	countRightMap := make(map[int]int)

	for _, num := range right {
		countRightMap[num]++
	}

	similarityScore := 0

	for _, num := range left {
		similarityScore += num * countRightMap[num]
	}

	fmt.Println(similarityScore)
}
