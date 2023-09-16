package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type instruction struct {
	p1     string
	action string
	amt    int16
	p2     string
}

const SELF string = "SELF"

func buildKeyAndList(instructions *[]instruction) (map[string]map[string]int16, []string) {
	key := make(map[string]map[string]int16)
	key[SELF] = make(map[string]int16)
	var names []string

	for _, ins := range *instructions {
		if !slices.Contains(names, ins.p1) {
			names = append(names, ins.p1)
			key[ins.p1] = make(map[string]int16)
			key[ins.p1][SELF] = 0
			key[SELF][ins.p1] = 0
		}

		if !slices.Contains(names, ins.p2) {
			names = append(names, ins.p2)
			key[ins.p2] = make(map[string]int16)
			key[ins.p2][SELF] = 0
			key[SELF][ins.p2] = 0
		}

		amt := ins.amt
		if ins.action == "lose" {
			amt = amt * -1
		}

		key[ins.p1][ins.p2] = int16(amt)
	}

	return key, names
}

func getInstructions(fileName string) []instruction {
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal("error reading working directory", err)
	}

	path := directory + "/" + fileName
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("error reading input file", err)
	}

	lines := strings.Split(string(file), "\n")
	var parts [][]string
	var lineParts []string

	for _, l := range lines {
		if l == "" {
			continue
		}
		lineParts = strings.Fields(l)
		parts = append(parts, lineParts)
	}

	var (
		p1           string
		action       string
		amt          int
		p2           string
		instructions []instruction
	)

	for _, p := range parts {
		p1 = p[0]
		action = p[2]
		amt, err = strconv.Atoi(p[3])
		p2, _ = strings.CutSuffix(p[10], ".")

		if err != nil {
			log.Panic("string convert error", err)
		}

		instructions = append(instructions, instruction{
			p1:     p1,
			action: action,
			amt:    int16(amt),
			p2:     p2,
		})
	}

	return instructions
}

func removeName(s []string, n string) []string {
	newSlice := copySlice(s)
	i := slices.Index(newSlice, n)
	if i != -1 {
		return slices.Replace(newSlice, i, i+1)
	}
	return newSlice
}

func copySlice(s []string) []string {
	newSlice := make([]string, len(s), cap(s))
	copy(newSlice, s)
	return newSlice
}

func calculateHappiness(peopleKey map[string]map[string]int16, names []string) int16 {
	var total int16

	for i, name := range names {
		prevIndex := i - 1
		if prevIndex < 0 {
			prevIndex = len(names) - 1
		}
		total += peopleKey[name][names[prevIndex]]

		nextIndex := i + 1
		if nextIndex == len(names) {
			nextIndex = 0
		}
		total += peopleKey[name][names[nextIndex]]
	}

	return total
}

func compileCombination(c chan []string, seated []string, remaining []string, firstSeating string) {
	newSeated := append(copySlice(seated), firstSeating)
	newRemaining := removeName(remaining, firstSeating)

	if len(newRemaining) > 0 {
		lwg := sync.WaitGroup{}
		for _, rn := range newRemaining {
			lwg.Add(1)
			r := copySlice(newRemaining)
			s := copySlice(newSeated)

			go func(name string) {
				defer lwg.Done()
				compileCombination(c, s, r, name)
			}(rn)
		}
		lwg.Wait()
	} else {
		c <- newSeated
	}
}

func getPossibleInts(
	key map[string]map[string]int16,
	names []string,
	withSelf bool,
) ([]int16, int16) {
	var vals []int16

	if withSelf {
		names = append(names, SELF)
	}

	wg := &sync.WaitGroup{}
	resultChan := make(chan []string)
	remaining := copySlice(names)

	for _, n := range names {
		wg.Add(1)
		r := copySlice(remaining)

		go func(name string) {
			defer wg.Done()
			compileCombination(
				resultChan,
				make([]string, 0, len(names)),
				r,
				name,
			)
		}(n)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var possibilities [][]string
	var highest int16
	var tempVal int16

	for list := range resultChan {
		possibilities = append(possibilities, list)
		tempVal = calculateHappiness(key, list)

		if tempVal > highest {
			highest = tempVal
		}
		vals = append(vals, tempVal)
	}

	return vals, highest
}

func main() {
	fmt.Println("-- Day 13 --")

	input := flag.String("input", "input.txt", "The file to read as input to program")
	flag.Parse()

	instructions := getInstructions(*input)
	peopleKey, names := buildKeyAndList(&instructions)

	fmt.Println("*** Part 1 ***")
	partOneVals, highestOne := getPossibleInts(peopleKey, names, false)
    fmt.Printf("# of Results (%d) --- Highest Value (%d)\n", len(partOneVals), highestOne)

	fmt.Println("*** Part 2 ***")
	partTwoVals, highestTwo := getPossibleInts(peopleKey, names, true)
    fmt.Printf("# of Results (%d) --- Highest Value (%d)\n", len(partTwoVals), highestTwo)
}
