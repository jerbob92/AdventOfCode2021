package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	// Illegal point lookup table for part 1.
	illegalPoints := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	// Keep track of the syntax error score for part 1.
	syntaxErrorScore := 0

	// Autocomplete point lookup table for part 2.
	autocompletePoints := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	// Store autocomplete scores for part 2.
	autocompleteScores := []int{}

	// First parse the input into a grid.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chunkDepth := -1
		chunkStarts := map[int]string{}
		corruptedLine := false
		for _, char := range line {
			if corruptedLine {
				break
			}
			stringChar := string(char)
			switch stringChar {
			case "(", "[", "{", "<":
				// Begin a new chunk.
				chunkDepth++
				chunkStarts[chunkDepth] = stringChar
			case ")", "]", "}", ">":
				expected := ""
				if chunkStarts[chunkDepth] == "(" {
					expected = ")"
				} else if chunkStarts[chunkDepth] == "[" {
					expected = "]"
				} else if chunkStarts[chunkDepth] == "{" {
					expected = "}"
				} else if chunkStarts[chunkDepth] == "<" {
					expected = ">"
				}
				// End chunk, test if this is the expected char.
				if stringChar != expected {
					syntaxErrorScore += illegalPoints[stringChar]
					corruptedLine = true
					continue
				}
				chunkDepth--
			}
		}

		// Part 2, if the line is not corrupted, it's incomplete.
		if !corruptedLine {
			autoCompleteScore := 0
			for currentChunk := chunkDepth; currentChunk >= 0; currentChunk-- {
				autoCompleteScore = (autoCompleteScore * 5) + autocompletePoints[chunkStarts[currentChunk]]
			}

			autocompleteScores = append(autocompleteScores, autoCompleteScore)
		}
	}

	// Sort the scores to get the middle one.
	sort.Ints(autocompleteScores)

	log.Printf("Part 1: total syntax error score: %d", syntaxErrorScore)
	log.Printf("Part 2: total syntax error score: %d", autocompleteScores[len(autocompleteScores)/2])
}
