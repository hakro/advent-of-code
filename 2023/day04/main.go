package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
    "slices"
)

func main() {
    // part1()
    part2()
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
    file, _ := os.ReadFile("input.txt")
    lines := strings.Split(strings.TrimSpace(string(file)), "\n")

    cardCount := make(map[int]int)
    // Foreach card (line in the input file)
    for i, l := range lines {
        if _, exists := cardCount[i]; !exists {
            cardCount[i] = 1
        } else {
            cardCount[i]++
        }
        cardNums := strings.Fields(strings.TrimSpace(strings.Split(strings.Split(l, ": ")[1], "|")[0]))
        yourNums := strings.Fields(strings.TrimSpace(strings.Split(strings.Split(l, ": ")[1], "|")[1]))
        cardScore := 0
        for _, yn := range yourNums {
            if slices.Contains(cardNums, yn) {
                cardScore++
            }
        }
        for n := 0; n < cardCount[i]; n++ {
            for j := 0; j < cardScore; j++ {
                // Out of bounds
                if i + j + 1 > len(lines) - 1 {
                    break
                }
                if _, exists := cardCount[i+j+1]; !exists {
                    cardCount[i+j+1] = 1
                } else {
                    cardCount[i+j+1]++
                }
            }
        }
    }
    totalScore := 0
    for _, v := range cardCount {
        totalScore += v
    }
    fmt.Println("------ The answer my friend is : ", totalScore, " ------")
}

