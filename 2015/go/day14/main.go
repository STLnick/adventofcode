package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type reindeer struct {
    name string
    speed int
    stamina int
    cooldown int
    state string
    counter int
    traveled int
    points int
}

func (r *reindeer) Print() {
    fmt.Printf("%-10s", r.name)
    fmt.Printf("%-3d", r.speed)
    fmt.Printf("%-3d", r.stamina)
    fmt.Printf("%-4d", r.cooldown)
    fmt.Printf("%-10s", r.state)
    fmt.Printf("%-4d", r.counter)
    fmt.Printf("%-8d", r.traveled)
    fmt.Printf("%-6d", r.points)
    fmt.Printf("\n")
}

func printLogHead() {
    fmt.Printf("%-10s", "Name")
    fmt.Printf("%-3s", "Spd")
    fmt.Printf("%-3s", "Stm")
    fmt.Printf("%-4s", "CD")
    fmt.Printf("%-10s", "State") 
    fmt.Printf("%-4s", "Ct")
    fmt.Printf("%-8s", "Travel")
    fmt.Printf("%-6s", "Points")
    fmt.Printf("\n")
}

// Rudolph can fly 22 km/s for 8 seconds, but then must rest for 165 seconds.
// [0]             [3]         [6]                                [13]
const SPEED_IDX int = 3
const STAMINA_IDX int = 6
const COOLDOWN_IDX int = 13

func main() {
	fmt.Println("-- Day 14 --")

	input := flag.String("i", "input.txt", "The file to read as input to program")
	limit := flag.Int("l", 1000, "How many ticks to run in simulation")
	flag.Parse()

    fmt.Println("Input: ", *input)
    fmt.Println("Limit: ", *limit)
	
    directory, err := os.Getwd()
	if err != nil {
		log.Fatal("error reading working directory", err)
	}

	path := directory + "/" + *input
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("error reading input file", err)
	}

	lines := strings.Split(string(file), "\n")
    reindeerList := make([]*reindeer, 0, len(lines))
	var (
        lineParts []string
        speed int
        stamina int
        cooldown int
    )

	for _, l := range lines {
		if l == "" {
			continue
		}
		lineParts = strings.Fields(l)
        speed, _ = strconv.Atoi(lineParts[SPEED_IDX])
        stamina, _ = strconv.Atoi(lineParts[STAMINA_IDX])
        cooldown, _ = strconv.Atoi(lineParts[COOLDOWN_IDX])

        newReindeer := &reindeer{
            name: lineParts[0],
            speed: speed,
            stamina: stamina,
            cooldown: cooldown,
            state: "FLYING",
            counter: stamina,
            traveled: 0,
            points: 0,
        }
        reindeerList = append(reindeerList, newReindeer)
	}

    ticks := 0
    var currentLeadDist int
    for ticks < *limit {
        ticks++

        for _, r := range reindeerList {
            if r.state == "FLYING" {
                r.traveled += r.speed
                if r.traveled > currentLeadDist {
                    currentLeadDist = r.traveled
                }
                r.counter--
                if r.counter == 0 {
                    r.state = "RESTING"
                    r.counter = r.cooldown
                }
            } else {
                r.counter--
                if r.counter == 0 {
                    r.state = "FLYING"
                    r.counter = r.stamina
                }
            }
        }

        // Check after everyone has advanced
        for _, r := range reindeerList {
            if r.traveled == currentLeadDist {
                r.points++
            }
        }
    }

    var winning int
    var winningName string
    var winningPts int
    var winningNamePts string

    fmt.Println()
    printLogHead()
    for _, r := range reindeerList {
        r.Print()
        if r.traveled > winning {
            winning = r.traveled
            winningName = r.name
        }
        if r.points > winningPts {
            winningPts = r.points
            winningNamePts = r.name
        }
    }
    fmt.Printf("\n%s traveled the furthest -> %d\n", winningName, winning)
    fmt.Printf("\n%s had the most points -> %d\n", winningNamePts, winningPts)
}
