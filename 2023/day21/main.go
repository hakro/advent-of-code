package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	// part2()
}

type Grid []string

type Point struct {
    x int
    y int
}

func countO(g Grid) int {
    count := 0
    for _, r := range g {
        for _, c := range r {
            if string(c) == "O" {
                count++
            }
        }
    }
    return count
}

func nextVisit(g Grid, p Point) {
    g[p.y] = g[p.y][:p.x] + "." + g[p.y][p.x + 1:]
    // North
    if p.y > 0 && string(g[p.y - 1][p.x]) != "#" {
        g[p.y - 1] = g[p.y - 1][:p.x] + "O" + g[p.y - 1][p.x + 1:]
    }
    // South
    if p.y < len(g) - 1 && string(g[p.y + 1][p.x]) != "#" {
        g[p.y + 1] = g[p.y + 1][:p.x] + "O" + g[p.y + 1][p.x + 1:]
    }
    // East
    if p.x < len(g[0]) - 1 && string(g[p.y][p.x + 1]) != "#" {
        g[p.y] = g[p.y][:p.x + 1] + "O" + g[p.y][p.x + 2:]
    }
    // West
    if p.x > 0 && string(g[p.y][p.x - 1]) != "#" {
        g[p.y] = g[p.y][:p.x - 1] + "O" + g[p.y][p.x:]
    }
}

func getVisited(g Grid) []Point {
    pts := []Point{}
    for j, r := range g {
        for i, c := range r {
            if string(c) == "O" {
                pts = append(pts, Point{i, j})
            }
        }
    }
    return pts
}

func part1() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

    start := Point{}
    grid := Grid{}
    for scanner.Scan() {
        t := strings.TrimSpace(scanner.Text())
        if strings.Contains(t, "S") {
            start.y = len(grid)
            start.x = strings.Index(t, "S")
        }
        grid = append(grid, t)
    }

    steps := 64

    nextVisit(grid, start)
    for i := 0; i < steps - 1; i++ {
        // visited := getVisited(grid)
        for _, v := range getVisited(grid) {
            nextVisit(grid, v)
        }
    }

    fmt.Println("----- The Answer My Friend Is : ", countO(grid), "-----")
}

func show(g Grid) {
    fmt.Println("XXXXXXXXXXXXXXXXXXX")
    for _, l := range g {
        fmt.Println(l)
    }
}

func part2() {
}
