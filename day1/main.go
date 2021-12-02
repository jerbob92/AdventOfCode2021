package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main(){
	slidingWindowSize := 3

	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	currentDepthIndex := 0

	// Keep track of amount of increments.
	amountOfIncrements := 0
	amountOfSlidingWindowIncrements := 0

	// Keep track of previous values.
	var previousDepth *int
	slidingWindowHistory := make([]int, slidingWindowSize)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		if previousDepth != nil && depth > *previousDepth {
			amountOfIncrements++
		}

		previousDepth = &depth

		// Compare previous window to current window.
		if currentDepthIndex >= slidingWindowSize {
			previousSlidingWindowSize := sumIntSlice(slidingWindowHistory)
			currentSlidingWindowSize := (previousSlidingWindowSize - slidingWindowHistory[0]) + depth
			if currentSlidingWindowSize > previousSlidingWindowSize {
				amountOfSlidingWindowIncrements++
			}
		}

		// Rewrite sliding window history
		slidingWindowHistory = append(slidingWindowHistory[1:], depth)
		currentDepthIndex++
	}

	log.Printf("Part 1: Amount of measurements larger than the previous measurement: %d", amountOfIncrements)
	log.Printf("Part 2: Amount of sliding measurements window larger than the previous sliding measurements window: %d", amountOfSlidingWindowIncrements)
}

func sumIntSlice(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}