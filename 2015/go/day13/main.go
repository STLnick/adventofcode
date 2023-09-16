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

func buildKeyAndList(instructions *[]instruction) (map[string]map[string]int16, []string) {
	key := make(map[string]map[string]int16)
	var names []string

	for _, ins := range *instructions {
		if !slices.Contains(names, ins.p1) {
			names = append(names, ins.p1)
			key[ins.p1] = make(map[string]int16)
		}

		if !slices.Contains(names, ins.p2) {
			names = append(names, ins.p2)
			key[ins.p2] = make(map[string]int16)
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

func removeName(n string, s *[]string) {
	i := slices.Index(*s, n)
    if i != -1 {
	    *s = slices.Replace(*s, i, i+1)
    }
}

func main() {
	fmt.Println("-- Day 13 --")

	input := flag.String("input", "input.txt", "The file to read as input to program")
	flag.Parse()

	wg := &sync.WaitGroup{}
	resultChanStr := make(chan []string)
	instructions := getInstructions(*input)
	peopleKey, names := buildKeyAndList(&instructions)
    _ = peopleKey // appease compiler for now 

    remaining := make([]string, len(names))
    copy(remaining, names)

	for _, n := range names {
		wg.Add(1)
		
        go func(name string) {
            defer wg.Done()
            resultChanStr <- []string{name}
        }(n)
	}

	go func() {
		wg.Wait()
		close(resultChanStr)
	}()

    // TODO: Loop All compiled combinations and find highest value
	for val := range resultChanStr {
		fmt.Println("resultChanStr val: ", val)
	}
}
