package main

import (
    "os"
    "fmt"
    "strings"
)

func main() {
    part1()
    // part2()
}

func hash(s string) int {
    current := 0
    for _, c := range strings.TrimSpace(s) {
        current = ((current + int(c)) * 17) % 256
    }
    return current
}

func part1() {
    file, _ := os.ReadFile("input.txt")
    values := strings.Split(string(file), ",")

    total := 0
    for _, v := range values {
        total += hash(v)
    }

    fmt.Println("----- The Part 1 Answer My Friend Is :", total, "-----")
}

func part2() {
}
