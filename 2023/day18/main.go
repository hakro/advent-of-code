package main

import (
    "os"
    "fmt"
    "math"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    part1()
    // part2()
}

type vertex struct {
    x int
    y int
}

func part1() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

    area := []vertex{vertex{0, 0}}
    bounds := 0

    for scanner.Scan() {
        dir := strings.Fields(scanner.Text())[0]
        moveCount, _ := strconv.Atoi(strings.Fields(scanner.Text())[1])
        bounds += moveCount
        lastPt := area[len(area) - 1]
        var newPt vertex
        switch dir {
        case "R":
            newPt = vertex{lastPt.x + moveCount, lastPt.y}
        case "L":
            newPt = vertex{lastPt.x - moveCount, lastPt.y}
        case "D":
            newPt = vertex{lastPt.x, lastPt.y + moveCount}
        case "U":
            newPt = vertex{lastPt.x, lastPt.y - moveCount}
        }
        area = append(area, newPt)
    }

    // After determining the internal area with shoelace,
    // Use Pick's theorem to determine the area + boundaries
    // i = Area(shoelace) - b/2 + 1
    // We're looking for internal + boundaries, so :
    // i + b = Area + b/2 + 1
    fmt.Println(int(shoelace(area)))
    fmt.Println(bounds / 2)
    res := int(shoelace(area)) + (bounds / 2) + 1
    fmt.Println("----- The Answer My Friend Is : ", res, " -----")
}

// https://en.wikipedia.org/wiki/Shoelace_formula
func shoelace(area []vertex) float64 {
    sum := 0
    for i := 0; i < len(area) - 1; i++ {
        sum += area[i].x * area[i + 1].y - area[i + 1].x * area[i].y
    }
    return math.Abs(float64(sum) / 2.0)
}

func part2() {
}
