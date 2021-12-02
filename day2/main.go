package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	HorizontalPosition int
	Depth              int

	// Only used in part 2
	Aim int
}

func main() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	subPart1 := &Submarine{}
	subPart2 := &Submarine{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commandParts := strings.Split(scanner.Text(), " ")
		if len(commandParts) < 2 {
			continue
		}

		// Parse the command in usable information.
		command := commandParts[0]
		amountString := commandParts[1]

		amount, err := strconv.Atoi(amountString)
		if err != nil {
			continue
		}

		switch command {
		case "forward":
			subPart1.HorizontalPosition += amount
			subPart2.HorizontalPosition += amount

			// In part 2, the aim * the forward amount determines the increase in depth.
			subPart2.Depth += subPart2.Aim * amount
		case "down":
			subPart1.Depth += amount
			subPart2.Aim += amount
		case "up":
			subPart1.Depth -= amount
			subPart2.Aim -= amount
		}
	}

	log.Printf("Part 1: horizontal position: %d, depth: %d, final position: %d", subPart1.HorizontalPosition, subPart1.Depth, subPart1.HorizontalPosition*subPart1.Depth)
	log.Printf("Part 2: horizontal position: %d, depth: %d, final position: %d", subPart2.HorizontalPosition, subPart2.Depth, subPart2.HorizontalPosition*subPart2.Depth)
}
