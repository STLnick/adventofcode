package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"slices"
	"strings"
	"unicode"
)

type transform struct {
	from string
	to   string
}

func (t transform) String() string {
	return fmt.Sprintf("%s => %s", t.from, t.to)
}

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func setup() ([]transform, string) {
	fileName := "input.txt"
	p2 := flag.Bool("p2", false, "use input2 for logic")
	useTest := flag.Bool("test", false, "use test file for input")
	useTest2 := flag.Bool("test2", false, "use test2 file for input")
	flag.Parse()

	if *useTest {
		fileName = "test.txt"
	} else if *useTest2 {
		fileName = "test2.txt"
		isPartTwo = true
	} else if *p2 {
		fileName = "input2.txt"
		isPartTwo = true
	}

	fmt.Println("File Name:", fileName)

	dir, err := os.Getwd()
	checkFatal(err, "error getting working directory")

	file, err := os.ReadFile(dir + "/" + fileName)
	checkFatal(err, "error reading input file")

	lines := strings.Split(string(file), "\n")
	var transforms []transform

	var startIdx int
	for i, line := range lines {
		if line == "" {
			startIdx = i + 1
			break
		}

		parts := strings.Split(line, " => ")
		transforms = append(transforms, transform{from: parts[0], to: parts[1]})
	}

	return transforms, lines[startIdx]
}

func performTransformsRecursive(s string, t transform, start int) string {
	if start >= len(s) {
		return s
	}

	idx := strings.Index(s[start:], t.from)
	if idx == -1 {
		return s
	} else {
		idx += start
	}

	cloned := strings.Clone(s)
	newStr := cloned[:idx] + t.to
	tailStart := idx + len(t.from)

	if tailStart >= len(s) {
		return newStr
	}

	newStr += cloned[tailStart:]
	return newStr + "::" + performTransformsRecursive(s, t, idx+1)
}

var isPartTwo = false
var logger = slog.Default()

func main() {
	fmt.Println("-- Day 19 --")

	transforms, start := setup()

	if !isPartTwo {
		var strs []string

		for _, t := range transforms {
			joinedStr := performTransformsRecursive(start, t, 0)
			for _, s := range strings.Split(joinedStr, "::") {
				if !slices.Contains(strs, s) && s != start {
					strs = append(strs, s)
				}
			}
		}

		fmt.Println("main() :: PART ONE :: # of strs", len(strs))
	} else {
		input := "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl"
		var (
			tokens   []string
			commaCt  int
			parensCt int
			c        rune
		)

		fmt.Printf("Tokens initialized with len (%d) (%v)\n", len(tokens), tokens)

		temp := string(input[0])

		for i := 1; i < len(input); i++ {
			c = rune(input[i])

			if unicode.IsUpper(c) {
				tokens = append(tokens, temp)

				if i == 0 || i == 1 {
					fmt.Printf("Tokens initialized with len (%d) (%v)\n", len(tokens), tokens)
				}

				if temp == "Y" {
					commaCt++
				} else if temp == "Rn" || temp == "Ar" {
					parensCt++
				}
				temp = string(c)
			} else {
				temp += string(c)
				if i == len(input)-1 {
					tokens = append(tokens, temp)
					if temp == "Y" {
						commaCt++
					} else if temp == "Rn" || temp == "Ar" {
						parensCt++
					}
				}
			}
		}

		fmt.Printf("%-8s%5d\n", "Tokens", len(tokens))
		fmt.Printf("%-8s%5d\n", "commaCt", commaCt)
		fmt.Printf("%-8s%5d\n", "parensCt", parensCt)

		// count(tokens) - count(parens) - 2*count("Y") - 1
		answer := len(tokens) - parensCt - (2 * commaCt) - 1

		fmt.Printf("%-8s%5d\n", "Answer", answer)
	}
}
