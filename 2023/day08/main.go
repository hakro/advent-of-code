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

    scanner.Scan() //Skip empty line
    elems := make(map[string][2]string)
    for scanner.Scan() {
        l := strings.Split(scanner.Text(), " = ")
        elems[l[0]] = [2]string{
            strings.TrimPrefix(strings.Split(l[1], ", ")[0], "("),
            strings.TrimSuffix(strings.Split(l[1], ", ")[1], ")"),
        }
    }
    for true {
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

