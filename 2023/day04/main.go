package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
    "slices"
)

func main() {
    part1()
    // part2()
}

func part1() {
    // This is the final answer
    totalScore := 0
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)
    // Foreach card (line in the input file)
    for scanner.Scan() {
        l := scanner.Text()
        cardNums := strings.Fields(strings.TrimSpace(strings.Split(strings.Split(l, ": ")[1], "|")[0]))
        yourNums := strings.Fields(strings.TrimSpace(strings.Split(strings.Split(l, ": ")[1], "|")[1]))
        cardScore := 0
        for _, yn := range yourNums {
            if slices.Contains(cardNums, yn) {
                if cardScore == 0 {
                    cardScore++
                    continue
                }
                cardScore *= 2
            }
        }
        totalScore += cardScore
    }

    fmt.Println("------ The answer my friend is : ", totalScore, " ------")
}

func part2() {
}

