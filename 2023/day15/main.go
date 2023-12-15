package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    part1()
    // part2()
}

func hash(s string) int {
    current := 0
    for _, c := range s {
        current = ((current + int(c)) * 17) % 256
    }
    return current
}

func part1() {
    file, _ := os.ReadFile("input.txt")
    values := strings.Split(strings.TrimSpace(string(file)), ",")

    // Used for part 2
    boxes := map[int][]string{}

    totalPart1 := 0
    for _, v := range values {
        h := hash(v)
        totalPart1 += h

        h = hash(strings.Split(strings.Split(v, "=")[0], "-")[0])
        // Part2
        if strings.HasSuffix(v, "-") {
            // for each lens in the box
            for i, l := range boxes[h] {
                if strings.Split(v, "-")[0] == strings.Split(l, "=")[0] {
                    // remove
                    boxes[h] = append(boxes[h][:i], boxes[h][i+1:]...)
                }
            }
        }

        if len(strings.Split(v, "=")) == 2 {
            added := false
            for i, l := range boxes[h] {
                if strings.Split(v, "=")[0] == strings.Split(l, "=")[0] {
                    added = true
                    // boxes[h] = slices.Insert(boxes[h], i, v)
                    boxes[h][i] = v
                    break
                }
            }
            if !added {
                    boxes[h] = append(boxes[h], v)
            }
        }
    }

    fmt.Println("----- The Part 1 Answer My Friend Is :", totalPart1, "-----")
    totalPart2 := 0
    for i, b := range boxes {
        for j, l := range b {
            // Focus Power
            focalLength, _ := strconv.Atoi(strings.Split(l, "=")[1])
            totalPart2 += (i + 1) * (j + 1) * focalLength
        }
    }
    fmt.Println("----- The Part 2 Answer My Friend Is :", totalPart2, "-----")
}

func part2() {
}
