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
    file, _ := os.Open("input-example.txt")
    scanner := bufio.NewScanner(file)

    // This the final result
    sum := 0
    sum2 := 0
    part2 := true
    blocks := make([][]string, 2)

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
                if part2 {
                    smudgeCount := 0
                    smudgeIndex := -1
                    for k := 0; k < len(sUp); k++ {
                        if sUp[k] != sDown[k] {
                            smudgeCount++
                            smudgeIndex = k
                        }
                    }
                    if smudgeCount == 1 {
                        sUp = sUp[:smudgeIndex] + flipSmudge(string(sUp[smudgeIndex])) + sUp[smudgeIndex+1:]
                        sum2 += 100 * (i + 1)
                    }
                } else {
                    // Part 1
                    if sUp == sDown {
                        sum += 100 * (i + 1)
                    }
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
                if part2 {
                    smudgeCount := 0
                    smudgeIndex := -1
                    for k := 0; k < len(sLeft); k++ {
                        if sLeft[k] != sRight[k] {
                            smudgeCount++
                            smudgeIndex = k
                        }
                    }
                    fmt.Println("count: ", smudgeCount)
                    if smudgeCount == 1 {
                        sLeft = sLeft[:smudgeIndex] + flipSmudge(string(sLeft[smudgeIndex])) + sLeft[smudgeIndex+1:]
                        sum2 += i + 1
                    }
                } else {
                    // Part1
                    if sLeft == sRight {
                        sum += i + 1
                    }
                }
            }
        }
    }
    fmt.Println("----- The Answer My Friend is:", sum ," -----")
    fmt.Println("----- The Answer My Friend for PART2 is:", sum2 ," -----")
}

func flipSmudge(s string) string {
    if s == "#" {
        return "."
    }
    return "#"
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
