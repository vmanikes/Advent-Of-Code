package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := helpers.GetFile("2015/03/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

func partOne(input string) int {
	var (
		row    = 0
		column = 0
	)

	result := make(map[string]int)
	result[fmt.Sprintf("%d,%d", row, column)] = 1

	for _, val := range input {
		switch string(val) {
		case ">":
			column++
		case "<":
			column--
		case "^":
			row--
		case "v":
			row++
		}

		result[fmt.Sprintf("%d,%d", row, column)] = 1
	}

	return len(result)
}

func partTwo(input string) int {
	var (
		robotRow    = 0
		robotColumn = 0
		santaRow    = 0
		santaColumn = 0
	)

	result := make(map[string]int)
	result[fmt.Sprintf("%d,%d", robotRow, robotColumn)] = 1
	result[fmt.Sprintf("%d,%d", santaRow, santaColumn)] = 1

	for idx, val := range input {
		if idx%2 == 0 {
			switch string(val) {
			case ">":
				santaColumn++
			case "<":
				santaColumn--
			case "^":
				santaRow--
			case "v":
				santaRow++
			}

			result[fmt.Sprintf("%d,%d", santaRow, santaColumn)] += 1
		} else {
			switch string(val) {
			case ">":
				robotColumn++
			case "<":
				robotColumn--
			case "^":
				robotRow--
			case "v":
				robotRow++
			}

			result[fmt.Sprintf("%d,%d", robotRow, robotColumn)] += 1
		}
	}

	return len(result)
}
