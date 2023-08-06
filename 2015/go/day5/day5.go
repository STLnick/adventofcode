package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pairs struct {
    Log map[string][]int
}

func NewPairs() *Pairs {
    return &Pairs{}
}

func (p *Pairs) Add(ch1 rune, ch2 rune, ch2Idx int) int {
    s := string(ch1) + string(ch2)
    currLen := len(p.Log[s])

    if currLen > 0 {
        lastEntry := p.Log[s][currLen - 1]
        if lastEntry == ch2Idx - 1 {
            return currLen
        }
    }

    p.Log[s] = append(p.Log[s], ch2Idx)

    return currLen + 1
}

func (p *Pairs) Reset() {
    p.Log = make(map[string][]int)
}

func main() {
	fmt.Println("-- Day 5 --")

	niceCt := 0
	scanner := bufio.NewScanner(os.Stdin)
    pairLog := NewPairs()
	var (
		hasPair        bool
		hasSplitDouble bool
		tailCh         rune
		midCh          rune
		currCh         rune
	)

	for scanner.Scan() {
		s := scanner.Text()
		hasPair = false
		hasSplitDouble = false
        pairLog.Reset()

		midCh = rune(s[0])
		currCh = rune(s[1])
        pairLog.Add(midCh, currCh, 1)

		for i := 2; i < len(s); i++ {
			tailCh = midCh
			midCh = currCh
			currCh = rune(s[i])

            if pairLog.Add(midCh, currCh, i) > 1 {
                hasPair = true
            }

            if !hasSplitDouble && tailCh == currCh {
                hasSplitDouble = true
            }

            if hasPair && hasSplitDouble {
                break
            }
		}

		if hasPair && hasSplitDouble {
			niceCt++
		}
	}

	fmt.Println("Nice strings: ", niceCt)
}
