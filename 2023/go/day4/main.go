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

func checkCard(cardStr string) int {
	var points int
	cardParts := strings.Split(cardStr, ": ")
	numberParts := strings.Split(cardParts[1], " | ")
	winningNums := strings.Fields(numberParts[0])
	cardNums := strings.Fields(numberParts[1])

	for _, cardNum := range cardNums {
		//fmt.Println("cardNum:", cardNum)
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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var (
		cardPoints int
		pointSum   int
	)

	for scanner.Scan() {
		cardPoints = checkCard(scanner.Text())
		//fmt.Println("cardPoints", cardPoints)
		pointSum += cardPoints
	}
	fmt.Println("pointSum", pointSum)

	if err := scanner.Err(); err != nil {
		handleError(err)
	}
}
