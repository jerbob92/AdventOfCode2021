package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Make an array 0-8, 8 is the max lifetime of a fish.
	fishTimers := make([]int64, 9)

	// First parse the input into draws and boards.
	scanner := bufio.NewScanner(file)
	scanner.Split(scanCommas)
	totalFish := int64(0)
	for scanner.Scan() {
		number := scanner.Text()
		timer, _ := strconv.Atoi(number)

		// Register the initial fish timers.
		fishTimers[timer]++

		// Register the initial fish count.
		totalFish++
	}

	// Method to progress x number of days, to do 2 parts in one.
	progressDays := func(days int) {

		// Progress each day.
		for day := 1; day <= days; day++ {

			// Keep track of how many fish need to be reset/reproduced after this run.
			resetFish := int64(0)
			for timer, fishCount := range fishTimers {
				// Timer 0 is special, fish needs a reset and reproduce.
				if timer == 0 {
					resetFish = fishCount
				} else {
					// Move every timer one down in the timer array, because one
					// day has progressed.
					fishTimers[timer - 1] = fishCount
				}
			}

			// Reset the timers for reset fish to 6.
			fishTimers[6] += resetFish

			// Add the resetFish as new fish with timer 8.
			fishTimers[8] = resetFish

			// Keep track of the total amount of fish we currently have.
			totalFish += resetFish
		}
	}

	part1Days := 80
	progressDays(part1Days)
	log.Printf("Part 1: Amount of fish after %d days: %d", part1Days, totalFish)

	part2Days := 256
	progressDays(part2Days-part1Days)
	log.Printf("Part 2: Amount of fish after %d days: %d", part2Days, totalFish)
}

// scanCommas is a helper method that can be used in a scanner to read input by comma.
func scanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}