package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("--- This is the play.go file ---")

    s := "teststringoneight5fivetwone"
	table := make(map[string]string)
	table["zero"] = "0"
	table["one"] = "1"
	table["two"] = "2"
	table["three"] = "3"
	table["four"] = "4"
	table["five"] = "5"
	table["six"] = "6"
	table["seven"] = "7"
	table["eight"] = "8"
	table["nine"] = "9"

    // Change the words to digits in order
    tryAgain := true
    for tryAgain {
        tryAgain = false
        for i := range s {
            foundWord := false
            for v := range table {
                if s[i: min(i + len(v), len(s))] == v {
                    foundWord = true
                    s = strings.Replace(s, v, table[v], 1)
                }
            }
            if foundWord {
                tryAgain = true
                break
            }
        }
    }
    return s
}
