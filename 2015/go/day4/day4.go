package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
)

func main() {
    fmt.Println("-- Day 4 --")
    if len(os.Args) == 1 || len(os.Args) > 2 {
        log.Fatal("Provide exactly one argument to run test.\n")
    }

    key := os.Args[1]
    lowestEligible := -1
    var data []byte
    num := 0

    fmt.Print("Testing at 1...")
    for lowestEligible < 0 {
        num += 1

        if num % 10000 == 0 {
            fmt.Printf(" %d...", num)
        }

        data = []byte(key + fmt.Sprint(num))
        sum := md5.Sum(data)
       
        if sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
            // Hash starts with at least 6 leading zeroes
            lowestEligible = num
            fmt.Print("\n\n")
        }
    }

    fmt.Println("Lowest eligible number: ", lowestEligible)
}
