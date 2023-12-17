package main

import (
    "os"
    "fmt"
    "bufio"
    "slices"
    "strings"
)

func main() {
    part1()
    // part2()
}

const (
    up = "up"
    down = "down"
    left = "left"
    right = "right"
)
type point struct {
    x int
    y int
}

var input []string = []string{}
var visited []point = []point{}
var visitedDirs map[point][]string = map[point][]string{}

func part1() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        input = append(input, strings.TrimSpace(scanner.Text()))
    }


    visit(point{0,0}, right)
    fmt.Println("---- The Part 1 Answer My Friend Is:", len(visited), "-----")


    // Part 2
    // clearAll()
    topRes := 0
    // Check from top to bottom
    for i := 0; i < len(input[0]); i++ {
        clearAll()
        visit(point{i, 0}, down)
        topRes = max(topRes, len(visited))
    }
    // Check from bottom to top
    for i := 0; i < len(input[0]); i++ {
        clearAll()
        visit(point{i, len(input) - 1}, up)
        topRes = max(topRes, len(visited))
    }
    // Check from left to right
    for i := 0; i < len(input); i++ {
        clearAll()
        visit(point{0, i}, right)
        topRes = max(topRes, len(visited))
    }
    // Check from right to left
    for i := 0; i < len(input); i++ {
        clearAll()
        visit(point{len(input[0]) - 1, i}, left)
        topRes = max(topRes, len(visited))
    }
    fmt.Println("---- The Part 2 Answer My Friend Is:", topRes, "-----")

}

func clearAll() {
    visited = nil
    clear(visitedDirs)
}

func visit(next point, direction string) {
    if next.x < 0 || next.x >= len(input[0]) || next.y < 0 || next.y >= len(input) {
        return
    }
    if string(input[next.y][next.x]) == "." && slices.Contains(visited, next) && slices.Contains(visitedDirs[next], direction) {
        return
    }

    if !slices.Contains(visited, next) {
        visited = append(visited, next)
        if !slices.Contains(visitedDirs[next], direction) {
            visitedDirs[next] = append(visitedDirs[next], direction)
        }
    }

    switch string(input[next.y][next.x]) {
    case ".":
        switch direction {
        case up:
            visit(point{next.x, next.y - 1}, direction)
        case down:
            visit(point{next.x, next.y + 1}, direction)
        case left:
            visit(point{next.x - 1, next.y}, direction)
        case right:
            visit(point{next.x + 1, next.y}, direction)
        }
    case "|":
        switch direction {
        case left, right:
            visit(point{next.x, next.y - 1}, up)
            visit(point{next.x, next.y + 1}, down)
        case up:
            visit(point{next.x, next.y - 1}, direction)
        case down:
            visit(point{next.x, next.y + 1}, direction)
        }
    case "-":
        switch direction {
        case up, down:
            visit(point{next.x + 1, next.y}, right)
            visit(point{next.x - 1, next.y}, left)
        case right:
            visit(point{next.x + 1, next.y}, direction)
        case left:
            visit(point{next.x - 1, next.y}, direction)
        }
    case "/":
        switch direction {
        case right:
            visit(point{next.x, next.y - 1}, up)
        case left:
            visit(point{next.x, next.y + 1}, down)
        case up:
            visit(point{next.x + 1, next.y}, right)
        case down:
            visit(point{next.x - 1, next.y}, left)
        }
    case "\\":
        switch direction {
        case right:
            visit(point{next.x, next.y + 1}, down)
        case left:
            visit(point{next.x, next.y - 1}, up)
        case up:
            visit(point{next.x - 1, next.y}, left)
        case down:
            visit(point{next.x + 1, next.y}, right)
        }
    }

}

func part2() {
}
