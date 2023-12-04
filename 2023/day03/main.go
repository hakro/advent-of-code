package main

import (
	"fmt"
	"os"
    "strings"
    "regexp"
    "unicode"
    "strconv"
)

func main() {
    // part1()
    part2()
}

func part1() {
    // This is the final answer :
    totalCount := 0
    file, _ := os.ReadFile("input.txt")
    pat := regexp.MustCompile(`\d+`)
    lines := strings.Split(strings.TrimSpace(string(file)), "\n")
    // Foreach line in the file
    for i, l := range lines {
        nums := pat.FindAllString(l, -1)
        positions := pat.FindAllStringIndex(l, -1)
        // Foreach number matched
        for j, n := range nums {
            firstIndex := positions[j][0]
            lastIndex := firstIndex + len(n) - 1
            // Check the left side
            if firstIndex > 0 {
                // Check direct left
                if isSymbol(l[firstIndex - 1]) {
                    val, _ := strconv.Atoi(n)
                    totalCount += val
                    continue
                }
                // Check upper left diagonally
                if i > 0 {
                    if isSymbol(lines[i-1][firstIndex - 1]) {
                        val, _ := strconv.Atoi(n)
                        totalCount += val
                        continue
                    }
                }
                // Check lower left diagonally
                if i < len(lines) - 1 {
                    if isSymbol(lines[i+1][firstIndex - 1]) {
                        val, _ := strconv.Atoi(n)
                        totalCount += val
                        continue
                    }
                }
            }
            // Check the right side
            if lastIndex < len(l) - 1 {
                // Check direct right
                if isSymbol(l[lastIndex + 1]) {
                    val, _ := strconv.Atoi(n)
                    totalCount += val
                    continue
                }
                // Check upper right diagonally
                if i > 0 {
                    if isSymbol(lines[i-1][lastIndex + 1]) {
                        val, _ := strconv.Atoi(n)
                        totalCount += val
                        continue
                    }
                }
                // Check lower right diagonally
                if i < len(lines) - 2 {
                    if isSymbol(lines[i+1][lastIndex + 1]) {
                        val, _ := strconv.Atoi(n)
                        totalCount += val
                        continue
                    }
                }
            }
            // Check the upper and lower sides
            for u := firstIndex; u <= lastIndex; u++ {
                if i > 0 {
                    if isSymbol(lines[i-1][u]) {
                        val, _ := strconv.Atoi(n)
                        totalCount += val
                        break
                    }
                }
                if i < len(lines) - 1 {
                    if isSymbol(lines[i+1][u]) {
                        val, _ := strconv.Atoi(n)
                        totalCount += val
                        break
                    }
                }
            }
        }
    }

    fmt.Println("----- The answer my friend, is : ", totalCount, " -----")
}

func part2() {
    starMap := make(map[string][]int) //Keys are formatted as line-row as star coordinates, values will be the surrounding numbers
    file, _ := os.ReadFile("input.txt")
    pat := regexp.MustCompile(`\d+`)
    lines := strings.Split(strings.TrimSpace(string(file)), "\n")
    // Foreach line in the file
    for i, l := range lines {
        nums := pat.FindAllString(l, -1)
        positions := pat.FindAllStringIndex(l, -1)
        // Foreach number matched
        for j, n := range nums {
            firstIndex := positions[j][0]
            lastIndex := firstIndex + len(n) - 1
            // Check the left side
            if firstIndex > 0 {
                // Check direct left
                if isSymbolStar(l[firstIndex - 1]) {
                    val, _ := strconv.Atoi(n)
                    starMap[string(i) + "-" + string(firstIndex - 1)] = append(starMap[string(i) + "-" + string(firstIndex - 1)], val)
                    continue
                }
                // Check upper left diagonally
                if i > 0 {
                    if isSymbolStar(lines[i-1][firstIndex - 1]) {
                        val, _ := strconv.Atoi(n)
                        starMap[string(i-1) + "-" + string(firstIndex - 1)] = append(starMap[string(i-1) + "-" + string(firstIndex - 1)], val)
                        continue
                    }
                }
                // Check lower left diagonally
                if i < len(lines) - 1 {
                    if isSymbolStar(lines[i+1][firstIndex - 1]) {
                        val, _ := strconv.Atoi(n)
                        starMap[string(i+1) + "-" + string(firstIndex - 1)] = append(starMap[string(i+1) + "-" + string(firstIndex - 1)], val)
                        continue
                    }
                }
            }
            // Check the right side
            if lastIndex < len(l) - 1 {
                // Check direct right
                if isSymbolStar(l[lastIndex + 1]) {
                    val, _ := strconv.Atoi(n)
                    starMap[string(i) + "-" + string(lastIndex + 1)] = append(starMap[string(i) + "-" + string(lastIndex + 1)], val)
                    continue
                }
                // Check upper right diagonally
                if i > 0 {
                    if isSymbolStar(lines[i-1][lastIndex + 1]) {
                        val, _ := strconv.Atoi(n)
                        starMap[string(i-1) + "-" + string(lastIndex + 1)] = append(starMap[string(i-1) + "-" + string(lastIndex + 1)], val)
                        continue
                    }
                }
                // Check lower right diagonally
                if i < len(lines) - 2 {
                    if isSymbolStar(lines[i+1][lastIndex + 1]) {
                        val, _ := strconv.Atoi(n)
                        starMap[string(i+1) + "-" + string(lastIndex + 1)] = append(starMap[string(i+1) + "-" + string(lastIndex + 1)], val)
                        continue
                    }
                }
            }
            // Check the upper and lower sides
            for u := firstIndex; u <= lastIndex; u++ {
                if i > 0 {
                    if isSymbolStar(lines[i-1][u]) {
                        val, _ := strconv.Atoi(n)
                        starMap[string(i-1) + "-" + string(u)] = append(starMap[string(i-1) + "-" + string(u)], val)
                        break
                    }
                }
                if i < len(lines) - 1 {
                    if isSymbolStar(lines[i+1][u]) {
                        val, _ := strconv.Atoi(n)
                        starMap[string(i+1) + "-" + string(u)] = append(starMap[string(i+1) + "-" + string(u)], val)
                        break
                    }
                }
            }
        }
    }

    totalCount := 0
    for _, v := range(starMap) {
        if len(v) == 2 {
            totalCount += v[0] * v[1]
        }
    }
    fmt.Println("----- The answer my friend, is : ", totalCount, " -----")
}

func isSymbol(b byte) bool {
    if !unicode.IsDigit(rune(b)) && string(b) != "." {
        return true
    }
    return false
}

func isSymbolStar(b byte) bool {
    if string(b) == "*" {
        return true
    }
    return false
}
