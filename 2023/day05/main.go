package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
    "strconv"
    "slices"
)

func main() {
    part1()
    // part2()
}

func part1() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // Get first line containing seeds
    scanner.Scan()
    seeds := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")

    type Rule struct {
        Dest int
        Src int
        Offset int
    }
    mapNames := []string{}
    mapRules := make(map[string][]Rule)
    currentMapName := ""
    for scanner.Scan() {
        if strings.TrimSpace(scanner.Text()) == "" {
            continue
        }
        if strings.HasSuffix(scanner.Text(), " map:") {
            mapNames = append(mapNames, scanner.Text())
            currentMapName = scanner.Text()
            continue
        }
        // This is a rule then
        dest, _ := strconv.Atoi(strings.Fields(scanner.Text())[0])
        src, _ := strconv.Atoi(strings.Fields(scanner.Text())[1])
        offset, _ := strconv.Atoi(strings.Fields(scanner.Text())[2])
        r := Rule{
            dest,
            src,
            offset,
        }
        mapRules[currentMapName] = append(mapRules[currentMapName], r)
    }
    // For all seeds, traverse all rules
    locations := []int{}
    for _, seedString := range seeds {
        seed, _ := strconv.Atoi(seedString)
        for _, mn := range(mapNames) {
            for _, rule := range(mapRules[mn]) {
                if seed >= rule.Src && seed < rule.Src + rule.Offset {
                    // seed = rule.Dest - rule.Src + seed
                    seed = seed - rule.Src + rule.Dest
                    break
                }
            }
        }
        locations = append(locations, seed)
    }
    fmt.Println("------- The answer my friend is ", slices.Min(locations), "-------")
}

func part2() {
}

