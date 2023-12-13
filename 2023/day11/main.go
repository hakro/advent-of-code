package main

import (
    "os"
    "bufio"
    "strings"
)

func main() {
    part1()
    // part2()
}


func part1() {
    file, _ := os.Open("input.txt")

    input := []string{}
    h := 0 // number of lines
    w := 0

    part2 := true
    scanner := bufio.NewScanner(file)
    // foreach line
    for scanner.Scan() {
        h++
        l := strings.TrimSuffix(scanner.Text(), "\n")
        input = append(input, l)
        w = len(l)

        // Insert horizonal double lines
        if !strings.Contains(l, "#") {
            h++
            input = append(input, l)
        }
    }

    doublePos := []int{}
    for i := 0; i < w; i++ {
        c := 0 // count of points
        for j := 0; j < h; j++ {
            if string(input[j][i]) == "." {
                c++
            }
        }
        if c == h {
            doublePos = append(doublePos, i)
        }
    }

    // Insert vertical double lines
    for i := len(doublePos) -1; i >= 0; i-- {
        for j, _ := range input {
            input[j] = input[j][:doublePos[i]] + "." + input[j][doublePos[i]:]
        }
    }

    type star struct {
        x int
        y int
    }
    stars := map[int]star{}
    c := 0
    for i, _ := range input {
        for j, _ := range input[i] {
            if string(input[i][j]) == "#" {
                c++
                stars[c] = star{i, j}
            }
        }
    }

    // This is the final answer
    shortPathSum := 0

    for i := 1; i <= len(stars); i++ {
        for j := i+1; j < len(stars) + 1; j++ {
            shortPathSum += abs(stars[j].x - stars[i].x) + abs(stars[j].y - stars[i].y)
        }
    }


    println("----- The Answer My Friend Is : ", shortPathSum, " -----")
}

func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}

func part2() {
}
