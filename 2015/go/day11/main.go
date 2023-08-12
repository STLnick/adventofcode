package main

import (
	"flag"
	"fmt"
)

func validatePassword(pw string) bool {
	var (
		straightValid bool
		pairsValid    bool
		chRune        rune
		pairs         int
	)
    lastPairEnd := -1
	frame := [3]rune{}
	excludedValid := true

	for i, ch := range pw {
		chRune = rune(ch)

        if *debug {
            fmt.Printf("Char: %c - char int val: %v - Rune: %c\n", ch, ch, chRune)
        }

		frame[0] = frame[1]
		frame[1] = frame[2]
		frame[2] = chRune

		if chRune == 'i' || chRune == 'l' || chRune == 'o' {
			excludedValid = false
		}

		if i >= 1 && !pairsValid {
			if frame[1] == frame[2] && lastPairEnd != i-1 {
                pairs++
                lastPairEnd = i
                if pairs > 1 {
                    pairsValid = true
                }
			}
		}

		if i >= 2 && !straightValid {
			if frame[1] == frame[0]+1 && frame[2] == frame[1]+1 {
				straightValid = true
			}
		}

        if straightValid && excludedValid && pairsValid {
            return true
        }
	}

    return false
}

func incrementPassword(pw string) string {
    var curr rune
    var newPw string
    incrementing := true 
    pwLen := len(pw)

    for i := pwLen - 1; i >= 0; i-- {
        curr = rune(pw[i])

        if incrementing {
            if curr + 1 > zRuneVal {
                newPw = string('a') + newPw
            } else {
                newPw = string(curr + 1) + newPw
                incrementing = false
            }
        } else {
            newPw = string(curr) + newPw
        }
    }

    return newPw
}

const zRuneVal = 122
var debug *bool

func main() {
	fmt.Println("-- Day 11 --")

	debug = flag.Bool("debug", false, "Flag to print debug statements")
	input := flag.String("input", "vzbxkghb", "The starting input")
	flag.Parse()

    var nextPw string
    testPw := *input

    for nextPw == "" {
        testPw = incrementPassword(testPw)
        
        //fmt.Println("Testing -> ", testPw)

        if validatePassword(testPw) {
            nextPw = testPw
        }
    }

    fmt.Println("NEXT PW: ", nextPw)
}
