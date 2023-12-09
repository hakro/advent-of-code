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
    totalCount := 0

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

        // Walk back and find the prection for this line
        for l := len(layers) - 1; l >= 0; l-- {
            totalCount += layers[l][len(layers[l]) - 1]
        }
    }
    fmt.Println("----- The Answer My Friend Is :", totalCount, " -----")
}

func part2() {
}

