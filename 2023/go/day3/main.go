package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func handleError(e error) {
  if e != nil {
    log.Panic(e)
  }
}

func getFileAsLines() []string {
  useTest := flag.Bool("t", false, "Flag to use test input file")
  flag.Parse()

  filename := "input.txt"
  if *useTest {
    filename = "input-test.txt"
  }

  pwd, err := os.Getwd()  
  handleError(err) 

  file, err := os.ReadFile(pwd + "/" + filename)
  handleError(err) 

  return strings.Split(string(file), "\n")
}

func convertFileToGrid() {
}

func isSchematicSymbol(r rune) bool {
  return r != period && (unicode.IsSymbol(r) || unicode.IsPunct(r))
}

func getNumberIndices(line string, start int) (int, int) {
  startX := start
  endX := start

  for startX >= 0 && unicode.IsDigit(rune(line[startX])) {
    startX--
  }
  // Move forward one to be at first digit position
  startX++
  for endX < len(line) && unicode.IsDigit(rune(line[endX])) {
    endX++
  }
  // Move backward one to be at last digit position
  endX--

  return startX, endX
}

const asterisk rune = '*'
const period rune = '.'

func main() {
  /**LOG*/ fmt.Println("Day 3")

  lines := getFileAsLines()
  gridChecks := [][]int{
    []int{-1, -1},
    []int{-1, 0},
    []int{-1, 1},
    []int{0, -1},
    []int{0, 0},
    []int{0, 1},
    []int{1, -1},
    []int{1, 0},
    []int{1, 1},
  }
  var (
    adjacentNums []int
    checkXIdx int
    checkYIdx int
    endX int
    gearRatioSum int
    startX int
    sum int
  )
  prevNum := -1

  for lIdx, line := range lines {
    for rIdx, r := range line {
      adjacentNums = adjacentNums[:0]

      if isSchematicSymbol(r) {
        // Part One
        // Found a symbol - check surrounding squares for a number
        for gridIdx, pair := range gridChecks {
          checkXIdx = rIdx + pair[1]
          checkYIdx = lIdx + pair[0]
          testLine := lines[checkYIdx]

          if gridIdx % 3 == 0 {
            prevNum = -1
          }

          if unicode.IsDigit(rune(testLine[checkXIdx])) {
            startX, endX = getNumberIndices(testLine, checkXIdx)
            num, err := strconv.Atoi(testLine[startX:endX+1])
            handleError(err)

            if num != prevNum {
              sum += num
              adjacentNums = append(adjacentNums, num)
            }
            prevNum = num
          }
        }

        // Part Two
        if r == asterisk && len(adjacentNums) == 2 {
          gearRatioSum += adjacentNums[0] * adjacentNums[1]
        }
      }
    }
  }

  /**LOG*/ fmt.Println("Sum of Part Numbers:", sum)
  /**LOG*/ fmt.Println("Sum of Gear Ratios:", gearRatioSum)
}
