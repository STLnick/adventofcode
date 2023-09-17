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

					iCap += (a * ingList[0].capacity)
					iDur += (a * ingList[0].durability)
					iFlav += (a * ingList[0].flavor)
					iTex += (a * ingList[0].texture)
					iCal += (a * ingList[0].calories)

					iCap += (b * ingList[1].capacity)
					iDur += (b * ingList[1].durability)
					iFlav += (b * ingList[1].flavor)
					iTex += (b * ingList[1].texture)
					iCal += (b * ingList[1].calories)

					iCap += (c * ingList[2].capacity)
					iDur += (c * ingList[2].durability)
					iFlav += (c * ingList[2].flavor)
					iTex += (c * ingList[2].texture)
					iCal += (c * ingList[2].calories)

					iCap += (d * ingList[3].capacity)
					iDur += (d * ingList[3].durability)
					iFlav += (d * ingList[3].flavor)
					iTex += (d * ingList[3].texture)
					iCal += (d * ingList[3].calories)
		
                    if iCap < 0 || iDur < 0 || iFlav < 0 || iTex < 0 {
                        score = 0
                    } else {
                        score = iCap * iDur * iFlav * iTex
                    }
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
