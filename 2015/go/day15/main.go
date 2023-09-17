package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func compileIngredients(fileName string) []ingredient {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	ingredients := make([]ingredient, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		lineParts := strings.Fields(line)
		capacity, _ := strconv.Atoi(strings.Replace(lineParts[2], ",", "", 1))
		durability, _ := strconv.Atoi(strings.Replace(lineParts[4], ",", "", 1))
		flavor, _ := strconv.Atoi(strings.Replace(lineParts[6], ",", "", 1))
		texture, _ := strconv.Atoi(strings.Replace(lineParts[8], ",", "", 1))
		calories, _ := strconv.Atoi(lineParts[10])

		newIng := ingredient{
			name:       lineParts[0],
			capacity:   capacity,
			durability: durability,
			flavor:     flavor,
			texture:    texture,
			calories:   calories,
		}

		ingredients = append(ingredients, newIng)
	}

	return ingredients
}

func addToTotals(ing ingredient, qty int, capT *int, durT *int, flavT *int, texT *int, calT *int) {
    *capT += (qty * ing.capacity)
    *durT += (qty * ing.durability)
    *flavT += (qty * ing.flavor)
    *texT += (qty * ing.texture)
    *calT += (qty * ing.calories)
}

func clampToZero(capT *int, durT *int, flavT *int, texT *int, calT *int) {
    if *capT < 0 {
        *capT = 0
    }
    if *durT < 0 {
        *durT = 0
    }
    if *flavT < 0 {
        *flavT = 0
    }
    if *texT < 0 {
        *texT = 0
    }
}

func main() {
	fmt.Println("-- Day 15 --")

	input := flag.String("i", "input.txt", "Name of input file")
	flag.Parse()

	fmt.Println("Input file:", *input)

	ingList := compileIngredients(*input)

	var (
		iCap    int
		iDur    int
		iFlav   int
		iTex    int
		iCal    int
		highest int
		score   int
		scores  []int
	)
	highestCombo := make([]int, len(ingList))

	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			for c := 0; c < 100; c++ {
				for d := 0; d < 100; d++ {
					if a+b+c+d != 100 {
						continue
					} else {
                        iCap = 0
                        iDur = 0
                        iFlav = 0
                        iTex = 0
                        iCal = 0
                    }

                    addToTotals(ingList[0], a, &iCap, &iDur, &iFlav, &iTex, &iCal)
                    addToTotals(ingList[1], b, &iCap, &iDur, &iFlav, &iTex, &iCal)
                    addToTotals(ingList[2], c, &iCap, &iDur, &iFlav, &iTex, &iCal)
                    addToTotals(ingList[3], d, &iCap, &iDur, &iFlav, &iTex, &iCal)
                    clampToZero(&iCap, &iDur, &iFlav, &iTex, &iCal)
	
                    if iCal != 500 {
                        continue // Part Two
                    }

                    score = iCap * iDur * iFlav * iTex
                    scores = append(scores, score)

                    if score > highest {
                        highest = score
                        copy(highestCombo, []int{a,b,c,d})
                    }
                }
            }
        }
	}

	fmt.Println("# of Scores:", len(scores))
	fmt.Printf("Highest Score: %d --- combo %v\n", highest, highestCombo)
}
