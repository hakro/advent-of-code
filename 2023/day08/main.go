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

var directions string
func getDirectionByStep(s int) string {
    return string(directions[s % len(directions)])
}

func part1() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

    scanner.Scan()
    directions = scanner.Text()
    steps := 0
    currentDir := "AAA"

    currentDirPart2 := []string{}
    scanner.Scan() //Skip empty line
    elems := make(map[string][2]string)
    for scanner.Scan() {
        l := strings.Split(scanner.Text(), " = ")
        if strings.HasSuffix(l[0], "A") {
            currentDirPart2 = append(currentDirPart2, l[0])
        }
        elems[l[0]] = [2]string{
            strings.TrimPrefix(strings.Split(l[1], ", ")[0], "("),
            strings.TrimSuffix(strings.Split(l[1], ", ")[1], ")"),
        }
    }
    part2 := true
    for true {
        if part2 {
            // Part2
            nbZ := 0 // Number of things ending in Z
            for _, e := range currentDirPart2 {
                if strings.HasSuffix(e, "Z") {
                    nbZ++
                }
            }
            fmt.Println(steps)
            // All of directions en in Z, terminate
            if nbZ == len(currentDirPart2) {
                break
            }
            for i, el := range currentDirPart2 {
                switch getDirectionByStep(steps) {
                case "L" :
                    currentDirPart2[i] = elems[el][0]
                case "R" :
                    currentDirPart2[i] = elems[el][1]
                }
            }
            steps++
            continue
        }

        // Part 1
        if currentDir == "ZZZ" {
            break
        }
        switch getDirectionByStep(steps) {
        case "L" :
            currentDir = elems[currentDir][0]
        case "R" :
            currentDir = elems[currentDir][1]
        }
        steps++
    }

    fmt.Println("------ The Answer My Friend Is : ", steps, " -------")
}

func part2() {
}

