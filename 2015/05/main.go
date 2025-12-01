package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"io"
	"strings"
)

func main() {
	file, err := helpers.GetFile("2015/05/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	fmt.Println(partOne_IsNice(lines))
	fmt.Println(partTwo_IsNice(lines))
}

func partOne_IsNice(input []string) int {
	var (
		totalCount int
	)

	for _, line := range input {
		var (
			vowelMatch  bool
			doubleMatch bool
			repeatMatch bool
		)

		if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
			doubleMatch = true
		}

		if getVowelCount(line) > 2 {
			vowelMatch = true
		}

		if getRepeatMatch(line) {
			repeatMatch = true
		}

		if vowelMatch && !doubleMatch && repeatMatch {
			totalCount++
		}
	}

	return totalCount
}

func getVowelCount(input string) int {
	var vowelCount int

	for _, char := range input {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowelCount++
		}
	}

	return vowelCount
}

func getRepeatMatch(input string) bool {
	for idx := 0; idx < len(input)-1; idx++ {
		if input[idx] == input[idx+1] {
			return true
		}
	}

	return false
}

func partTwo_IsNice(input []string) int {
	var totalCount int

	for _, line := range input {
		if hasRepeatedPair(line) && hasRepeatWithGap(line) {
			totalCount++
		}
	}

	return totalCount
}

// Rule 1: pair of two letters appears at least twice without overlapping
func hasRepeatedPair(s string) bool {
	pairMap := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]

		// If we've seen this pair before and it doesn't overlap
		if idx, exists := pairMap[pair]; exists {
			if i-idx > 1 {
				return true
			}
		} else {
			pairMap[pair] = i
		}
	}

	return false
}

// Rule 2: at least one letter repeats with exactly one in between (xyx pattern)
func hasRepeatWithGap(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
