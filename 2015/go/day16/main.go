package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//type auntSue struct {
//    id int
//    children int
//    cats int
//    samoyeds int
//    pomeranians int
//    akitas int
//    vizslas int
//    goldfish int
//    trees int
//    cars int
//    perfumes int
//}

type auntSue struct {
	id    int
	props map[string]int
}

func NewSue(id int) auntSue {
	return auntSue{
		id:    id,
		props: make(map[string]int),
	}
}

func compileSues() []auntSue {
	input := flag.String("i", "input.txt", "Input file to pipe to program")
	flag.Parse()

	fmt.Println("Input File:", *input)

	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/" + *input)
	if err != nil {
		log.Fatal(err)
	}

	sues := make([]auntSue, 500)
	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		if line == "" {
			continue
		}

		colonIdx := strings.Index(line, ":")
		sliced := line[colonIdx+2:]
		props := func() [][]string {
			var res [][]string

			propStrs := strings.Split(sliced, ", ")
			for _, propStr := range propStrs {
				res = append(res, strings.Split(propStr, ": "))
			}

			return res
		}()

		var val int
		newSue := NewSue(i + 1)
		for _, prop := range props {
			val, _ = strconv.Atoi(prop[1])
			newSue.props[prop[0]] = val
		}

		sues[i] = newSue
	}

	return sues
}

func main() {
	fmt.Println("-- Day 16 --")
	sues := compileSues()
	output := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	var (
		sueVal        int
		ok            bool
		matches       bool
	)

	for _, sue := range sues {
		matches = true

		for key, val := range output {
			sueVal, ok = sue.props[key]
			if ok && sueVal != val {
				matches = false
                break
			}
		}

		if matches {
            fmt.Println("Matching Sue ID:", sue.id)
            break
		}
	}
}
