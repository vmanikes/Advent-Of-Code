package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("2025/06/input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rows := len(lines) - 1 // all but last row (data rows)

	// Find max width
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	// Ensure all rows have same width (pad if necessary)
	for i := range lines {
		for len(lines[i]) < width {
			lines[i] += " "
		}
	}
	opsRow := lines[len(lines)-1]

	// Process columns right-to-left
	answer := 0
	col := width - 1

	for col >= 0 {
		// Skip separator columns (all spaces in data rows AND operator is space)
		for col >= 0 && isDataSpace(lines, rows, col) && opsRow[col] == ' ' {
			col--
		}
		if col < 0 {
			break
		}

		// Find the operator for this problem (scan leftward until we find it)
		var op byte
		for c := col; c >= 0; c-- {
			if opsRow[c] == '+' || opsRow[c] == '*' {
				op = opsRow[c]
				break
			}
			if isDataSpace(lines, rows, c) && opsRow[c] == ' ' {
				break
			}
		}

		// Collect numbers for this problem (right-to-left columns)
		var nums []int

		for col >= 0 {
			// Check if this is a separator column
			if isDataSpace(lines, rows, col) && opsRow[col] == ' ' {
				break
			}

			// Read this column top-to-bottom to form a number
			num := 0
			hasDigit := false
			for r := 0; r < rows; r++ {
				ch := lines[r][col]
				if ch >= '0' && ch <= '9' {
					num = num*10 + int(ch-'0')
					hasDigit = true
				}
			}

			if hasDigit {
				nums = append(nums, num)
			}
			col--
		}

		// Compute result for this problem
		if len(nums) > 0 && op != 0 {
			res := nums[0]
			for j := 1; j < len(nums); j++ {
				if op == '+' {
					res += nums[j]
				} else {
					res *= nums[j]
				}
			}
			answer += res
		}
	}

	fmt.Println(answer)
}

func isDataSpace(lines []string, rows int, col int) bool {
	for r := 0; r < rows; r++ {
		if col < len(lines[r]) && lines[r][col] != ' ' {
			return false
		}
	}
	return true
}
