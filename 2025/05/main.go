package main

import (
	"AdventOfCode/helpers"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := helpers.GetFile("2025/05/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges := make([]string, 0)
	ingredients := make([]string, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		value := scanner.Text()

		if strings.Contains(value, "-") {
			ranges = append(ranges, value)
		} else if len(value) > 0 {
			ingredients = append(ingredients, value)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(partOne(ranges, ingredients))
}

type Interval struct {
	Start int
	End   int
}

func partOne(ranges []string, ingredients []string) int {
	// ---- Step 1: parse ----
	intervals := make([]Interval, 0, len(ranges))
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		intervals = append(intervals, Interval{a, b})
	}

	fmt.Println(partTwo(intervals))

	// ---- Step 2: sort by start ----
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	fmt.Println("sorted intervals: ", intervals)

	// ---- Step 3: merge ----
	merged := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := &merged[len(merged)-1]
		curr := intervals[i]

		if curr.Start <= last.End+1 {
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			merged = append(merged, curr)
		}
	}

	fmt.Println("merged: ", merged)

	// ---- Step 4: Binary search for each ingredient ----
	fresh := 0
	for _, ing := range ingredients {
		id, _ := strconv.Atoi(ing)

		// binary search in merged ranges
		idx := sort.Search(len(merged), func(i int) bool {
			return merged[i].Start > id
		})

		if idx > 0 && merged[idx-1].Start <= id && id <= merged[idx-1].End {
			fresh++
		}
	}

	return fresh
}

func partTwo(intervals []Interval) int {
	// Step 1 — sort by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// Step 2 — merge
	merged := []Interval{intervals[0]}
	for _, curr := range intervals[1:] {
		last := &merged[len(merged)-1]

		if curr.Start <= last.End+1 {
			// Overlapping or touching
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			// No overlap
			merged = append(merged, curr)
		}
	}

	// Step 3 — sum all merged interval lengths
	total := 0
	for _, iv := range merged {
		total += (iv.End - iv.Start + 1)
	}

	return total
}
