package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type Grid struct {
	Grid [][]int
}

func main() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := &Grid{
		Grid: [][]int{},
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

	// Keep track of the lowest locations for part 2.
	lowestLocations := [][]int{}

	// Loop through the rows and cols to calculate the risk.
	totalRisk := 0
	for rowI := 0; rowI < len(grid.Grid); rowI++ {
		for colI := 0; colI < len(grid.Grid[rowI]); colI++ {
			val := grid.Grid[rowI][colI]
			lowestLocation := true
			if colI > 0 {
				if grid.Grid[rowI][colI-1] <= val {
					lowestLocation = false
				}
			}
			if colI < len(grid.Grid[rowI])-1 {
				if grid.Grid[rowI][colI+1] <= val {
					lowestLocation = false
				}
			}
			if rowI > 0 {
				if grid.Grid[rowI-1][colI] <= val {
					lowestLocation = false
				}
			}
			if rowI < len(grid.Grid)-1 {
				if grid.Grid[rowI+1][colI] <= val {
					lowestLocation = false
				}
			}

			if lowestLocation {
				totalRisk += 1 + val
				lowestLocations = append(lowestLocations, []int{rowI, colI})
			}
		}
	}

	log.Printf("Part 1: the sum of the risk levels of all low points: %d", totalRisk)

	basinSizes := []int{}
	for _, lowestLocation := range lowestLocations {
		// Find basin size, our own number also counts towards the basin.
		basinSize := 1 + findAndMarkNeighbours(lowestLocation[0], lowestLocation[1], grid)
		basinSizes = append(basinSizes, basinSize)
	}

	// Sort the basin sizes.
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	log.Printf("Part 2: the multiplication of the basin size of the 3 largest basins: %d * %d * %d = %d", basinSizes[0], basinSizes[1], basinSizes[2], basinSizes[0]*basinSizes[1]*basinSizes[2])
}

// findAndMarkNeighbours is a recursive method that find neighbour values on
// a grid, count them and then find their neighbours until no higher values
// can be found or the higher value is a 9.
func findAndMarkNeighbours(x, y int, grid *Grid) int {
	// Save the original value of the position.
	itemValue := grid.Grid[x][y]

	// Reset my grid value so neighbours don't count me towards the basin size.
	grid.Grid[x][y] = 9

	// Find higher neighbours, count them towards the basin size and find
	// neighbours in the neighbours.
	basinSize := 0
	if y > 0 {
		if grid.Grid[x][y-1] > itemValue && grid.Grid[x][y-1] < 9 {
			basinSize += 1 + findAndMarkNeighbours(x, y-1, grid)
		}
	}
	if y < len(grid.Grid[x])-1 {
		if grid.Grid[x][y+1] > itemValue && grid.Grid[x][y+1] < 9 {
			basinSize += 1 + findAndMarkNeighbours(x, y+1, grid)
		}
	}
	if x > 0 {
		if grid.Grid[x-1][y] > itemValue && grid.Grid[x-1][y] < 9 {
			basinSize += 1 + findAndMarkNeighbours(x-1, y, grid)
		}
	}
	if x < len(grid.Grid)-1 {
		if grid.Grid[x+1][y] > itemValue && grid.Grid[x+1][y] < 9 {
			basinSize += 1 + findAndMarkNeighbours(x+1, y, grid)
		}
	}

	return basinSize
}
