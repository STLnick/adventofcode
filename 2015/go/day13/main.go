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
) {
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
		for _, rn := range remaining {
			lwg.Add(1)
            s := make([]*seatedPerson, len(seated), cap(seated))
            //s := make([]*seatedPerson, 0, cap(seated))
            copy(s, seated)
            r := make([]string, len(remaining))
            copy(r, remaining)

			go func(name string) {
                //fmt.Println("runPossibility() :: Name:", name)
				defer lwg.Done()
				//runPossibility(c, peopleKey, seated, r, name)
                
                // When providing a fresh copy for SEATED getting 0 values back
				runPossibility(c, peopleKey, s, r, name)
			}(rn)
		}
        //go func() {
        //    lwg.Wait()
        //}()
        lwg.Wait()
	} else {
        currSp.nextVal = (*peopleKey)[currSp.name][seated[0].name]
        seated[0].prevVal = (*peopleKey)[seated[0].name][currSp.name]

        c <- calculateChange(&seated)
    }
}

func runPossibilityStr(c chan []string, seated []string, remaining []string, firstSeating string) {
	seated = append(seated, firstSeating)
	removeName(firstSeating, &remaining)

    if len(remaining) > 0 {
		lwg := &sync.WaitGroup{}
		for _, rn := range remaining {
			lwg.Add(1)

            s := make([]string, len(seated), cap(seated))
            copy(s, seated)

            r := make([]string, len(remaining))
            copy(r, remaining)

			go func(name string) {
                //fmt.Println("runPossibilityStr() :: Name:", name)
				defer lwg.Done()
				runPossibilityStr(c, s, r, name)
                
                // When providing a fresh copy for SEATED getting 0 values back
				//runPossibilityStr(c, peopleKey, s, r, name)
			}(rn)
		}

        lwg.Wait()
	} else {
        //c <- calculateChange(&seated)
        c <- seated
    }
}

func main() {
	fmt.Println("-- Day 13 --")

	input := flag.String("input", "input.txt", "The file to read as input to program")
	flag.Parse()

	wg := &sync.WaitGroup{}
	resultChan := make(chan int16)
	//resultChanStr := make(chan []string)
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
            initialSeating := make([]*seatedPerson, 0, len(names))
            //initialSeatingStr := make([]string, 0, len(names))

            runPossibility(
                resultChan,
                &peopleKey,
                initialSeating,
                r,
                name,
            )

            // Provide all possible string/seating combinations

            //runPossibilityStr(resultChanStr, initialSeatingStr, r, name)
        }(n)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

    // TODO: Loop All compiled combinations and find hightest value

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
