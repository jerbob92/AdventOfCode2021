package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type VentMap struct {
	Vents map[int]map[int]int

	// Only used for printing.
	MaxX int
	MaxY int
}

// ensureLocation will make sure the maps have a value for the location.
// It also keeps track the map size.
func (v *VentMap) ensureLocation(x, y int) {
	if v.Vents == nil {
		v.Vents = map[int]map[int]int{}
	}

	// Ensure x location.
	if _, ok := v.Vents[x]; !ok {
		v.Vents[x] = map[int]int{}
	}

	// Ensure y location.
	if _, ok := v.Vents[x][y]; !ok {
		v.Vents[x][y] = 0
	}

	// Keep track of max x and max for printing a pretty map.
	if x > v.MaxX {
		v.MaxX = x
	}
	if y > v.MaxY {
		v.MaxY = y
	}
}

// Mark will mark a vent on the map and ensure the location is on the map.
func (v *VentMap) Mark(x, y int) {
	// First ensure the location.
	v.ensureLocation(x, y)

	// Now that we have the location on the map, add a vent.
	v.Vents[x][y]++
}

// Print is a helper to draw the current state of the map.
func (v *VentMap) Print() {
	for yi := 0; yi <= v.MaxY; yi++ {
		row := []string{}
		for xi := 0; xi <= v.MaxX; xi++ {
			if _, ok := v.Vents[xi]; ok {
				if val, ok := v.Vents[xi][yi]; ok {
					row = append(row, strconv.Itoa(val))
				} else {
					row = append(row, ".")
				}
			} else {
				row = append(row, ".")
			}
		}
		log.Println(strings.Join(row, ""))
	}
}

// GetNumberOfOverlaps counts the positions of the map that more than 1
// occurrence.
func (v *VentMap) GetNumberOfOverlaps() int {
	counter := 0
	for _, xVal := range v.Vents {
		for _, yVal := range xVal {
			if yVal > 1 {
				counter++
			}
		}
	}

	return counter
}

func main() {
	part1Map := drawMap(false)
	//part1Map.Print()
	log.Printf("Part 1: Total number of overlaps: %d", part1Map.GetNumberOfOverlaps())

	part2Map := drawMap(true)
	//part2Map.Print()
	log.Printf("Part 2: Total number of overlaps: %d", part2Map.GetNumberOfOverlaps())
}

func drawMap(drawDiagonals bool) *VentMap {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ventMap := &VentMap{}

	// First parse the input into draws and boards.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Convert the instructions to x,y positions.
		lineParts := strings.Split(line, " -> ")
		xy1 := strings.Split(lineParts[0], ",")
		xy2 := strings.Split(lineParts[1], ",")
		x1, _ := strconv.Atoi(xy1[0])
		y1, _ := strconv.Atoi(xy1[1])
		x2, _ := strconv.Atoi(xy2[0])
		y2, _ := strconv.Atoi(xy2[1])

		// Part one only considers vents that have matching x or matching y.
		if x1 == x2 || y1 == y2 {
			// Switch values around if necessary, makes looping easier.
			startX := x1
			endX := x2
			if endX < startX {
				startX = x2
				endX = x1
			}

			startY := y1
			endY := y2
			if endY < startY {
				startY = y2
				endY = y1
			}

			// Loop through the x and y values until we reach the end position.
			for xi := startX; xi <= endX; xi++ {
				for yi := startY; yi <= endY; yi++ {
					// Mark the vent on this location.
					ventMap.Mark(xi, yi)
				}
			}
		} else if drawDiagonals {
			// For part 2 we also need to draw diagonals.
			// Starting positions
			startX := x1
			endX := x2
			startY := y1
			endY := y2

			// Only switch around the first loop this time.
			// We don't want the line to go the wrong direction.
			if endX < startX {
				startX = x2
				endX = x1
				startY = y2
				endY = y1
			}

			// Figure out if we need to draw 45 or -45 degrees.
			goBackWards := false
			if endY < startY {
				goBackWards = true
			}

			// Draw diagonal line
			for xi := startX; xi <= endX; xi++ {
				// Start a different loop based on the direction.
				if !goBackWards {
					for yi := startY; yi <= endY; yi++ {
						// Mark the vent on this location.
						ventMap.Mark(xi, yi)

						// Automatically go to the next row since we go
						// diagonally, so we're not marking the full row.
						xi++
					}
				} else {
					for yi := startY; yi >= endY; yi-- {
						// Mark the vent on this location.
						ventMap.Mark(xi, yi)

						// Automatically go to the next row since we go
						// diagonally, so we're not marking the full row.
						xi++
					}
				}
			}
		}
	}

	return ventMap
}
