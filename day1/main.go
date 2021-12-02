package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
)

func main(){
	input, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Split on newlines
	depthStrings := bytes.Split(input, []byte("\n"))

	// Parse lines into numbers.
	var depths = make([]int, len(depthStrings))
	for i := range depthStrings {
		depth, err := strconv.Atoi(string(depthStrings[i]))
		if err != nil {
			continue
		}
		depths[i] = depth
	}

	// Keep track of amount of increments
	amountOfIncrements := 0
	amountOfSlidingWindowIncrements := 0
	slidingWindowSize := 3

	for i := range depths {
		// Compare to last depth.
		if i > 0 && depths[i] > depths[i-1] {
			amountOfIncrements++
		}

		// Compare to sliding window.
		if i >= slidingWindowSize {
			previousSlidingWindowSize := depths[i-1] + depths[i-2] + depths[i-3]
			currentSlidingWindowSize := depths[i] + depths[i-1] + depths[i-2]
			if currentSlidingWindowSize > previousSlidingWindowSize {
				amountOfSlidingWindowIncrements++
			}
		}
	}

	log.Printf("Part 1: Amount of measurements larger than the previous measurement: %d", amountOfIncrements)
	log.Printf("Part 2: Amount of sliding measurements window larger than the previous sliding measurements window: %d", amountOfSlidingWindowIncrements)
}