package main

import (
	"AdventOfCode/helpers"
	"bufio"
	"fmt"
)

// Directions: up, down, left, right, and 4 diagonals
var dirs = [][2]int{
	{-1, 0},  // up
	{1, 0},   // down
	{0, -1},  // left
	{0, 1},   // right
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

// countAdjacentRolls counts how many adjacent @ symbols a cell has
func countAdjacentRolls(matrix [][]string, i, j int) int {
	adjacentAt := 0
	for _, dir := range dirs {
		ni := i + dir[0]
		nj := j + dir[1]
		if ni >= 0 && ni < len(matrix) && nj >= 0 && nj < len(matrix[ni]) {
			if matrix[ni][nj] == "@" {
				adjacentAt++
			}
		}
	}
	return adjacentAt
}

// findRemovableRolls finds all rolls that can be removed (have < 4 adjacent rolls)
func findRemovableRolls(matrix [][]string) [][2]int {
	removable := make([][2]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "@" {
				if countAdjacentRolls(matrix, i, j) < 4 {
					removable = append(removable, [2]int{i, j})
				}
			}
		}
	}
	return removable
}

func main() {
	file, err := helpers.GetFile("2025/04/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]string, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		matrixLine := make([]string, 0)

		for _, c := range scanner.Text() {
			matrixLine = append(matrixLine, string(c))
		}

		matrix = append(matrix, matrixLine)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Part 1: Count rolls removable in first pass
	firstPassRolls := len(findRemovableRolls(matrix))
	fmt.Println("Part 1 - Rolls removable in first pass:", firstPassRolls)

	// Part 2: Keep removing rolls until no more can be removed
	totalRemoved := 0
	for {
		removable := findRemovableRolls(matrix)
		if len(removable) == 0 {
			break
		}

		// Remove all accessible rolls
		for _, pos := range removable {
			matrix[pos[0]][pos[1]] = "."
		}
		totalRemoved += len(removable)

		fmt.Printf("Removed %d rolls this pass\n", len(removable))
	}

	fmt.Println("Part 2 - Total rolls removed:", totalRemoved)
}
