package main

import (
    "strings"
    "os"
    "slices"
)

func main() {
    part1()
    // part2()
}


func part1() {
    file, _ := os.ReadFile("input.txt")
    h := 0 // Number of lines
    w := 0 // Number of chars per line
    for i := len(file) - 1; i >= 0; i-- {
        if file[i] == 10 {
            // 10 is \n
            h++ // add a line
            file = append(file[:i], file[i+1:]...)
        }
    }
    w = len(file) / h
    // Location of the S character
    sLoc := 0
    for i := 0 ; i < len(file) ; i++ {
        if string(file[i]) == "S" {
            sLoc = i
        }
    }

    var validMoves = map[int]string{
        0: "|7FS", //Up
        1: "-7JS", //Right
        2: "|JLS", //Down
        3: "-FLS", //Left
    }
    var validNextMove = map[string][]int {
        "|": []int{0, 2},
        "-": []int{1, 3},
        "L": []int{0, 1},
        "J": []int{0, 3},
        "7": []int{2, 3},
        "F": []int{2, 1},
        "S": []int{0, 1, 2, 3},
    }
    // Traverse the maze
    currentLoc := -1
    lastLoc := -1
    move := 0
    loopPoints := []int{} // Used for Part 2
    for currentLoc != sLoc {
        if move == 0 {
            currentLoc = sLoc
        }
        // Check all 4 directions
        loop:
        for d := 0; d < len(validMoves); d++ {
            switch d {
            case 0:
                if currentLoc - w >= 0 &&
                    strings.Contains(validMoves[d], string(file[currentLoc - w])) &&
                    slices.Contains(validNextMove[string(file[currentLoc])], d) {
                    if currentLoc - w != lastLoc {
                        lastLoc = currentLoc
                        currentLoc = currentLoc - w
                        break loop
                    }
                }
            case 1:
                if currentLoc % w < w - 1 &&
                    strings.Contains(validMoves[d], string(file[currentLoc + 1])) &&
                    slices.Contains(validNextMove[string(file[currentLoc])], d) {
                    if currentLoc + 1 != lastLoc {
                        lastLoc = currentLoc
                        currentLoc = currentLoc + 1
                        break loop
                    }
                }
            case 2:
                if currentLoc + w < len(file) &&
                    strings.Contains(validMoves[d], string(file[currentLoc + w])) &&
                    slices.Contains(validNextMove[string(file[currentLoc])], d) {
                    if currentLoc + w != lastLoc {
                        lastLoc = currentLoc
                        currentLoc = currentLoc + w
                        break loop
                    }
                }
            case 3:
                if currentLoc % w != 0 &&
                    strings.Contains(validMoves[d], string(file[currentLoc - 1])) &&
                    slices.Contains(validNextMove[string(file[currentLoc])], d) {
                    if currentLoc - 1 != lastLoc {
                        lastLoc = currentLoc
                        currentLoc = currentLoc - 1
                        break loop
                    }
                }
            }
        }
        move++
        loopPoints = append(loopPoints, currentLoc) // For part 2
    }

    for i, _ := range file {
        if (i+1) % w == 0 {
            print("\n")
        }
        if slices.Contains(loopPoints, i) {
            print("X")
        } else {
            print(".")
        }
    }

    println()
    println("----- The Answer My Friend Is : ", (move + 1) / 2, " -----")
}

func part2() {
}
