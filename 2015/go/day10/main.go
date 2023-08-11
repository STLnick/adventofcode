package main

import (
    "flag"
    "fmt"
    "strconv"
)

func parseChunk(start int, str string) (string, int) {
    strLen := len(str)

    if start == strLen - 1 {
        return "1" + string(str[start]), start + 1
    }

    ch := rune(str[start])
    pos := start + 1
    count := 1

    for i := pos; i < strLen; i++ {
        if ch == rune(str[pos]) {
            count++
            pos++
        } else {
            return strconv.Itoa(count) + string(ch), pos
        }
    }

    return strconv.Itoa(count) + string(ch), pos
}

func parseString(s string) string {
    var (
        result string
        chunk  string
        pos    int
    )

    for pos < len(s) {
        chunk, pos = parseChunk(pos, s)
        result += chunk
    }

    return result
}

func main() {
    fmt.Println("-- Day 10 --")

    input := flag.String("in", "1321131112", "Input number string to start with")
    count := flag.Int("n", 40, "Iteration count to parse")
    flag.Parse()

    final := *input
    for i := 0; i < *count; i++ {
        final = parseString(final)
    }

    fmt.Printf("Input:\t\t%s\n", *input)
    //fmt.Printf("Result after %d iterations:\t%s\n", *count, final)
    fmt.Printf("Length of result: %v\n", len(final))
}
