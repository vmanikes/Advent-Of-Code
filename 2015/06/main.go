package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	Op     string
	X1, Y1 int
	X2, Y2 int
}

var re = regexp.MustCompile(`^(turn on|turn off|toggle)\s+(\d+),(\d+)\s+through\s+(\d+),(\d+)$`)

func ParseInstruction(s string) (*Instruction, error) {
	m := re.FindStringSubmatch(s)
	if m == nil {
		return nil, fmt.Errorf("invalid instruction: %q", s)
	}

	// convert all numbers
	x1, _ := strconv.Atoi(m[2])
	y1, _ := strconv.Atoi(m[3])
	x2, _ := strconv.Atoi(m[4])
	y2, _ := strconv.Atoi(m[5])

	return &Instruction{
		Op: m[1],
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}, nil
}

func main() {
	file, err := helpers.GetFile("2015/06/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	instructions := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		instruction, err := ParseInstruction(line)
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, *instruction)
	}

	fmt.Println(partOne(instructions))
	fmt.Println(partTwo(instructions))
}

func partOne(instructions []Instruction) int {
	matrix := make([][]int, 1000)
	for i := range matrix {
		matrix[i] = make([]int, 1000)
	}

	for _, instruction := range instructions {
		switch instruction.Op {
		case "turn on":
			for x := instruction.X1; x <= instruction.X2; x++ {
				for y := instruction.Y1; y <= instruction.Y2; y++ {
					matrix[x][y] = 1
				}
			}
		case "turn off":
			for x := instruction.X1; x <= instruction.X2; x++ {
				for y := instruction.Y1; y <= instruction.Y2; y++ {
					matrix[x][y] = 0
				}
			}
		case "toggle":
			for x := instruction.X1; x <= instruction.X2; x++ {
				for y := instruction.Y1; y <= instruction.Y2; y++ {
					if matrix[x][y] == 1 {
						matrix[x][y] = 0
					} else {
						matrix[x][y] = 1
					}
				}
			}
		}
	}

	count := 0

	for _, row := range matrix {
		for _, col := range row {
			if col == 1 {
				count++
			}
		}
	}

	return count
}

func partTwo(instructions []Instruction) int {
	matrix := make([][]int, 1000)
	for i := range matrix {
		matrix[i] = make([]int, 1000)
	}

	for _, instruction := range instructions {
		switch instruction.Op {
		case "turn on":
			for x := instruction.X1; x <= instruction.X2; x++ {
				for y := instruction.Y1; y <= instruction.Y2; y++ {
					matrix[x][y] += 1
				}
			}
		case "turn off":
			for x := instruction.X1; x <= instruction.X2; x++ {
				for y := instruction.Y1; y <= instruction.Y2; y++ {
					if matrix[x][y] > 0 {
						matrix[x][y] -= 1
					}
				}
			}
		case "toggle":
			for x := instruction.X1; x <= instruction.X2; x++ {
				for y := instruction.Y1; y <= instruction.Y2; y++ {
					matrix[x][y] += 2
				}
			}
		}
	}

	count := 0

	for _, row := range matrix {
		for _, col := range row {
			count += col
		}
	}

	return count
}
