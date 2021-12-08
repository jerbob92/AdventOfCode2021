package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type DisplayValue struct {
	SignalPatterns []string
	OutputValues   []string
	Values         []string
	KnownSignals   map[int]string
}

// OverlapTable maps how much overlap numbers have with other numbers.
// The first index is the number for which we know the signal.
// The second index is the amount of overlaps.
// The value is the numbers that have that amount of overlap.
// The table is structured like this for the quickest lookups.
var OverlapTable = map[int]map[int][]int{
	1: {
		1: []int{2, 5, 6},
		2: []int{0, 3, 9},
	},
	4: {
		2: []int{2},
		3: []int{0, 3, 5, 6},
		4: []int{9},
	},
	7: {
		2: []int{2, 5, 6},
		3: []int{0, 3, 9},
	},
	8: {
		5: []int{2, 3, 5},
		6: []int{0, 6, 9},
	},
}

func main() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	displayValues := []*DisplayValue{}

	// Keep track of unique values for part 1.
	uniqueValueCount := 0

	// First parse the input into signal patterns and output values.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " | ")
		signalPatterns := strings.Split(lineParts[0], " ")
		outputValues := strings.Split(lineParts[1], " ")
		displayValue := &DisplayValue{
			SignalPatterns: signalPatterns,
			OutputValues:   outputValues,
			Values:         append(signalPatterns, outputValues...),
			KnownSignals:   map[int]string{},
		}

		// Calculate part 1:
		for outputValueI := range displayValue.OutputValues {
			switch len(displayValue.OutputValues[outputValueI]) {
			case 2, 3, 4, 7:
				uniqueValueCount++
			}
		}

		// Prepare part 2.
		// Save the known signals for unique numbers.
		for valueI := range displayValue.Values {
			output := displayValue.Values[valueI]
			switch len(output) {
			case 2:
				displayValue.KnownSignals[1] = output
			case 3:
				displayValue.KnownSignals[7] = output
			case 4:
				displayValue.KnownSignals[4] = output
			case 7:
				displayValue.KnownSignals[8] = output
			}
		}

		displayValues = append(displayValues, displayValue)
	}

	log.Printf("Part 1: digits 1, 4, 7, or 8 appear %d times", uniqueValueCount)

	// Part 2:
	totalOfAllNumbers := 0
	for displayI := range displayValues {
		displayOutput := ""

		// Loop through the output values, convert them to number one by one.
		for outputValueI := range displayValues[displayI].OutputValues {
			output := displayValues[displayI].OutputValues[outputValueI]

			// Check the output length, some numbers are already unique.
			switch len(output) {
			case 2:
				displayOutput += "1"
			case 3:
				displayOutput += "7"
			case 4:
				displayOutput += "4"
			case 5, 6:
				// If we have 5 or 6 signals, we have to try to figure out which number it is.
				// The numbers left are 0, 2, 3, 5, 6, 9

				// Keep track of the amount of overlaps with known numbers.
				signalOverlaps := map[int]int{}

				// Figure out the amount of overlaps with this display value and the known numbers.
				// We need this for the lookup table later on.
				for knownSignalNumber, knownSignal := range displayValues[displayI].KnownSignals {
					if _, ok := signalOverlaps[knownSignalNumber]; !ok {
						signalOverlap := amountOfSignalOverlap(output, knownSignal)
						signalOverlaps[knownSignalNumber] = signalOverlap
					}
				}

				// Keep track of the amount of points for overlap with each known number.
				// Since each combination of overlap in unknown and known numbers
				// results in a unique number, we can then figure out which number it is.
				candidatePoints := map[int]int{}

				// Loop through the overlap stats.
				for signalOverlapsNumber, signalOverlapsOverlap := range signalOverlaps {

					// Lookup the current known number in the overlap table.
					if table, ok := OverlapTable[signalOverlapsNumber]; ok {

						// Get the candidates from the map that have this amount of overlap.
						if candidates, ok := table[signalOverlapsOverlap]; ok {

							// For each candidate, add one point.
							for candidateI := range candidates {
								candidatePoints[candidates[candidateI]]++
							}
						} else {
							log.Printf("Could not find candidates for number %d and overlap amount %d", signalOverlapsNumber, signalOverlapsOverlap)
						}
					} else {
						log.Printf("Could not find overlap table for number %d", signalOverlapsNumber)
					}
				}

				// Find the canidate number with the highest overlap.
				highestCandidate := -1
				highestCandidateOverlap := 0
				for number, overlap := range candidatePoints {
					if highestCandidate == -1 || overlap > highestCandidateOverlap {
						highestCandidate = number
						highestCandidateOverlap = overlap
					}
				}

				displayOutput += strconv.Itoa(highestCandidate)
			case 7:
				displayOutput += "8"
			}
		}

		parsedNumber, _ := strconv.Atoi(displayOutput)
		totalOfAllNumbers += parsedNumber
	}

	log.Printf("Part 2: sum of all output numbers: %d", totalOfAllNumbers)
}

func amountOfSignalOverlap(a, b string) int {
	overlap := 0
	aIndex := map[string]bool{}
	for i := range a {
		aIndex[string(a[i])] = true
	}

	for i := range b {
		if _, ok := aIndex[string(b[i])]; ok {
			overlap++
		}
	}

	return overlap
}
