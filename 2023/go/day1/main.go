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

type NumEntry struct {
    value int
    index int
}

func handleError(e error) {
    if e != nil {
        log.Panic(e)
    }
}

var lookupTable = map[string]int{
    "one": 1, 
    "two": 2, 
    "three": 3, 
    "four": 4, 
    "five": 5, 
    "six": 6, 
    "seven": 7, 
    "eight": 8, 
    "nine": 9, 
}

func getDigitString(s string) int {
    if len(s) == 1 && unicode.IsDigit(rune(s[0])) {
        val, err := strconv.Atoi(s)
        handleError(err)
        return val
    }

    val, ok := lookupTable[s]
    if ok {
        return val
    }

    return -1
}

func main() {
    input := flag.String("i", "input.txt", "The file to be read as input for program")
    flag.Parse()

    directory, err := os.Getwd()
    handleError(err)
    
    file, err := os.ReadFile(directory + "/" + *input)
    handleError(err)

    lines := strings.Split(string(file), "\n")
    var (
        first NumEntry
        last NumEntry
        idx int
        numEntries []NumEntry
        sum int
    )

    for _, l := range lines {
        if l == "" {
            continue;
        }
                
        numEntries = []NumEntry{}

        // Find words of number in string
        for numWord, numVal := range lookupTable {
            idx = strings.Index(l, numWord)
            if idx != -1 {
                numEntries = append(numEntries, NumEntry{value: numVal, index: idx})
            }

            idx = strings.LastIndex(l, numWord)
            if idx != -1 {
                numEntries = append(numEntries, NumEntry{value: numVal, index: idx})
            }
        }

        // Find digits in string
        for cIdx, c := range l {
            if unicode.IsDigit(c) {
                numEntries = append(numEntries, NumEntry{value: int(c - '0'), index: cIdx})
            }
        }

        // Determine which is first and last entry
        for idx, entry := range numEntries {
            if idx == 0 {
                first = entry
                last = entry
            }

            if entry.index < first.index {
                first = entry
            }

            if entry.index > last.index {
                last = entry
            }
        }

        if first.value == 0 {
            log.Panic("no digits found in input string:", l)
        }

        sum += (first.value * 10) + last.value
    }

    /**LOG*/ fmt.Println("Sum of line digits:", sum)
}
