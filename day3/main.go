package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	amountOfReports := 0
	var reportSums []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bits := scanner.Text()

		// Initialize the array to keep track of the report sums.
		// We assume every line has the same length.
		if reportSums == nil {
			reportSums = make([]int, len(bits))
		}

		// Loop through every report.
		for i, char := range bits {
			// When a bit is positive, add one to the current position in the
			// report.
			if string(char) == "1" {
				reportSums[i]++
			}
		}

		amountOfReports++
	}

	epsilonRate := int64(0)
	gammaRate := int64(0)

	// Loop through every sum of the values in the report.
	for pos, reportSum := range reportSums {
		if reportSum > (amountOfReports / 2) {
			// If most reports on this position are positive, add one to epsilonRate.
			epsilonRate |= 1 << (len(reportSums) - pos - 1)
		} else {
			// If most reports on this position are negative, add one to gammaRate.
			gammaRate |= 1 << (len(reportSums) - pos - 1)
		}
	}

	log.Printf("Part 1: Epsilon rate: %b (%d), gamma rate: %b (%d), power consumtion: %d", epsilonRate, epsilonRate, gammaRate, gammaRate, epsilonRate*gammaRate)
}
