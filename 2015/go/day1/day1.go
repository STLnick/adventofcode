package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Hello, Day 1")
    var directions string
    basementStep := -1
    floor := 0

    fmt.Print("Reading input...")
    if len(os.Args) == 1 {
        reader := bufio.NewReader(os.Stdin)
        directions, _ = reader.ReadString('\n')
    } else {
        directions = os.Args[1]
    }

    fmt.Print(" ...Calculating...\n")
    for i, step := range directions {
        if step == '(' {
            floor += 1
        } else if step == ')' {
            floor -= 1
        }

        if floor < 0 && basementStep == -1 {
            basementStep = i + 1
        }
    }

    fmt.Printf("ANSWER: floor (%d) - first entered basement on step (%d)", floor, basementStep)
}
