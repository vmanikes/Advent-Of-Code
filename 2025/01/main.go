package main

import (
	"AdventOfCode/2025/helpers"
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	MaxRingSize = 100
	StartValue  = 50
)

func main() {
	file, err := helpers.GetFile("01/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := make([]int, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if strings.HasPrefix(scanner.Text(), "L") {
			pointer := strings.ReplaceAll(scanner.Text(), "L", "")

			pointerValue, err := strconv.Atoi(pointer)
			if err != nil {
				log.Fatalln("unable to convert value to int: ", err)
			}
			values = append(values, -pointerValue)
		}

		if strings.HasPrefix(scanner.Text(), "R") {
			pointer := strings.ReplaceAll(scanner.Text(), "R", "")

			pointerValue, err := strconv.Atoi(pointer)
			if err != nil {
				log.Fatalln("unable to convert value to int: ", err)
			}
			values = append(values, pointerValue)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	partOne(values)
	partTwo(values)
}

func partOne(values []int) {
	r := ring.New(MaxRingSize)

	for i := 0; i < MaxRingSize; i++ {
		r.Value = i
		r = r.Next()
	}

	r = r.Move(StartValue)

	key := 0

	for _, value := range values {
		r = r.Move(value)

		if r.Value.(int) == 0 {
			key++
		}
	}

	fmt.Println("Solution 1:", key)
}

func partTwo(values []int) {
	r := ring.New(MaxRingSize)

	for i := 0; i < MaxRingSize; i++ {
		r.Value = i
		r = r.Next()
	}

	r = r.Move(StartValue)

	key := 0

	for _, value := range values {
		for i := 0; i < int(math.Abs(float64(value))); i++ {
			if value < 0 {
				r = r.Prev()
			} else {
				r = r.Next()
			}

			if r.Value.(int) == 0 {
				key++
			}
		}
	}

	fmt.Println("Solution 2:", key)
}
