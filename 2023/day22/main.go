package main

import (
	"os"
	"fmt"
	"bufio"
    "slices"
    "strconv"
	"strings"
)

func main() {
	part1()
    // test()
}

type brick struct {
    x1 int
    y1 int
    z1 int

    x2 int
    y2 int
    z2 int

    above []*brick
    below []*brick
}

func (b *brick) overlap (obj brick) bool {
    xOverlap := max(b.x1, obj.x1) <= min(b.x2, obj.x2)
    yOverlap := max(b.y1, obj.y1) <= min(b.y2, obj.y2)
    if xOverlap && yOverlap {
        return true
    }
    return false
}

func (b *brick) fallSteps (world []brick) int {
    maxSteps := b.z1 - 1 // Max fall is to z = 1
    for _, obj := range world {
        if b.overlap(obj) {
            maxSteps = b.z1 - obj.z1 - 1
        }
    }
    return maxSteps
}

func (b *brick) moveBy (steps int) {
    b.z1 -= steps
    b.z2 -= steps
}

var world []brick

func part1() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    // minZ := 0 // To remember the z of the lowest brick
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        b := brick{}
        b.x1, _ = strconv.Atoi(strings.Split(strings.Split(strings.TrimSpace(scanner.Text()), "~")[0], ",")[0])
        b.y1, _ = strconv.Atoi(strings.Split(strings.Split(strings.TrimSpace(scanner.Text()), "~")[0], ",")[1])
        b.z1, _ = strconv.Atoi(strings.Split(strings.Split(strings.TrimSpace(scanner.Text()), "~")[0], ",")[2])

        b.x2, _ = strconv.Atoi(strings.Split(strings.Split(strings.TrimSpace(scanner.Text()), "~")[1], ",")[0])
        b.y2, _ = strconv.Atoi(strings.Split(strings.Split(strings.TrimSpace(scanner.Text()), "~")[1], ",")[1])
        b.z2, _ = strconv.Atoi(strings.Split(strings.Split(strings.TrimSpace(scanner.Text()), "~")[1], ",")[2])

        added := false
        for i, obj := range world {
            if b.z1 < obj.z1 {
                world = slices.Insert(world, i, b)
                added = true
                break
            }
        }
        if len(world) == 0 || !added {
            world = append(world, b)
        }
    }

    // Here, the world is populated with bricks ordered on the Z axis
    // Apply gravity
    for i, b := range world {
        if b.z1 == 1 || b.z2 == 1 {
            // Already in place, do nothing
            continue
        }
        fallSteps := b.fallSteps(world[:i])
        world[i].moveBy(fallSteps)
    }

    removable := 0 // Final answer

    for i, b := range world {
        for j, obj := range world {
            if obj.z1 == b.z1 - 1 && b.overlap(obj){
                world[i].below = append(world[i].below, &world[j])
            }
            if obj.z1 == b.z1 + 1 && b.overlap(obj){
                world[i].above = append(world[i].above, &world[j])
            }
        }
        // fmt.Println(b, "has", len(b.above), "above and", len(b.below),"below")
    }

    // Count
    for _, b := range world {
        if len(b.above) == 0 {
            removable++
            continue
        }

        for _, above := range b.above {
            if len(above.below) > 1 {
                removable++
                break
            }
        }
    }

    fmt.Println("----- The Answer My Friend Is : ", removable, " -----")
}
