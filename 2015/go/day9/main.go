package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type DistanceMap map[string]map[string]int

func removeKey(s []string, target string) []string {
    var res []string
    for _, k := range s {
        if k != target {
            res = append(res, k)
        }
    }

    return res
}

func findSolution(solutions *[]int, current string, remaining []string, sum int) {
    if loopLogging {
        fmt.Printf("[ %s ]\n", current)
    }

    if len(remaining) == 0 {
        if loopLogging {
            fmt.Printf("+++ Possible Solution: %v\n", sum)
        }
        *solutions = append(*solutions, sum)
        return;
    }

    for _, k := range remaining {
        k := k
        if loopLogging {
            fmt.Printf("--- Moving from %s ---> to ---> %s = +%v (SUM %v)\n", current, k, dm[current][k], sum + dm[current][k])
        }
        findSolution(solutions, k, removeKey(remaining, k), sum + dm[current][k])
    }
}

func findSolutions(current string, remaining []string, start int) []int {
    var solutions []int
    findSolution(&solutions, current, removeKey(remaining, current), start)
    return solutions
}

func findSolutionsConcurrent(current string, remaining []string) []int {
    var solutions []int
    var wg sync.WaitGroup

    // I know my current key
    // Sum == 0 right now

    for _, k := range remaining {
        wg.Add(1)
        go func(_k string, _keys []string) {
            data := findSolutions(_k, _keys, dm[current][_k])
            solutions = append(solutions, data...)
            wg.Done()
        }(k, removeKey(remaining, k))

    }

    wg.Wait()

    //findSolution(&solutions, current, removeKey(remaining, current), 0)

    return solutions
}

func synchronousSolution(keys []string) []int {
    if loopLogging {
        fmt.Println("SYNCHRONOUS RUN")
    }

    var solutions []int

    for _, k := range keys {
        if loopLogging {
            fmt.Println("MAIN LOOP :: Key: ", k)
        }
        findSolution(&solutions, k, removeKey(keys, k), 0)
    }

    return solutions
}

func concurrentSolution(keys []string) []int {
    if loopLogging {
        fmt.Printf("CONCURRENT RUN --- %v goroutines\n", len(keys))
    }

    var solutions []int
    var wg sync.WaitGroup

    for _, k := range keys {
        wg.Add(1)
        go func(_k string, _keys []string) {
            data := findSolutionsConcurrent(_k, removeKey(_keys, _k))
            solutions = append(solutions, data...)
            wg.Done()
        }(k, keys)
    }

    wg.Wait()
    return solutions
}

func parseSolutions(s []int) {
    var lowest int 
    var highest int 
    for i, v := range s {
        if i == 0 {
            lowest = v
        }
        if v < lowest {
            lowest = v
        }
        if v > highest {
            highest = v
        }
    }

    if loopLogging {
        fmt.Printf("\t***** Calculated %v solutions ***\n", len(s))
        fmt.Printf("\t  ===> Lowest value: %v\n", lowest)
        fmt.Printf("\t  ===> Highest value: %v\n", highest)
    }
}

func setup() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var (
        err  error
		s    []string
		keys []string
		key1 string
		key2 string
		val  int
        ok   bool
	)

	for scanner.Scan() {
		s = strings.Split(scanner.Text(), " ")
		key1 = s[0]
		key2 = s[2]
		val, err = strconv.Atoi(s[4])
        if err != nil {
            panic("Done f'd up")
        }

        _, ok = dm[key1]
        if !ok {
            keys = append(keys, key1)
            dm[key1] = make(map[string]int)
        }

        _, ok = dm[key2]
        if !ok {
            keys = append(keys, key2)
            dm[key2] = make(map[string]int)
        }
 
        dm[key1][key2] = val
        dm[key2][key1] = val
	}

    return keys
}

// Flags
var loopLogging = false

// Globals
var dm = make(DistanceMap)

func main() {
	fmt.Println("-- Day 9 --")

    count := flag.Int("count", 10, "set iteration count for tests")
    enableLogs := flag.Bool("logs", false, "display debugging logs")
    skipSync := flag.Bool("ss", false, "skip sync solution")
	flag.Parse()
    loopLogging = *enableLogs

    keys := setup()
    var numCalc int

    if *skipSync {
        fmt.Print("Sync starting")
        syncStart := time.Now()
        for i := 0; i < *count; i++ {
            if i % 10 == 0 {
                fmt.Print(".")
            }
            ss := synchronousSolution(keys)
            parseSolutions(ss)
        }
        syncDur := time.Since(syncStart)
        fmt.Print(" Done!\n")
        fmt.Printf("\t- Duration:\t%v\n", syncDur)
    }

    conStart := time.Now()
    fmt.Print("Concurrent starting")
    for i := 0; i < *count; i++ {
        if i % 10 == 0 {
            fmt.Print(".")
        }
        cs := concurrentSolution(keys)
        parseSolutions(cs)
    
        if i == 0 {
            numCalc = len(cs)
        }
    }
    conDur := time.Since(conStart)
    fmt.Print(" Done!\n")
    fmt.Printf("\t- Duration:\t%v\n", conDur)

    fmt.Printf("Num of solutions Calculated: %v\n", numCalc)

	fmt.Println("-- END --")
}
