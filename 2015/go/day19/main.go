package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type transform struct {
	from string
	to   string
}

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func setup() ([]transform, string) {
	fileName := "input.txt"
	useTest := flag.Bool("test", false, "use test file for input")
	flag.Parse()

	if *useTest {
		fileName = "test.txt"
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

func performTransforms(s string, t transform, start int) string {
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
	return newStr + "::" + performTransforms(s, t, idx+1)
}

func main() {
	fmt.Println("-- Day 19 --")

	transforms, start := setup()
	var strs []string

	for _, t := range transforms {
		joinedStr := performTransforms(start, t, 0)
		for _, s := range strings.Split(joinedStr, "::") {
			if !slices.Contains(strs, s) && s != start {
				strs = append(strs, s)
			}
		}
	}

	fmt.Println("main() :: # of strs", len(strs))
}
