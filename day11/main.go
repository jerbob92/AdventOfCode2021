package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	Grid    [][]int
	Flashed map[int]map[int]bool
}

type Pos struct {
	X int
	Y int
}

// Print prints the current grid.
func (g *Grid) Print() {
	for yi := 0; yi < len(g.Grid); yi++ {
		row := []string{}
		for xi := 0; xi < len(g.Grid); xi++ {
			row = append(row, strconv.Itoa(g.Grid[yi][xi]))
		}
		log.Println(strings.Join(row, ""))
	}
}

// HasFlashed returns whether a location has flashed.
func (g *Grid) HasFlashed(x, y int) bool {
	if _, ok := g.Flashed[y]; ok {
		if val, ok := g.Flashed[y][x]; ok {
			if val {
				return true
			}
		}
	} else {
		g.Flashed[y] = map[int]bool{}
	}

	return false
}

func main() {
	file, err := os.Open("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := &Grid{
		Grid:    [][]int{},
		Flashed: map[int]map[int]bool{},
	}

	// First parse the input into a grid.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, char := range line {
			number, _ := strconv.Atoi(string(char))
			row[i] = number
		}

		grid.Grid = append(grid.Grid, row)
	}

	// Keep track of the total amount of flashes.
	totalFlashes := 0

	// Keep track of the current step.
	step := 0
	for true {
		step++

		// Keep track of the amount of flashes in this step for part 2.
		stepFlashes := 0

		// Loop through the grid.
		for yi := 0; yi < len(grid.Grid); yi++ {
			for xi := 0; xi < len(grid.Grid); xi++ {

				// I can only flash once per step.
				if !grid.HasFlashed(xi, yi) {

					// Add some energy.
					grid.Grid[yi][xi]++

					// Should I flash?
					if grid.Grid[yi][xi] == 10 {

						// I'm flashing,
						// Drain my energy.
						grid.Grid[yi][xi] = 0

						// Make sure I won't flash again.
						grid.Flashed[yi][xi] = true

						// Flash my neighbours, register my flash and that of my neighbours.
						stepFlashes+= 1 + flashNeighbours(xi, yi, grid)
					}
				}
			}
		}

		// Reset the flashes after the step.
		grid.Flashed = map[int]map[int]bool{}

		// Register the total amount of flashes.
		totalFlashes += stepFlashes

		// Part 1 answer.
		if step == 100 {
			log.Printf("Part 1: total amount of flashes: %d", totalFlashes)
		}

		// Part 2 answer and quit.
		if stepFlashes == len(grid.Grid) * len(grid.Grid) {
			log.Printf("Part 2: first step that all octopuses flashed: %d", step)
			break
		}
	}

}

// flashNeighbours flashes the neighbours, pun intended.
func flashNeighbours(x, y int, grid *Grid) int {
	neighbours := []Pos{
		// Same row
		{
			X: x - 1,
			Y: y,
		},
		{
			X: x + 1,
			Y: y,
		},
		// Row above
		{
			X: x - 1,
			Y: y - 1,
		},
		{
			X: x,
			Y: y - 1,
		},
		{
			X: x + 1,
			Y: y - 1,
		},
		// Row below
		{
			X: x - 1,
			Y: y + 1,
		},
		{
			X: x,
			Y: y + 1,
		},
		{
			X: x + 1,
			Y: y + 1,
		},
	}

	totalFlashes := 0

	// Flash every neighbour.
	for _, neighbour := range neighbours {
		// Skip invalid neighbours.
		if neighbour.X < 0 || neighbour.Y < 0 || neighbour.Y > len(grid.Grid)-1 || neighbour.X > len(grid.Grid[neighbour.Y])-1 {
			continue
		}

		// I can only flash once per step.
		if !grid.HasFlashed(neighbour.X, neighbour.Y)  {

			// Add some energy.
			grid.Grid[neighbour.Y][neighbour.X]++

			// Should I flash?
			if grid.Grid[neighbour.Y][neighbour.X] == 10 {
				// I'm flashing,
				// Drain my energy.
				grid.Grid[neighbour.Y][neighbour.X] = 0

				// Make sure I won't flash again.
				grid.Flashed[neighbour.Y][neighbour.X] = true

				// Flash my neighbours, register my flash and that of my neighbours.
				totalFlashes += 1 + flashNeighbours(neighbour.X, neighbour.Y, grid)
			}
		}
	}

	return totalFlashes
}
