package main

import (
	"AdventOfCode/helpers"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	file, err := helpers.GetFile("2018/01/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	sum := 0
	sumMap := make(map[int]struct{})

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		sum += value

		if _, ok := sumMap[sum]; ok {
			fmt.Println("Duplicate value:", value)
		} else {
			sumMap[value] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
