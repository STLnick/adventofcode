package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const GRID_ON string = "#"
const GRID_OFF string = "."

type Grid struct {
	board [][]string
	cols  int
	rows  int
}

func (g Grid) String() string {
	var s string
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.cols; col++ {
			s += g.board[row][col]
		}
		s += "\n"
	}
	return s
}

func (g Grid) IsCorner(row int, col int) bool {
	if row == 0 {
		if col == 0 || col == g.cols-1 {
			return true
		}
	} else if row == g.rows-1 {
		if col == 0 || col == g.cols-1 {
			return true
		}
	}
	return false
}

func (g Grid) ProcessTick(isPartTwo bool) {
	var str string
	var onNeighbors int
	snapshot := cloneGrid(g)

	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.cols; col++ {
			str = g.board[row][col]
			onNeighbors = getOnNeighborCount(snapshot, row, col)

			if str == GRID_ON {
				if onNeighbors < 2 || onNeighbors > 3 {
					g.board[row][col] = GRID_OFF
					if isPartTwo && g.IsCorner(row, col) {
						g.board[row][col] = GRID_ON
					}
				}
			} else {
				if onNeighbors == 3 {
					g.board[row][col] = GRID_ON
				}
			}
		}
	}
}

func (g Grid) LightsOn() int {
	var count int
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.cols; col++ {
			if g.board[row][col] == GRID_ON {
				count++
			}
		}
	}
	return count
}

func getOnNeighborCount(grid Grid, row int, col int) int {
	neighbors := [][]int{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}
	var count int

	for _, n := range neighbors {
		if n[0] < 0 || n[0] == grid.rows || n[1] < 0 || n[1] == grid.cols {
			continue
		}
		if grid.board[n[0]][n[1]] == GRID_ON {
			count++
		}
	}
	return count
}

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func buildGrid(lines []string, isTest bool) Grid {
	var cols int
	var rows int

	if isTest {
		cols = 6
		rows = 6
	} else {
		cols = 100
		rows = 100
	}

	grid := Grid{
		board: make([][]string, rows),
		cols:  cols,
		rows:  rows,
	}

	for row, line := range lines {
		if line == "" {
			continue
		}
		newRow := make([]string, cols)
		grid.board[row] = newRow
		for col, char := range line {
			grid.board[row][col] = string(char)
		}
	}

	return grid
}

func cloneGrid(grid Grid) Grid {
	cols := grid.cols
	rows := grid.rows
	newGrid := Grid{
		cols:  cols,
		rows:  rows,
		board: make([][]string, rows),
	}

	for row := 0; row < rows; row++ {
		newRow := make([]string, cols)
		copy(newRow, grid.board[row])
		newGrid.board[row] = newRow
	}

	return newGrid
}

func main() {
	fmt.Println("-- Day 18 --")

	fileName := "input.txt"
	useTest := flag.Bool("test", false, "use test file for input")
	ticks := flag.Int("t", 100, "ticks to run simulation")
	flag.Parse()

	if *useTest {
		fileName = "test.txt"
	}

	fmt.Println("File Name:", fileName)
	fmt.Println("Ticks:", *ticks)

	dir, err := os.Getwd()
	checkFatal(err, "error getting working directory")
	file, err := os.ReadFile(dir + "/" + fileName)
	checkFatal(err, "error reading input file")
	lines := strings.Split(string(file), "\n")

	grid := buildGrid(lines, *useTest)
	grid2 := cloneGrid(grid)

	for counter := 0; counter < *ticks; counter++ {
		grid.ProcessTick(false)
	}
	fmt.Printf("*** Lights on ( %d ) ***\n", grid.LightsOn())

	// Turn on corners
	grid2.board[0][0] = GRID_ON
	grid2.board[0][grid2.cols-1] = GRID_ON
	grid2.board[grid2.rows-1][0] = GRID_ON
	grid2.board[grid2.rows-1][grid2.cols-1] = GRID_ON

	for counter := 0; counter < *ticks; counter++ {
		grid2.ProcessTick(true)
	}
	fmt.Printf("*** PART TWO: Lights on ( %d ) ***\n", grid2.LightsOn())
}
