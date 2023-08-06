package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Contains(haystack *[]string, needle string) bool {
	for _, s := range *haystack {
		if s == needle {
			return true
		}
	}

	return false
}

func ContainsInt(haystack *[]int, needle int) bool {
	for _, n := range *haystack {
		if n == needle {
			return true
		}
	}

	return false
}

func AddDistinct(slice *[]string, s string) {
	if !Contains(slice, s) {
		*slice = append(*slice, s)
	}
}

func SpliceMany(s []string, indices []int) []string {
    spliced := []string{}
    for i := 0; i < len(s); i++ {
        if !ContainsInt(&indices, i) {
            spliced = append(spliced, s[i])
        }
    }
    return spliced
}

func RunInstructions(instructions []string, wires map[string]uint16) {
	var (
		splitInstr []string
		dest       string
		op         string
		argOne     string
		argTwo     string
		valOne     uint16
		valTwo     uint16
		err        error
		temp       int
		ok         bool
	)
	indicesToSplice := []int{}
	wireNames := []string{}

	for len(instructions) > 0 {
        indicesToSplice = []int{}

		for i, instr := range instructions {
			splitInstr = strings.Split(instr, " ")

			switch len(splitInstr) {
			case 3: // num
				dest = splitInstr[2]

                if dest == "b" {
                    // Override for Part 2 with "a" value from Part 1
                    valOne = 956
                } else if unicode.IsNumber(rune(splitInstr[0][0])) {
					temp, err = strconv.Atoi(splitInstr[0])
					if err != nil {
						log.Fatal(err)
					}
					valOne = uint16(temp)
				} else {
					valOne, ok = wires[splitInstr[0]]
					if !ok {
						continue
					}
				}

				indicesToSplice = append(indicesToSplice, i)
				AddDistinct(&wireNames, dest)
				wires[dest] = valOne
			case 4: // NOT {id}
				valOne, ok = wires[splitInstr[1]]
				if !ok {
					continue
				}

				valOne = ^valOne
				indicesToSplice = append(indicesToSplice, i)
				dest = splitInstr[3]
				AddDistinct(&wireNames, dest)
				wires[dest] = valOne
			case 5: // arg - command - arg
				argOne = splitInstr[0]
				op = splitInstr[1]
				argTwo = splitInstr[2]
				dest = splitInstr[4]
				AddDistinct(&wireNames, dest)

				if unicode.IsNumber(rune(argOne[0])) {
					temp, err = strconv.Atoi(argOne)
					if err != nil {
						log.Fatal(err)
					}
					valOne = uint16(temp)
				} else {
					valOne, ok = wires[argOne]
					if !ok {
						continue
					}
				}

				if unicode.IsNumber(rune(argTwo[0])) {
					temp, err = strconv.Atoi(argTwo)
					if err != nil {
						log.Fatal(err)
					}
					valTwo = uint16(temp)
				} else {
					valTwo, ok = wires[argTwo]
					if !ok {
						continue
					}
				}

				indicesToSplice = append(indicesToSplice, i)

				switch op {
				case "AND":
					wires[dest] = uint16(valOne & valTwo)
				case "OR":
					wires[dest] = uint16(valOne | valTwo)
				case "LSHIFT":
					wires[dest] = uint16(valOne << valTwo)
				case "RSHIFT":
					wires[dest] = uint16(valOne >> valTwo)
				}
			}
		}

        instructions = SpliceMany(instructions, indicesToSplice)
	}
}

func main() {
	fmt.Println("-- Day 7 --")

	scanner := bufio.NewScanner(os.Stdin)
	instructions := []string{}
	wires := make(map[string]uint16)
	//wires2 := make(map[string]uint16)

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

    RunInstructions(instructions, wires)

	fmt.Println("[1] Wire A value with overridden b: ", wires["a"])
}
