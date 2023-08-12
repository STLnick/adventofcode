package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

type parseResult struct {
	idx int
	str string
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func parseChunk(start int, str string) (string, int) {
	strLen := len(str)

	if start == strLen-1 {
		return "1" + string(str[start]), start + 1
	}

	ch := rune(str[start])
	pos := start + 1
	count := 1

	for i := pos; i < strLen; i++ {
		if ch == rune(str[pos]) {
			count++
			pos++
		} else {
			return strconv.Itoa(count) + string(ch), pos
		}
	}

	return strconv.Itoa(count) + string(ch), pos
}

func parseString(s string) string {
	var (
		result string
		chunk  string
		pos    int
	)

	for pos < len(s) {
		chunk, pos = parseChunk(pos, s)
		result += chunk
	}

	return result
}

func main() {
	fmt.Println("-- Day 10 --")

	input := flag.String("in", "1321131112", "Input number string to start with")
	count := flag.Int("n", 40, "Iteration count to parse")
	fractions := flag.Int("f", 2, "Number of chunks to split input into")
	runs := flag.Int("r", 1, "Number of times to run logic")
	flag.Parse()

	runtimes := map[string]time.Duration{
        "chunk": 0,
        "parse": 0,
        "merge": 0,
    }
	start := time.Now()

	var (
		foundNext bool
		ch1       byte
		ch2       byte
		cursor1   int
		cursor2   int
        finalLen  int
	)

    for r := 0; r < *runs; r++ {
        fmt.Print(".")

        inputStr := *input

        for i := 0; i < *count; i++ {
            chunkStart := time.Now()
            var chunks []string
            cursor1 = 0
            cursor2 = 0

            for j := 1; j < *fractions; j++ {
                foundNext = false
                cursor1 = cursor2
                if cursor1 == len(inputStr) {
                    break
                }
                cursor2 = j*(len(inputStr) / *fractions) - 1

                // If last chunk was long enough to reach next calculated starting point
                if cursor2 <= cursor1 {
                    cursor2 = cursor1 + 1
                }

                for !foundNext && cursor2 < len(inputStr) {
                    if cursor2 != len(inputStr)-1 {
                        ch1 = inputStr[cursor2]
                        ch2 = inputStr[cursor2+1]
                        if ch1 != ch2 {
                            foundNext = true
                        }
                    }

                    cursor2++
                }

                chunks = append(chunks, inputStr[cursor1:cursor2])
            }

            // Append last part of string unless done in last loop iteration
            if cursor1 != len(inputStr) {
                chunks = append(chunks, inputStr[cursor2:])
            }

            chunkRuntime := time.Since(chunkStart)
            runtimes["chunk"] += chunkRuntime
            parseStart := time.Now()

            jobs := make(chan parseResult)
            wg := &sync.WaitGroup{}

            for k, c := range chunks {
                wg.Add(1)
                go func(idx int, chunk string) {
                    defer wg.Done()
                    parsed := parseString(chunk)
                    result := parseResult{idx: idx, str: parsed}
                    jobs <- result
                }(k, c)
            }

            go func() {
                wg.Wait()
                close(jobs)
            }()

            resultList := []parseResult{}
            for val := range jobs {
                resultList = append(resultList, val)
            }

            if len(chunks) != len(resultList) {
                panic(
                    fmt.Sprintf("![ERROR]: Lost a chunk from results! Had %d - Have %d", len(chunks), len(resultList)),
                )
            }

            parseRuntime := time.Since(parseStart)
            runtimes["parse"] = runtimes["parse"] + parseRuntime
            mergeStart := time.Now()

            strSlots := make([]string, len(resultList))
            for _, pr := range resultList {
                strSlots[pr.idx] = pr.str
            }
            inputStr = strings.Join(strSlots, "")

            mergeRuntime := time.Since(mergeStart)
            runtimes["merge"] = runtimes["merge"] + mergeRuntime
        }

        if r == 0 {
            finalLen = len(inputStr)
        }
    }

    fmt.Printf("Done!\n")
    fmt.Printf("(INFO) Ran %d Times\n", *runs)
	fmt.Printf("(INFO) Input: %s\n", *input)

	runtime := time.Since(start)
	runtimeNs := float64(runtime.Abs().Nanoseconds())
	chunkNs := runtimes["chunk"].Abs().Nanoseconds()
	parseNs := runtimes["parse"].Abs().Nanoseconds()
	mergeNs := runtimes["merge"].Abs().Nanoseconds()
	chunkDec:= float64(chunkNs) / runtimeNs
	parseDec := float64(parseNs) / runtimeNs
	mergeDec := float64(mergeNs) / runtimeNs

	fmt.Printf("-  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -\n")
    fmt.Printf("\tANSWER: FINAL length of result: %v\n\n", finalLen)
	fmt.Printf("\t[Total Runtime]: \t%v\n", runtime)
	fmt.Printf("\t- [Chunk Runtime]: \t%v\t(%v%%)\n", runtimes["chunk"], roundFloat(float64(100)*chunkDec, 2))
	fmt.Printf("\t- [Parse Runtime]: \t%v\t(%v%%)\n", runtimes["parse"], roundFloat(float64(100)*parseDec, 2))
	fmt.Printf("\t- [Merge Runtime]: \t%v\t(%v%%)\n", runtimes["merge"], roundFloat(float64(100)*mergeDec, 2))
	fmt.Printf("-  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -\n")
}
