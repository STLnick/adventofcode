package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    fmt.Println("-- Day 6 --")

    scanner := bufio.NewScanner(os.Stdin)
    grid := [1000][1000]int{}
    var (
        currInstrStr string
        commandVal int
        idx int
        startRange []string
        startX int
        startY int
        endRange []string
        endX int
        endY int
    )

    for scanner.Scan() {
        currInstrStr = scanner.Text()
        splitInstr := strings.Split(currInstrStr, " ")

        if splitInstr[0] == "toggle" {
            commandVal = 2
            idx = 1
        } else {
            if splitInstr[1] == "on" {
                commandVal = 1
            } else {
                commandVal = -1
            }
            idx = 2
        }

        startRange = strings.Split(splitInstr[idx], ",")
        startX, _ = strconv.Atoi(startRange[0])
        startY, _ = strconv.Atoi(startRange[1])
        endRange = strings.Split(splitInstr[idx + 2], ",")
        endX, _ = strconv.Atoi(endRange[0])
        endY, _ = strconv.Atoi(endRange[1])

        for x := startX; x <= endX; x++ {
            for y := startY; y <= endY; y++ {
                grid[x][y] += commandVal
                if grid[x][y] < 0 {
                    grid[x][y] = 0
                }
            }
        }
    }

    brightness := 0
    for _, row := range grid {
        for _, val := range row {
            brightness += val
        }
    }

    fmt.Println("Total brightness: ", brightness)
}
