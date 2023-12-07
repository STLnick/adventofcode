package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func handleError(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func parseCardNumbers(cardStr string) ([]string, []string) {
	cardParts := strings.Split(cardStr, ": ")
	numberParts := strings.Split(cardParts[1], " | ")
	winningNums := strings.Fields(numberParts[0])
	cardNums := strings.Fields(numberParts[1])

	return winningNums, cardNums
}

func checkCardPoints(cardStr string) int {
	var points int
	winningNums, cardNums := parseCardNumbers(cardStr)

	for _, cardNum := range cardNums {
		if slices.Contains(winningNums, cardNum) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return points
}

func checkCardMatches(cardStr string) int {
	var matches int
	winningNums, cardNums := parseCardNumbers(cardStr)

	for _, cardNum := range cardNums {
		if slices.Contains(winningNums, cardNum) {
			matches += 1
		}
	}

	return matches
}

func printMap(lineMap map[int]int) {
	//for k, v := range lineMap {
	for i := 0; i < len(lineMap); i++ {
		fmt.Printf("Card %d :: number of cards %d\n", i, lineMap[i])
	}
}

func main() {
	fmt.Println("-- Day 4 --")

	useTestFile := flag.Bool("t", false, "Flag if program should use test file as input")
	flag.Parse()

	filename := "input.txt"
	if *useTestFile {
		filename = "input-test.txt"
	}

	dir, err := os.Getwd()
	handleError(err)

	file, err := os.Open(dir + "/" + filename)
	handleError(err)
	//defer file.Close()

	scanner := bufio.NewScanner(file)
	var (
		cardPoints int
		pointSum   int
	)

	for scanner.Scan() {
		cardPoints = checkCardPoints(scanner.Text())
		pointSum += cardPoints
	}

	fmt.Println("Part 1: pointSum", pointSum)

	if err := scanner.Err(); err != nil {
		handleError(err)
	}
	err = file.Close()
	handleError(err)

	// Part 2
	var cardMatches int
	part2File, err := os.ReadFile(dir + "/" + filename)
	lines := strings.Split(string(part2File), "\n")
	lines = lines[:len(lines)-1]

	lineMap := make(map[int]int)
	for i, line := range lines {
		if line == "" {
			continue
		}
		lineMap[i] = 1
	}

	for i := 0; i < len(lines); i++ {
		cardMatches = checkCardMatches(lines[i])
		if cardMatches == 0 {
			continue
		}

		cardsToPlay := lineMap[i]
		for j := 0; j < cardsToPlay; j++ {
			localMatches := cardMatches
			for idx := i + 1; idx < len(lines) && localMatches > 0; idx++ {
				lineMap[idx] += 1
				localMatches -= 1
			}
		}
	}

	printMap(lineMap)

	var cardCount int
	for i := 0; i < len(lineMap); i++ {
		cardCount += lineMap[i]
	}

	fmt.Println("Part 2: cardCount", cardCount)
}
