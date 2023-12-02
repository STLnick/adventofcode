package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handleError(e error) {
    if e != nil {
        log.Panic(e)
    }
}

func getInputLinesAsStrings() []string {
    input := flag.String("i", "input.txt", "File to be read as input to program")
    useTestFile := flag.Bool("t", false, "Flag if program should read test file for input")
    flag.Parse()

    var inputFileName string
    if *useTestFile {
        inputFileName = "input-test.txt"
    } else {
        inputFileName = *input
    }

    pwd, err := os.Getwd()
    handleError(err)

    file, err := os.ReadFile(pwd + "/" + inputFileName)
    handleError(err)

    return strings.Split(string(file), "\n")
}

func processLine(line string) (int, int) {
    valid := true
    lineParts := strings.Split(line, ":")
    idStr := lineParts[0]
    dataStr := lineParts[1]
    id, err := strconv.Atoi(strings.Fields(idStr)[1])
    handleError(err)

    maxVals := map[string]int{
        "blue": 0,
        "green": 0,
        "red": 0,
    }
    var (
        color string
        num int
        splitSetChunk []string
    )

    for _, set := range strings.Split(dataStr, ";") {
        for _, splitSet := range strings.Split(set, ", ") {
            splitSetChunk = strings.Fields(splitSet)
            num, err = strconv.Atoi(splitSetChunk[0])
            handleError(err)
            color = splitSetChunk[1]
       
            if num > cubes[color] {
                valid = false
            }

            if num > maxVals[color] {
                maxVals[color] = num
            }
        }
    }
    
    if !valid {
        id = -1
    }

    power := maxVals["blue"] * maxVals["green"] * maxVals["red"]

    return id, power
}

var cubes = map[string]int{
    "blue": 14,
    "green": 13,
    "red": 12,
}

func main() {
    var sum int
    var powerSum int
    lines := getInputLinesAsStrings()

    for _, line := range lines {
        if line == "" {
            continue
        }
        
        id, power := processLine(line)
        powerSum += power
        if id != -1 {
             sum += id
        }
    }

    /**LOG*/ fmt.Println("Sum of Valid IDs: ", sum)
    /**LOG*/ fmt.Println("Sum of Powers: ", powerSum)
}
