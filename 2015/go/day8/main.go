package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("-- Day 8 --")

	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner

	codeChars := 0
	encodedChars := 0
	memChars := 0
	var (
		idx int
		s   string
	)

	for scanner.Scan() {
		s = scanner.Text()
		idx = 0
        var c rune

		for idx < len(s) {
            c = rune(s[idx])
			if c == '\\' {
				switch rune(s[idx+1]) {
				case '\\':
					codeChars += 2
					encodedChars += 4
					memChars++
					idx += 2
				case '"':
					codeChars += 2
					encodedChars += 4
					memChars++
					idx += 2
				case 'x':
					codeChars += 4
					encodedChars += 5
					memChars++
					idx += 4
				}
            } else if c == '"' {
                codeChars++
                encodedChars += 2
                idx++
			} else {
                codeChars++
                encodedChars++
                memChars++
				idx++
			}
		}

        encodedChars += 2 // surrounding quotes ""
	}

	fmt.Printf("Number of Code Characters: \t%d\n", codeChars)
	fmt.Printf("Number of Memory Characters: \t%d\n", memChars)
    fmt.Printf("Code - Memory = %d\n", codeChars - memChars)
    fmt.Println("- - - - - - - - - - - - - - - - -")
	fmt.Printf("Number of _Encoded_ Characters: \t%d\n", encodedChars)
    fmt.Printf("Encoded - Code = %d\n", encodedChars - codeChars)
}
