package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var delimiters = [...]rune{'[', ']', '{', '}', ',', ':'}

func isDelimiter(r rune) bool {
	for _, delim := range delimiters {
		if r == delim {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("-- Day 12 -- ")

	input := flag.String("in", "input.txt", "Name of input file for testing")
	flag.Parse()

	fmt.Println("Input: ", *input)
    
    directory, err := os.Getwd()
	if err != nil {
		log.Fatal("error reading working directory", err)
	}

	path := directory + "/" + *input

	file, err := os.ReadFile(path)
	if err != nil {
		log.Panic("Error opening file: ", err)
	}

	var (
		//sum         int
		curr        rune
		inString    bool
		buildingNum bool
		numStr      string
		numVal      int
        level       int
        key         string
	)
    levelSums := make(map[int]int)
    levelRed := make(map[int]bool)

	for i, ch := range string(file) {
		curr = rune(ch)

		if curr == '"' {
			inString = !inString

            if !inString {
                if key == "red" && rune(file[i - 5]) == ':' {
                    levelRed[level] = true
                }
                key = ""
            }

			continue
		}

        if inString {
            key += string(curr)
            continue
        }

        if unicode.IsSpace(curr) {
            continue
        }

		if isDelimiter(curr) { 
			if buildingNum {
				numVal, err = strconv.Atoi(numStr)
				if err != nil {
					log.Panic("error parsing num str", err)
				}

                levelSums[level] += numVal

                buildingNum = false
				numStr = ""
			}

            if curr == '{' {
                level++
            } else if curr == '}' {
                if !levelRed[level] {
                    // Commit sum to parent levelSum if not red
                    levelSums[level - 1] += levelSums[level]
                }

                levelSums[level] = 0
                levelRed[level] = false
                level--
            }

			continue
		}

		buildingNum = true
		numStr += string(curr)
	}

    var final int
    if !levelRed[0] {
        final = levelSums[0]
    }

	fmt.Println("Final SUM: ", final)
}
