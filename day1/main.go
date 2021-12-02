package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Configure the window size to compare with.
	slidingWindowSize := 3

	// Keep track of amount of increments.
	amountOfIncrements := 0
	amountOfSlidingWindowIncrements := 0

	// Keep track of state.
	depthHistory := make([]int, slidingWindowSize)
	currentDepthIndex := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		if currentDepthIndex > 0 && depth > depthHistory[slidingWindowSize-1] {
			amountOfIncrements++
		}

		// Compare previous window to current window.
		if currentDepthIndex >= slidingWindowSize {
			previousSlidingWindowSize := sumIntSlice(depthHistory)
			currentSlidingWindowSize := (previousSlidingWindowSize - depthHistory[0]) + depth
			if currentSlidingWindowSize > previousSlidingWindowSize {
				amountOfSlidingWindowIncrements++
			}
		}

		// Rewrite sliding window history
		copy(depthHistory, depthHistory[1:])
		depthHistory[slidingWindowSize-1] = depth
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
