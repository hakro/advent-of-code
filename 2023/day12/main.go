package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    o := []int{1,1,3}
    check("???.###", o)
    println(c)
    // part1()
    // part2()
}

var c int = 0
func check(s string, occ []int) {
    fmt.Println("--- s: ", s, " --oc:", occ)
    if len(s) == 0 || len(occ) == 0 {
        c++
        return
    }

    if strings.HasPrefix(s, ".") {
        check(s[1:], occ)
    }

    if strings.HasPrefix(s, "#") {
        if strings.HasPrefix(s, strings.Repeat("#", occ[0])) {
            check(s[occ[0]:], occ[1:])
        }
    }

    if strings.HasPrefix(s, "?") {
        check(strings.Replace(s, "?", "#", 1), occ)
        check(strings.Replace(s, "?", ".", 1), occ)
    }
}


func part1() {
    file, _ := os.Open("input-example.txt")
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        springs := strings.Fields(scanner.Text())[0]
        //occurences : The second part of each line
        occ := []int{}
        for _, o := range strings.Split(strings.Fields(scanner.Text())[1], ",") {
            n, _ := strconv.Atoi(o)
            occ = append(occ, n)
        }
        println(springs)
    }
}

func part2() {
}
