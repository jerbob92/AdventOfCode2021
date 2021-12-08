package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type DisplayValue struct {
	OutputValues []string
	KnownSignals map[int]string
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
			OutputValues: outputValues,
			KnownSignals: map[int]string{},
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
		for _, output := range append(signalPatterns, outputValues...) {
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
					signalOverlap := amountOfSignalOverlap(output, knownSignal)
					signalOverlaps[knownSignalNumber] = signalOverlap
				}

				// Keep track of the amount of points for overlap with each known number.
				// Since each combination of overlap in unknown and known numbers
				// results in a unique number, we can then figure out which number it is.
				candidatePoints := map[int]int{}

				// Loop through the overlap stats to get the candidates per overlap count.
				for signalOverlapsNumber, signalOverlapsOverlap := range signalOverlaps {
					// Lookup the current known number in the overlap table.
					// Get the candidates from that map that have this amount of overlap.
					// For each candidate in that map value, add one point.
					for candidateI := range OverlapTable[signalOverlapsNumber][signalOverlapsOverlap] {
						candidatePoints[OverlapTable[signalOverlapsNumber][signalOverlapsOverlap][candidateI]]++
					}
				}

				// Find the candidate number with the highest overlap.
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

// amountOfSignalOverlap returns the amount of chars 2 strings overlap.
func amountOfSignalOverlap(a, b string) int {
	// Create an index of known chars in string a.
	aIndex := map[string]bool{}
	for i := range a {
		aIndex[string(a[i])] = true
	}

	// Loop through string b to see if it is also in a using the index.
	overlap := 0
	for i := range b {
		if _, ok := aIndex[string(b[i])]; ok {
			overlap++
		}
	}

	return overlap
}
