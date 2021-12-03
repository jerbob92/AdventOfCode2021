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
	reportSize := 12
	reportSums := make([]int, reportSize)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bits := scanner.Text()
		for i, char := range bits {
			if string(char) == "1" {
				reportSums[i]++
			}
		}

		amountOfReports++
	}

	epsilonRate := int64(0)
	gammaRate := int64(0)
	for pos, reportSum := range reportSums {
		if reportSum > (amountOfReports / 2) {
			epsilonRate |= (1 << (reportSize - pos - 1))
		} else {
			gammaRate |= (1 << (reportSize - pos - 1))
		}
	}

	log.Printf("Epsilon rate: %b (%d), gamma rate: %b (%d), power consumtion: %d", epsilonRate, epsilonRate, gammaRate, gammaRate, epsilonRate*gammaRate)
}
