package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

func main() {
    part1()
    // part2()
}

func part1() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

    // This the final result
    sum := 0
    blocks := make([][]string, 100)

    n := 0
    for scanner.Scan() {
        if len(scanner.Text()) == 0 {
            n++
            continue
        }
        blocks[n] = append(blocks[n], strings.TrimSuffix(scanner.Text(), "\n"))
    }

    for _, b := range(blocks) {
        // for each line in that block
        for i := 0; i < len(b) - 1; i++ {
            if b[i] == b[i+1] {
                sUp := ""
                sDown := ""
                offset := min(i, len(b) - i - 2)
                for j := 0 ; j < offset; j++ {
                    sUp += b[i - offset + j]
                    sDown += b[i+1 + offset -j]
                }
                if sUp == sDown {
                    sum += 100 * (i + 1)
                }
            }
        }
        // Check vertical mirrors
        for i := 0; i < len(b[0]) - 1; i++ {
            if getCol(b, i) == getCol(b, i+1) {
                offset := min(i, len(b[0]) - i - 2)
                sLeft := ""
                sRight := ""
                for j := 0 ; j < offset; j++ {
                    sLeft += getCol(b, i - offset + j)
                    sRight += getCol(b, i+1 + offset -j)
                }
                if sLeft == sRight {
                    sum += i + 1
                }
            }
        }
    }
    fmt.Println("----- The Answer My Friend is:", sum ," -----")
}

// Get the column from a block
func getCol(block []string, c int) string {
    s := ""
    for _, l := range block {
        s += string(l[c])
    }
    return s
}

func part2() {
}
