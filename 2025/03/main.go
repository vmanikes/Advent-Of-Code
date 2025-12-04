package main

import (
	"AdventOfCode/helpers"
	"bufio"
	"fmt"
)

func main() {
	file, err := helpers.GetFile("2025/03/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	sum := 0

	for _, line := range lines {
		sum += getMaxNumber(line)
	}

	fmt.Println("part one", sum)
}

func partTwo(lines []string) {
	var sum uint64 = 0

	for _, line := range lines {
		sum += getMaxNumberK(line, 12)
	}

	fmt.Println("part one", sum)
}

func getMaxNumber(line string) int {
	n := len(line)
	digits := make([]int, n)
	for i, c := range line {
		digits[i] = int(c - '0')
	}

	// Build suffix max array
	suffixMax := make([]int, n+1) // last element = 0
	for i := n - 1; i >= 0; i-- {
		suffixMax[i] = suffixMax[i+1]
		if digits[i] > suffixMax[i] {
			suffixMax[i] = digits[i]
		}
	}

	// Find best two-digit number
	maxVal := 0
	for i := 0; i < n-1; i++ {
		right := suffixMax[i+1]
		val := digits[i]*10 + right
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

func getMaxNumberK(line string, k int) uint64 {
	n := len(line)
	if k <= 0 || n == 0 {
		return 0
	}
	if k >= n {
		// take all digits
		var r uint64
		for i := 0; i < n; i++ {
			r = r*10 + uint64(line[i]-'0')
		}
		return r
	}

	stack := make([]byte, 0, k)
	toRemove := n - k

	for i := 0; i < n; i++ {
		c := line[i]
		for len(stack) > 0 && stack[len(stack)-1] < c && toRemove > 0 {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, c)
	}

	// keep first k digits
	stack = stack[:k]

	// build integer result (fits in 64-bit for k <= 18; k=12 is safe)
	var res uint64 = 0
	for i := 0; i < k; i++ {
		res = res*10 + uint64(stack[i]-'0')
	}
	return res
}
