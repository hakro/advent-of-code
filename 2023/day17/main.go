package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    part1()
    // part2()
}

type point struct {
    x int
    y int
}

var w int
var h int

func getNeighbors(p point) []point {
    n := []point{}
    if p.x > 0 {
        n = append(n, point{p.x - 1, p.y})
    }
    if p.x < w - 1 {
        n = append(n, point{p.x + 1, p.y})
    }
    if p.y > 0 {
        n = append(n, point{p.x, p.y - 1})
    }
    if p.y < h - 1 {
        n = append(n, point{p.x, p.y + 1})
    }
    return n
}

// return the heat loss value of a point from the grid
func getHeat(grid *[]string, p point) int {
    v, _ := strconv.Atoi(string((*grid)[p.y][p.x]))
    return v
}

// Get the shortest path from queue, given the current heat
func getShortest(grid *[]string, q *[]point, totalHeat int) point {
    res := point{}
    heat := 100000
    index := -1
    for i, p := range *q {
        if getHeat(grid, p) + totalHeat < heat {
            res = p
            index = i
        }
    }
    *q = append((*q)[:index], ((*q)[index+1:])...)
    return res
}

func part1() {
    file, _ := os.Open("input-example.txt")
    scanner := bufio.NewScanner(file)

    grid := []string{}
    for scanner.Scan() {
        grid = append(grid, strings.TrimSpace(scanner.Text()))
    }
    w = len(grid[0])
    h = len(grid)

    // Start with point 0,0 with a distance of 0
    start := point{0, 0}
    end := point{w - 1, h - 1}
    // current := start
    visited := map[point]int{start: 0}

    // Queue that will hold places to visit next
    q := []point{}
    q = append(q, getNeighbors(point{0,0})...)

    totalHeat := 0
    for len(q) != 0 {
        curr := getShortest(&grid, &q, totalHeat)
        fmt.Println(curr, q)
        totalHeat+= getHeat(&grid, curr)
        visited[curr] = totalHeat
        if curr == end {
            break
        }
        for _, n := range getNeighbors(curr) {
            if _, ok := visited[n]; !ok {
                q = append(q, n)
            }
        }

    }
    fmt.Println("totalHeat: ", totalHeat)
}

func part2() {
}
