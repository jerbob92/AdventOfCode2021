package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

	combinedCandidates := map[bool][]string{}

	// Reset scanner to scan again.
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	// We already know the sums of the first bit, because that contains all
	// rows that we have seen, use that to build a first list of candidates.
	for scanner.Scan() {
		bits := scanner.Text()

		oneValues := reportSums[0]
		zeroValues := amountOfReports - oneValues
		mostCommonValue := "0"
		leastCommonValue := "1"

		if oneValues >= zeroValues {
			mostCommonValue = "1"
		}

		if zeroValues <= oneValues {
			leastCommonValue = "0"
		}

		if string(bits[0]) == mostCommonValue {
			combinedCandidates[true] = append(combinedCandidates[true], bits)
		}

		if string(bits[0]) == leastCommonValue {
			combinedCandidates[false] = append(combinedCandidates[false], bits)
		}
	}

	// Loop through the candidates for the different values.
	for mostCommon, candidates := range combinedCandidates {
		// Loop through the candidates until we have eliminated all but one.
		// We already did position one, so we can skip that.
		bitPosition := 1
		for bitPosition < len(reportSums) {
			// First find the most common value for the current candidates.
			oneValues := 0
			for _, candidate := range candidates {
				if string(candidate[bitPosition]) == "1" {
					oneValues++
				}
			}

			zeroValues := len(candidates) - oneValues
			mostCommonValue := "0"
			leastCommonValue := "1"

			if oneValues >= zeroValues {
				mostCommonValue = "1"
			}

			if zeroValues <= oneValues {
				leastCommonValue = "0"
			}

			// Filter out the candidates that do not match.
			newCandidates := []string{}
			for _, candidate := range candidates {
				if mostCommon && string(candidate[bitPosition]) == mostCommonValue {
					newCandidates = append(newCandidates, candidate)
				}

				if !mostCommon && string(candidate[bitPosition]) == leastCommonValue {
					newCandidates = append(newCandidates, candidate)
				}
			}
			candidates = newCandidates

			// When there is one value left, stop
			if len(candidates) == 1 {
				break
			}

			// Move to the next bit.
			bitPosition++
		}

		// Overwrite the candidates with the new list.
		combinedCandidates[mostCommon] = candidates
	}

	oxygenGeneratorRating, _ := strconv.ParseInt(combinedCandidates[true][0], 2, 64)
	CO2ScrubberRating, _ := strconv.ParseInt(combinedCandidates[false][0], 2, 64)

	log.Printf("Part 2: oxygen generator rating: %s (%d), CO2 scrubber rating: %s (%d), life support rating: %d", combinedCandidates[true][0], oxygenGeneratorRating, combinedCandidates[false][0], CO2ScrubberRating, oxygenGeneratorRating*CO2ScrubberRating)
}
