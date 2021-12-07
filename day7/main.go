package main

import (
	"bufio"
	"bytes"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	crabPositions := map[int]int{}

	// First parse the input into crab positions.
	scanner := bufio.NewScanner(file)
	scanner.Split(scanCommas)
	maxPosition := int(0)
	for scanner.Scan() {
		number := scanner.Text()
		position, _ := strconv.Atoi(number)

		// Register the crab positions and count.
		crabPositions[position]++

		// Register the max crab position.
		if position > maxPosition {
			maxPosition = position
		}
	}

	calculateFuelCost := func(exponentialFuelCost bool) (int, int) {
		// Keep track of the cheapest target position.
		cheapestTargetPosition := 0
		cheapestTargetPositionFuelUsed := 0

		// Loop through all possible target positions.
		for position := 0; position <= maxPosition; position++ {
			// Keep track of all fuel used for this target position.
			fuelUsedForPosition := 0

			// Loop through all current crab positions.
			for currentPosition, crabCount := range crabPositions {
				// Calculate te distance, sadly math.Abs only supports floats.
				distance := int(math.Abs(float64(position) - float64(currentPosition)))
				if exponentialFuelCost {
					fuelUsedForPosition += (distance * (distance + 1) / 2) * crabCount
				} else {
					fuelUsedForPosition += distance * crabCount
				}
			}

			// Check if the current target position is the cheapest.
			if cheapestTargetPositionFuelUsed == 0 || fuelUsedForPosition < cheapestTargetPositionFuelUsed {
				cheapestTargetPositionFuelUsed = fuelUsedForPosition
				cheapestTargetPosition = position
			}
		}

		return cheapestTargetPosition, cheapestTargetPositionFuelUsed
	}

	part1CheapestTargetPosition, part1CheapestTargetPositionFuelUsed := calculateFuelCost(false)
	log.Printf("Part 1: cheapest target position is %d with fuel usage: %d", part1CheapestTargetPosition, part1CheapestTargetPositionFuelUsed)

	part2CheapestTargetPosition, part2CheapestTargetPositionFuelUsed := calculateFuelCost(true)
	log.Printf("Part 2: cheapest target position is %d with exponential fuel usage: %d", part2CheapestTargetPosition, part2CheapestTargetPositionFuelUsed)
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
