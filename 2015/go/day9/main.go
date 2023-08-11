package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

/*
   London to Dublin = 464
   London to Belfast = 518
   London to XYZ = 100
   Dublin to Belfast = 141
   Dublin to XYZ = 125
   XYZ to Belfast = 50

   Possibilities:
   London:
    - Dublin
        - Belfast
            - XYZ
        - XYZ
            - Belfast
    - Belfast
        - Dublin
            - XYZ
        - XYZ
            - Dublin
    - XYZ
        - Dublin
            - Belfast
        - Belfast
            - Dublin

   [key1][key2] = value
   [ ] [L]  [B]  [D]  [X]
   [L] 0    518  464  100
   [B] 518  0    141  50
   [D] 464  141  0    125
   [X] 100  50   125  0

   What is the shortest route's distance value? (6 possible routes)
   L => D => B = 605
*/

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

func findSolution(solutions *[]int, dm DistanceMap, current string, remaining []string, sum int) {
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
        findSolution(solutions, dm, k, removeKey(remaining, k), sum + dm[current][k])
    }
}

func findSolutions(dm DistanceMap, current string, remaining []string) []int {
    var solutions []int
    findSolution(&solutions, dm, current, removeKey(remaining, current), 0)
    return solutions
}

const loopLogging = false

func main() {
	fmt.Println("-- Day 9 --")

	scanner := bufio.NewScanner(os.Stdin)
    distanceMap := make(DistanceMap)
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

        _, ok = distanceMap[key1]
        if !ok {
            keys = append(keys, key1)
            distanceMap[key1] = make(map[string]int)
        }

        _, ok = distanceMap[key2]
        if !ok {
            keys = append(keys, key2)
            distanceMap[key2] = make(map[string]int)
        }
 
        distanceMap[key1][key2] = val
        distanceMap[key2][key1] = val
	}

    useConcurrency := true
    var solutions []int

    if !useConcurrency {
        fmt.Println("SYNCHRONOUS RUN")
        for _, k := range keys {
            if loopLogging {
                fmt.Println("MAIN LOOP :: Key: ", k)
            }
            findSolution(&solutions, distanceMap, k, removeKey(keys, k), 0)
        }
    } else {
        fmt.Printf("-- Running %v Go Routines\n", len(keys))

        var wg sync.WaitGroup
        for _, k := range keys {
            wg.Add(1)
            go func(key string, keys []string) {
                data := findSolutions(distanceMap, key, removeKey(keys, key))
                solutions = append(solutions, data...)
                wg.Done()
            }(k, keys)
        }

        wg.Wait()
    }

    var lowest int 
    var highest int 
    for i, v := range solutions {
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

    fmt.Printf("\t***** Calculated %v solutions ***\n", len(solutions))
    fmt.Printf("\t  ===> Lowest value: %v\n", lowest)
    fmt.Printf("\t  ===> Highest value: %v\n", highest)
	fmt.Println("-- END --")
}
