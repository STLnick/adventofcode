package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Hello, Nick!")

    if len(os.Args) > 1 {
        fmt.Println("Arg 1: ", os.Args[1])
    }
}
