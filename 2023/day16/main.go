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
    fmt.Println("---- The Answer My Friend Is:", len(visited), "-----")

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
