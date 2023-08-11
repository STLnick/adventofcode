package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func findSolution(solutions *[]int, dm DistanceMap, current string, remaining []string, sum int, indent int) {
    indentStr := strings.Repeat("\t", indent)

    fmt.Printf("%s[ %s ]\n", indentStr, current)

    if len(remaining) == 0 {
        fmt.Printf("%s+++ Possible Solution: %v\n", indentStr, sum)
        *solutions = append(*solutions, sum)
        return;
    }

    for _, k := range remaining {
        fmt.Printf("%s--- Moving from %s ---> to ---> %s = +%v (SUM %v)\n", indentStr, current, k, dm[current][k], sum + dm[current][k])
        findSolution(solutions, dm, k, removeKey(remaining, k), sum + dm[current][k], indent + 1)
    }
}

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

    fmt.Println("Keys: ", keys)
    fmt.Printf("Keys Addr : %p\n\n", &keys)

    var solutions []int
    for _, k := range keys {
        fmt.Println("MAIN LOOP :: Key: ", k)
        findSolution(&solutions, distanceMap, k, removeKey(keys, k), 0, 0)
    }

    var lowest int 
    for i, v := range solutions {
        if i == 0 {
            lowest = v
        }
        if v < lowest {
            lowest = v
        }
    }

    fmt.Printf("Calculated %v solutions :: Lowest value was %v\n", len(solutions), lowest)

	fmt.Println("-- END --")
}
