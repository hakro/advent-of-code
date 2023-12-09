package main

import (
    "os"
    "fmt"
    "bufio"
    "slices"
    "strings"
    "strconv"
)

func main() {
    part1()
    // part2()
}

func part1() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

    // This is the final result
    totalCountPart1 := 0
    totalCountPart2 := 0

    for scanner.Scan() {
        line := strings.Fields(scanner.Text())

        layers := [][]int{}
        layers = append(layers, []int{})
        // Add the first layer, which is the input
        for _, s := range line {
            i, _ := strconv.Atoi(s)
            layers[0] = append(layers[0], i)
        }
        // Add the next layers
        i := 0
        for true {
            if slices.Min(layers[i]) == 0 && slices.Max(layers[i]) == 0 {
                // All elems are 0s, so leave the loop
                break
            }
            i++
            // fill next line
            layers = append(layers, []int{})
            for j := 0; j < len(layers[i - 1]) - 1; j++ {
                layers[i] = append(layers[i], layers[i - 1][j + 1] - layers[i - 1][j])
            }
        }

        lineCountPart2 := 0
        // Walk back and find the prection for this line
        for l := len(layers) - 1; l >= 0; l-- {
            totalCountPart1 += layers[l][len(layers[l]) - 1]
            if l > 0 {
                lineCountPart2 = layers[l-1][0] - lineCountPart2
            }
        }
        totalCountPart2 += lineCountPart2
    }
    fmt.Println("----- Part 1 - The Answer My Friend Is :", totalCountPart1, " -----")
    fmt.Println("----- Part 2 - The Answer My Friend Is :", totalCountPart2, " -----")
}

func part2() {
}

