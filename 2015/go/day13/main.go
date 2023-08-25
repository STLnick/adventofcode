package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Alice would gain 54 happiness units by sitting next to Bob.
//   0          2   3                                     10
// person / lose/gain / amt / otherPerson

type instruction struct {
	p1     string
	action string
	amt    int16
	p2     string
}

type seatedPerson struct {
	name    string
	prevVal int16
	nextVal int16
}

func findHighestGain(key *map[string]map[string]int16, seated string, remaining *[]string) string {
	var n string
	var max int16
	valMap := (*key)[seated]

	for i, name := range *remaining {
		if i == 0 || valMap[name] > max {
			max = valMap[name]
			n = name
		}
	}

	return n
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
	path := "/Users/nickray/Code/adventofcode/2015/go/day13/" + fileName
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
	*s = slices.Replace(*s, i, i+1)
}

func seatPerson(
	n string,
	lastSeated *string,
	remaining *[]string,
	seated *[]*seatedPerson,
	seatsOpen *int,
	currentSeat *int,
) {
	newSp := seatedPerson{name: n}
	*lastSeated = newSp.name
	removeName(*lastSeated, remaining)
	*seated = append(*seated, &newSp)
	*seatsOpen--
	*currentSeat++
}

func calculateChange(seated *[]*seatedPerson) int16 {
	var change int16
	for _, sp := range *seated {
		change += (*sp).prevVal + (*sp).nextVal
	}
	return change
}

func runPossibility(
	firstSeating string,
	peopleKey map[string]map[string]int16,
	names []string,
) int16 {
	seatsOpen := len(names)
	remaining := make([]string, len(names))
	copy(remaining, names)

	var (
		seated     []*seatedPerson
		lastSeated string
		prevSp     *seatedPerson
		currSp     *seatedPerson
	)
	currentSeat := -1

	seatPerson(firstSeating, &lastSeated, &remaining, &seated, &seatsOpen, &currentSeat)
	currSp = seated[currentSeat]

	for seatsOpen > 0 {
		n := findHighestGain(&peopleKey, lastSeated, &remaining)
		seatPerson(n, &lastSeated, &remaining, &seated, &seatsOpen, &currentSeat)
		prevSp = currSp
		currSp = seated[currentSeat]
		prevSp.nextVal = peopleKey[prevSp.name][currSp.name]
		currSp.prevVal = peopleKey[currSp.name][prevSp.name]
	}
	currSp.nextVal = peopleKey[currSp.name][seated[0].name]
	seated[0].prevVal = peopleKey[seated[0].name][currSp.name]

    //for _, sp := range seated {
        //fmt.Println(sp)
    //}

	return calculateChange(&seated)
}

func main() {
	fmt.Println("-- Day 13 --")

	input := flag.String("input", "input.txt", "The file to read as input to program")
	flag.Parse()

	instructions := getInstructions(*input)
	peopleKey, names := buildKeyAndList(&instructions)
	var (
		p       int16
		highest int16
	)

    /*
     * I'm not getting the right answer - highest possibility returned in
     * current logic's state is 649 (too low).
     *
     * I think I need to refactor to run it in a recursive, "try every single
     * configuration possible" way.
     *
     * 1. create a channel that accepts int16 (possible changes)
     * 2. Run X different threads with each name used as a root node
     * 3. Each thread will also exhaust every config and post a val to channel
     * 4. Loop channel values to find highest sent
     */

	for i, n := range names {
        fmt.Println("\n- - - - - - - - - - - - - - - - - - - -")
		if p = runPossibility(n, peopleKey, names); i == 0 || p > highest {
			highest = p
		}
		fmt.Printf("--->  Possibility #%d - %d  <---\n", i+1, p)
	}

    fmt.Println("\n- - - - - - - - - - - - - - - - - - - -")
	fmt.Println("Highest possible change value: ", highest)
    fmt.Println("- - - - - - - - - - - - - - - - - - - -")
}
