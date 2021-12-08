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

	// First parse the input into signal patterns and output values.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " | ")
		signalPatterns := strings.Split(lineParts[0], " ")
		outputValues := strings.Split(lineParts[1], " ")
		displayValues = append(displayValues, &DisplayValue{
			SignalPatterns: signalPatterns,
			OutputValues:   outputValues,
			Values:         append(signalPatterns, outputValues...),
			KnownSignals:   map[int]string{},
		})
	}

	// Part 1:
	uniqueValueCount := 0
	for displayI := range displayValues {
		for outputValueI := range displayValues[displayI].OutputValues {
			output := displayValues[displayI].OutputValues[outputValueI]
			outputLen := len(output)
			switch outputLen {
			case 2, 3, 4, 7:
				uniqueValueCount++
			default:
				// Unknown value.
			}
		}
	}

	log.Printf("Part 1: digits 1, 4, 7, or 8 appear %d times", uniqueValueCount)

	// Prepare part 2.
	for displayI := range displayValues {
		for valueI := range displayValues[displayI].Values {
			output := displayValues[displayI].Values[valueI]
			outputLen := len(output)
			switch outputLen {
			case 2:
				displayValues[displayI].KnownSignals[1] = output
			case 3:
				displayValues[displayI].KnownSignals[7] = output
			case 4:
				displayValues[displayI].KnownSignals[4] = output
			case 7:
				displayValues[displayI].KnownSignals[8] = output
			}
		}
	}

	// Part 2:
	totalOfAllNumbers := 0
	for displayI := range displayValues {
		displayOutput := ""
		for outputValueI := range displayValues[displayI].OutputValues {
			output := displayValues[displayI].OutputValues[outputValueI]
			outputLen := len(output)
			switch outputLen {
			case 2:
				displayOutput += "1"
			case 3:
				displayOutput += "7"
			case 4:
				displayOutput += "4"
			case 5, 6:
				signalOverlaps := map[int]int{}
				// If we have 5 or 6 signals, we have to try to figure out which number it is.
				// The numbers left are 0, 2, 3, 5, 6, 9
				for knownSignalNumber, knownSignal := range displayValues[displayI].KnownSignals {
					if _, ok := signalOverlaps[knownSignalNumber]; !ok {
						signalOverlap := amountOfSignalOverlap(output, knownSignal)
						signalOverlaps[knownSignalNumber] = signalOverlap
					}
				}

				candidatePoints := map[int]int{}
				for signalOverlapsNumber, signalOverlapsOverlap := range signalOverlaps {
					if table, ok := OverlapTable[signalOverlapsNumber]; ok {
						if candidates, ok := table[signalOverlapsOverlap]; ok {
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


				// Find the number with the highest overlap.
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
