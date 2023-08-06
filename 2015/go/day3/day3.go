package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    fmt.Println("--- Day 3 ---")

    scanner := bufio.NewScanner(os.Stdin)
    visited := 1
    x := 0
    y := 0
    xR := 0
    yR := 0
    mapList := make(map[int][]int)
    mapList[x] = append(mapList[x], y)
    santasTurn := false
    var newHouse bool
    var testX int
    var testY int

    for scanner.Scan() {
        directions := scanner.Text()

        for _, move := range directions {
            santasTurn = !santasTurn
            newHouse = true

            switch move {
            case '^':
                if santasTurn {
                    y += 1
                } else {
                    yR += 1
                }
            case '>':
                if santasTurn {
                    x += 1
                } else {
                    xR += 1
                }
            case 'v':
                if santasTurn {
                    y -= 1
                } else {
                    yR -= 1
                }
            case '<':
                if santasTurn {
                    x -= 1
                } else {
                    xR -= 1
                }
            }

            if santasTurn {
                testX = x
                testY = y
            } else {
                testX = xR
                testY = yR
            }

            for _, yCoord := range mapList[testX] {
                if yCoord == testY {
                    newHouse = false
                    break
                }
            }

            if newHouse {
                mapList[testX] = append(mapList[testX], testY)
                visited += 1
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Err: ", err)
    }

    fmt.Printf("Visted houses (%d)\n", visited)
}
