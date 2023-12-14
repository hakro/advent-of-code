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

    input := []string{}
    for scanner.Scan() {
        input = append(input, strings.TrimSpace(scanner.Text()))
        if len(input) == 1 {
            continue
        }
        for i, c := range(scanner.Text()) {
            if string(c) == "O" {
                // Check up north
                switchLoc := -1
                for l := len(input) - 2; l >=0; l-- {
                    if string(input[l][i]) != "." {
                        // Can't go north
                        break
                    }
                    switchLoc = l
                }
                //Swap
                if switchLoc != -1 {
                    input[len(input) - 1] = input[len(input) - 1][:i] + string(".") + input[len(input) - 1][i+1:]
                    input[switchLoc] = input[switchLoc][:i] + string("O") + input[switchLoc][i+1:]
                }
            }
        }
    }
    // This is the final Answer
    sum := 0
    for i, l := range(input) {
        sum += strings.Count(l, "O") * (len(input) - i)
    }

    fmt.Println("----- The Answer My Friend Is : ", sum, " -----")
}

func part2() {
}
