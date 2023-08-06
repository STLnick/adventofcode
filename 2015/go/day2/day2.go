package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
    fmt.Println("Day 2")
    var measurements string
    total := 0
    ribbonTotal := 0
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        measurements = scanner.Text()
        mSlice := strings.Split(measurements, "x")
        if len(mSlice) != 3 {
            log.Fatal("Must provide proper measurements in LxWxH format")
        }

        l, _ := strconv.Atoi(mSlice[0])
        w, _ := strconv.Atoi(mSlice[1])
        h, _ := strconv.Atoi(mSlice[2])

        var min float64
        lw := l * w
        wh := w * h
        hl := h * l
        min = math.Min(float64(lw), float64(wh))
        min = math.Min(min, float64(hl))
        area := 2 * (lw + wh + hl)
        total += area + int(min)

        var minPerim int

        if l < w {
            if w < h {
                minPerim = 2*l + 2*w
            } else {
                minPerim = 2*l + 2*h
            }
        } else {
            if l < h {
                minPerim = 2*w + 2*l
            } else {
                minPerim = 2*w + 2*h
            }
        }

        ribbonTotal += (l * w * h) + minPerim
    }


    fmt.Printf("Needed surface area of wrapping paper: %d feet\n", total)
    fmt.Printf("Needed length of ribbon: %d feet\n", ribbonTotal)
}
