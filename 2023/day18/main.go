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

    part2 := true

    for scanner.Scan() {
        dir := strings.Fields(scanner.Text())[0]
        moveCount, _ := strconv.Atoi(strings.Fields(scanner.Text())[1])

        if part2 {
            hexString := strings.Fields(scanner.Text())[2]
            hexString = hexString[1 : len(hexString) - 1]
            hexMove, _ := strconv.ParseInt(hexString[1:len(hexString) - 1], 16, 32)
            moveCount = int(hexMove)
            switch string(hexString[len(hexString) - 1]) {
            case "0":
                dir = "R"
            case "1":
                dir = "D"
            case "2":
                dir = "L"
            case "3":
                dir = "U"
            }
        }

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
