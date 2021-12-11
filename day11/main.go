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

// Flash flashes a single position. Returns the flashes.
func (g *Grid) Flash(x, y int) int {
	flashes := 0

	// I can only flash once per step.
	if !g.HasFlashed(x, y) {

		// Add some energy.
		g.Grid[y][x]++

		// Should I flash?
		if g.Grid[y][x] == 10 {

			// I'm flashing,
			// Drain my energy.
			g.Grid[y][x] = 0

			// Make sure I won't flash again.
			g.Flashed[y][x] = true

			// Flash my neighbours, register my flash and that of my neighbours.
			flashes+= 1 + g.FlashNeighbours(x, y)
		}
	}

	return flashes
}

// FlashNeighbours flashes the neighbours, pun intended.
func (g *Grid) FlashNeighbours(x, y int) int {
	// Collect all possible neighbours.
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
		if neighbour.X < 0 || neighbour.Y < 0 || neighbour.Y > len(g.Grid)-1 || neighbour.X > len(g.Grid[neighbour.Y])-1 {
			continue
		}

		// Flash my position and count the flashes.
		totalFlashes += g.Flash(neighbour.X, neighbour.Y)
	}

	return totalFlashes
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

// ResetFlashed Resets the map that keep tracks of flashes.
func (g *Grid) ResetFlashed() {
	g.Flashed = map[int]map[int]bool{}
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
				// Flash my position and count the flashes.
				stepFlashes += grid.Flash(xi, yi)
			}
		}

		// Reset the flashes after the step.
		grid.ResetFlashed()

		// Register the total amount of flashes.
		totalFlashes += stepFlashes

		// Part 1 answer, total amount of flashes after 100 steps.
		if step == 100 {
			log.Printf("Part 1: total amount of flashes: %d", totalFlashes)
		}

		// Part 2 answer, when all octopuses flash at the same time.
		if stepFlashes == len(grid.Grid) * len(grid.Grid) {
			log.Printf("Part 2: first step that all octopuses flashed: %d", step)

			// Quit when we have the answer for part 2.
			break
		}
	}
}
