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

func part1() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

    //workflows
    wfs := map[string]string{}
    ratings := []map[string]int{}
    // Get all workflows
    for scanner.Scan() {
        if scanner.Text() == "" {
            break
        }
        wfs[strings.Split(scanner.Text(), "{")[0]] = strings.TrimSuffix(strings.Split(scanner.Text(), "{")[1], "}")
    }

    // Get all ratings
    for scanner.Scan() {
        rt := strings.Split(scanner.Text()[1:len(scanner.Text()) - 1], ",")
        m := map[string]int{}
        for _, r := range(rt) {
            k := strings.Split(r, "=")[0]
            v, _ := strconv.Atoi(strings.Split(r, "=")[1])
            m[k] = v
        }
        ratings = append(ratings, m)
    }

    // Foreach rating, go through the workflow
    // fmt.Println(wfs)
    // fmt.Println(len(ratings))
    totalAnswer := 0
    for _, rating := range ratings {
        dest := "in"
        for dest != "A" && dest != "R" {
            for _, rule := range strings.Split(wfs[dest], ",") {
                if strings.Contains(rule, ":") {
                    // dest = strings.Split(rule, ":")[0]
                    c := strings.Split(rule, ":")[0]
                    d := strings.Split(rule, ":")[1]
                    if strings.Contains(c, "<") {
                        k := strings.Split(c, "<")[0]
                        v, _ := strconv.Atoi(strings.Split(c, "<")[1])
                        if rating[k] < v {
                            dest = d
                            if dest == "A" {
                                totalAnswer += sumRating(rating)
                            }
                            break
                        }
                    }
                    if strings.Contains(c, ">") {
                        k := strings.Split(c, ">")[0]
                        v, _ := strconv.Atoi(strings.Split(c, ">")[1])
                        if rating[k] > v {
                            dest = d
                            if dest == "A" {
                                totalAnswer += sumRating(rating)
                            }
                            break
                        }
                    }
                } else {
                    dest = rule
                }
                if dest == "A" {
                    if dest == "A" {
                        totalAnswer += sumRating(rating)
                    }
                }
            }
        }
    }
    fmt.Println("----- The Answer My Friend Is : ", totalAnswer, " -----")
}

func sumRating(r map[string]int) int {
    s := 0
    for _, p := range r {
        s += p
    }
    return s
}

func part2() {
}
