package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	X int
	Y int
}

type BingoBoard struct {
	Rows        [][]int
	Columns     [][]int
	RowSums     []int
	ColSums     []int
	Sum         int
	NumberIndex map[int][]Pos
	GridSize    int
	Drawn       map[int]bool
	Won         bool
}

// String prints the current state of the board
func (b *BingoBoard) String() string {
	printedRows := []string{}
	colSums := make([]int, b.GridSize)
	for i, row := range b.Rows {
		vals := []string{}
		total := 0
		for colI, val := range row {
			valToPrint := strconv.Itoa(val)
			if b.Drawn[val] {
				valToPrint = "_" + valToPrint + "_"
			}
			vals = append(vals, valToPrint)
			total+= val
			colSums[colI] += val
		}

		vals = append(vals, "Begin sum:" + strconv.Itoa(total))
		vals = append(vals, "Current sum:" + strconv.Itoa(b.RowSums[i]))

		printedRows = append(printedRows, strings.Join(vals, "\t"))
	}

	printedRows = append(printedRows, "")

	colSumsFormatted := []string{}
	for _, val := range colSums {
		colSumsFormatted = append(colSumsFormatted, strconv.Itoa(val))
	}
	printedRows = append(printedRows, strings.Join(colSumsFormatted, "\t"))

	colSumsFormatted = []string{}
	for _, val := range b.ColSums {
		colSumsFormatted = append(colSumsFormatted, strconv.Itoa(val))
	}
	printedRows = append(printedRows, strings.Join(colSumsFormatted, "\t"))

	return "\n" + strings.Join(printedRows, "\n")
}

// Draw draws the number for this board and changes the internal state.
func (b *BingoBoard) Draw(number int) {
	if _, ok := b.NumberIndex[number]; ok {
		b.Drawn[number] = true
		for _, pos := range b.NumberIndex[number] {
			b.RowSums[pos.Y] -= number
			b.ColSums[pos.X] -= number
			b.Sum -= number

			// Check if any rows or columns have a sum of 0.
			if b.RowSums[pos.Y] == 0 || b.ColSums[pos.X] == 0 {
				seenAllNumbers := false

				// Validate we actually had all numbers, since 0 is a thing.
				// Validate if row is indeed completely drawn.
				if b.RowSums[pos.Y] == 0 {
					seenAllNumbers = true
					for _, val := range b.Rows[pos.Y] {
						if _, ok := b.Drawn[val]; !ok {
							seenAllNumbers = false
						}
					}
				}

				// Check if column is indeed completely drawn.
				if b.ColSums[pos.X] == 0 {
					seenAllNumbers = true
					for _, val := range b.Columns[pos.X] {
						if _, ok := b.Drawn[val]; !ok {
							seenAllNumbers = false
						}
					}
				}

				if seenAllNumbers {
					b.Won = true
				}
			}
		}
	}
}

func main() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	boards := []*BingoBoard{}
	var drawnNumbers []int

	currentBoard := -1
	currentBoardRow := 0

	// First parse the input into draws and boards.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// The first line is the drawing of the numbers.
		if line != "" && currentBoard == -1 {
			drawnNumbersLine := strings.Split(line, ",")
			drawnNumbers = make([]int, len(drawnNumbersLine))
			for i, drawnNumber := range drawnNumbersLine {
				parsedNumber, _ := strconv.Atoi(drawnNumber)
				drawnNumbers[i] = parsedNumber
			}
			continue
		}

		// Reset for new baord.
		if line == "" {
			currentBoard++
			currentBoardRow = 0
			boards = append(boards, &BingoBoard{
				Rows:        [][]int{},
				Columns:     [][]int{},
				RowSums:     []int{},
				ColSums:     []int{},
				Drawn:       map[int]bool{},
				NumberIndex: map[int][]Pos{},
			})
			continue
		}

		// Add a new row.
		boards[currentBoard].Rows = append(boards[currentBoard].Rows, []int{})
		boards[currentBoard].RowSums = append(boards[currentBoard].RowSums, 0)

		// Scan the line into "words", every word is a number.
		rowScanner := bufio.NewScanner(bytes.NewBufferString(line))
		rowScanner.Split(bufio.ScanWords)
		col := 0
		for rowScanner.Scan() {
			if currentBoardRow == 0 {
				// Calculate the grid size.
				boards[currentBoard].GridSize++

				// Add a new column.
				boards[currentBoard].Columns = append(boards[currentBoard].Columns, []int{})
				boards[currentBoard].ColSums = append(boards[currentBoard].ColSums, 0)
			}

			parsedNumber, _ := strconv.Atoi(rowScanner.Text())

			// We keep track of the sum of the board, we need it later.
			boards[currentBoard].Sum += parsedNumber

			// Set the parsedNumber in the row and col.
			boards[currentBoard].Rows[currentBoardRow] = append(boards[currentBoard].Rows[currentBoardRow], parsedNumber)
			boards[currentBoard].RowSums[currentBoardRow] += parsedNumber
			boards[currentBoard].Columns[col] = append(boards[currentBoard].Columns[col], parsedNumber)
			boards[currentBoard].ColSums[col] += parsedNumber

			// Keep an index of all numbers on this board.
			if _, ok := boards[currentBoard].NumberIndex[parsedNumber]; !ok {
				boards[currentBoard].NumberIndex[parsedNumber] = []Pos{}
			}
			boards[currentBoard].NumberIndex[parsedNumber] = append(boards[currentBoard].NumberIndex[parsedNumber], Pos{
				X: col,
				Y: currentBoardRow,
			})

			col++
		}

		currentBoardRow++
	}

	gotWinners := 0

	// Start drawing numbers.
	for _, drawnNumber := range drawnNumbers {
		for i, board := range boards {
			// Skip when already won.
			if board.Won {
				continue
			}

			// Draw the number in the board.
			board.Draw(drawnNumber)

			// Check whether the board one.
			if board.Won {
				gotWinners++

				// Only first winner matters for part 1.
				if gotWinners == 1 {
					log.Printf("Part 1: Board %d won first when drawing number %d, score: %d", i+1, drawnNumber, board.Sum*drawnNumber)
				}

				// We know the last winner when the amount of winners equals the amount of boards.
				if gotWinners == len(boards) {
					log.Printf("Part 2: Board %d won last when drawing number %d, score: %d", i+1, drawnNumber, board.Sum*drawnNumber)
				}
			}
		}
	}
}
