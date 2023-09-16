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

type seatedPerson struct {
	name    string
	prevVal int16
	nextVal int16
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

func seatPerson(
	n string,
	remaining *[]string,
	seated *[]*seatedPerson,
) *seatedPerson {
	newSp := seatedPerson{name: n}
	removeName(newSp.name, remaining)
	*seated = append(*seated, &newSp)
    return (*seated)[len(*seated) - 1]
}

func calculateChange(seated *[]*seatedPerson) int16 {
	var change int16
	for _, sp := range *seated {
        //fmt.Printf("Prev val: (%d)  ::  Next val: (%d)\n", (*sp).prevVal, (*sp).nextVal)
		change += (*sp).prevVal + (*sp).nextVal
	}
    //fmt.Printf("Calculating change for %d seated people: %d\n", len(*seated), change)
	return change
}

func runPossibility(
	c chan int16,
	peopleKey *map[string]map[string]int16,
	seated []*seatedPerson,
	remaining []string,
	firstSeating string,
    isFirstRun bool,
) {
    if isFirstRun {
        fmt.Println("Remaining to begin parent loop:", remaining)
    }

	var (
		prevSp *seatedPerson
		currSp *seatedPerson
	)

	if len(seated) > 0 {
		prevSp = seated[len(seated) - 1]
	}

    currSp = seatPerson(firstSeating, &remaining, &seated)

	if prevSp != nil {
		prevSp.nextVal = (*peopleKey)[prevSp.name][currSp.name]
		currSp.prevVal = (*peopleKey)[currSp.name][prevSp.name]
	}

    if len(remaining) > 0 {
		lwg := &sync.WaitGroup{}
		for i, rn := range remaining {
            if i == 0 {
                //fmt.Println("Remaining:", remaining)
            }
            tempName := rn
			lwg.Add(1)
			go func(name string) {
                //fmt.Println("runPossibility() :: Name:", name)
				defer lwg.Done()
				runPossibility(c, peopleKey, seated, remaining, name, false)
			}(tempName)
		}
        //go func() {
        //    lwg.Wait()
        //}()
        lwg.Wait()
	} else {
        // Assign values when seating last person
        currSp.nextVal = (*peopleKey)[currSp.name][seated[0].name]
        seated[0].prevVal = (*peopleKey)[seated[0].name][currSp.name]

        c <- calculateChange(&seated)
    }
}

func main() {
	fmt.Println("-- Day 13 --")

	input := flag.String("input", "input.txt", "The file to read as input to program")
	flag.Parse()

	wg := &sync.WaitGroup{}
	resultChan := make(chan int16)
	instructions := getInstructions(*input)
	peopleKey, names := buildKeyAndList(&instructions)
    remaining := make([]string, len(names))
    copy(remaining, names)
	var highest int16
                
	for _, n := range names {
		wg.Add(1)
		
        go func(name string) {
            defer wg.Done()
            r := make([]string, len(names))
            copy(r, remaining)
            initialSeating := make([]*seatedPerson, 0, len(names) + 1)
            runPossibility(
                resultChan,
                &peopleKey,
                initialSeating,
                r,
                name,
                true,
            )
        }(n)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

    var vals []int16
	for val := range resultChan {
		//fmt.Println("resultChan val: ", val)
        vals = append(vals, val)
		if val > highest {
			highest = val
		}
	}

	fmt.Println("\n- - - - - - - - - - - - - - - - - - - -")
	fmt.Println("Highest possible change value: ", highest)
    fmt.Println("Number of values received:", len(vals))
	fmt.Println("- - - - - - - - - - - - - - - - - - - -")
}
