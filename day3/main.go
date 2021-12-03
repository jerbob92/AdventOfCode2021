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

	epsilonRate := ""
	gammaRate := ""

	for _, reportSum := range reportSums {
		if reportSum > (amountOfReports / 2) {
			epsilonRate += "1"
			gammaRate += "0"
		} else {
			epsilonRate += "0"
			gammaRate += "1"
		}
	}

	epsilonRateInt, _ := strconv.ParseInt(epsilonRate, 2, 64)
	gammaRateInt, _ := strconv.ParseInt(gammaRate, 2, 64)

	log.Printf("Epsilon rate: %s (%d), gamma rate: %s (%d), power consumtion: %d", epsilonRate, epsilonRateInt, gammaRate, gammaRateInt, epsilonRateInt * gammaRateInt)
}
