package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    part1()
    // part2()
}

func part1() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    times := strings.Fields(scanner.Text())[1:]

    scanner.Scan()
    scanner.Text()
    distances := strings.Fields(scanner.Text())[1:]

    // Final answer
    answer := 1

    // For every race
    for i := 0; i < len(times); i++ {
        raceAnswer := 0
        raceTime, _ := strconv.Atoi(times[i])
        raceRecord, _ := strconv.Atoi(distances[i])
        for t:= 0; t < raceTime; t++ {
            d := (raceTime - t) * t
            if d > raceRecord {
                raceAnswer++
            }
        }
        answer *= raceAnswer
    }
    fmt.Println("------ The Answer My Friend is: ", answer, " --------")
}

func part2() {
}

