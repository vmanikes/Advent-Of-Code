package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := `2558912-2663749,1-19,72-85,82984-100358,86-113,193276-237687,51-69,779543-880789,13004-15184,2768-3285,4002-4783,7702278-7841488,7025-8936,5858546565-5858614010,5117615-5149981,4919-5802,411-466,126397-148071,726807-764287,7454079517-7454227234,48548-61680,67606500-67729214,9096-10574,9999972289-10000034826,431250-455032,907442-983179,528410-680303,99990245-100008960,266408-302255,146086945-146212652,9231222-9271517,32295166-32343823,32138-36484,4747426142-4747537765,525-652,333117-414840,13413537-13521859,1626-1972,49829276-50002273,69302-80371,8764571787-8764598967,5552410836-5552545325,660-782,859-1056`

	ranges := strings.Split(input, ",")

	invalidSum := 0

	for _, r := range ranges {

		rSplit := strings.Split(r, "-")
		if len(rSplit) != 2 {
			panic("invalid ranges")
		}

		start, err := strconv.Atoi(rSplit[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(rSplit[1])
		if err != nil {
			panic(err)
		}

		for idx := start; idx <= end; idx++ {
			if !isNumberValid2(idx) {
				invalidSum += idx
			}
		}
	}

	fmt.Println(invalidSum)
}

func isNumberValid(number int) bool {
	digitCount := 0
	for x := number; x > 0; x /= 10 {
		digitCount++
	}

	if digitCount%2 == 1 {
		return true
	}

	denominator := int(math.Pow10(digitCount / 2))

	lastX := number % denominator

	firstX := number
	for firstX >= denominator {
		firstX /= 10
	}

	return lastX != firstX
}

func isNumberValid2(n int) bool {
	digits := 0
	for x := n; x > 0; x /= 10 {
		digits++
	}

	if digits == 1 {
		return true
	}

	// Try every possible chunk size
	// chunkSize must divide digits evenly AND must allow at least 2 repeats.
	for chunkSize := 1; chunkSize <= digits/2; chunkSize++ {
		if digits%chunkSize != 0 {
			continue
		}

		repeats := digits / chunkSize

		// Extract first chunk
		firstChunk := getLeadingDigits(n, chunkSize)

		// Reconstruct the number by repeating the chunk
		reconstructed := 0
		for i := 0; i < repeats; i++ {
			reconstructed = reconstructed*int(math.Pow10(chunkSize)) + firstChunk
		}

		if reconstructed == n {
			return false // invalid
		}
	}

	return true // valid
}

func getLeadingDigits(n, k int) int {
	// Count digits
	digits := 0
	for x := n; x > 0; x /= 10 {
		digits++
	}

	// Remove (digits - k) digits from right
	for i := 0; i < digits-k; i++ {
		n /= 10
	}

	return n
}
