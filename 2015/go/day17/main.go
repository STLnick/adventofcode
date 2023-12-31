package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " :: ", err)
	}
}

func setup() ([]int, int) {
	input := flag.String("i", "input.txt", "input to program")
	total := flag.Int("t", 150, "total to start")
	flag.Parse()

	dir, err := os.Getwd()
	checkFatal(err, "error getting working directory")

	file, err := os.ReadFile(dir + "/" + *input)
	checkFatal(err, "error opening input file")

	strList := strings.Split(string(file), "\n")
	var intVal int
	var intList []int

	for _, val := range strList {
		if val == "" {
			continue
		}
		intVal, err = strconv.Atoi(val)
		checkFatal(err, "error parsing int from string")
		intList = append(intList, intVal)
	}

	return intList, *total
}

func getCount(total int, n int, i int) int {
	if n < 0 {
		return 0
	} else if total == 0 {
        key := len(containers) - n
		if _, ok := solutionMap[key]; !ok {
			solutionMap[key] = 0
		}
		solutionMap[key]++
		return 1
	} else if i == len(containers) || total < 0 {
		return 0
	}

	return getCount(total, n, i+1) + getCount(total-containers[i], n-1, i+1)
}

var containers []int
var solutionMap map[int]int

func main() {
	fmt.Println("-- Day 17 --")
	var total int
	solutionMap = make(map[int]int)
	containers, total = setup()

    // Part One
	count := getCount(total, len(containers), 0)
	fmt.Print("*** Count ", count, " ***\n\n")

    // Part Two
    for k, v := range solutionMap {
        fmt.Printf("%d Containers used in (%d) solutions\n", k, v)
    }
}
