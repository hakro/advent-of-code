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

    // Turn this to true for part 2
    // Very slow, but works !!!
    part2 := true

    if (!part2) {
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
    } else {
        //Part 2
        for i := 0; i < len(seeds); i += 2 {
            // from, _ := strconv.ParseInt(seeds[i], 10, 64)
            // offset, _ := strconv.ParseInt(seeds[i+1], 10, 64)
            from, _ := strconv.Atoi(seeds[i])
            offset, _ := strconv.Atoi(seeds[i+1])
            for j := from; j < from + offset; j++ {
                seed := j
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
        }
    }
    fmt.Println("------- The answer my friend is ", slices.Min(locations), "-------")
}

