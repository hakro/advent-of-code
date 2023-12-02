package main

// Link : https://adventofcode.com/2023/day/1
// Input : https://adventofcode.com/2023/day/1/input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Turn this on for Day 1's second Part
var part2 bool = true

func main() {

	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalCalibrationValue := 0
	// Read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCalibrationValue := ""
		t := scanner.Text()
        //
        fmt.Println("-----")
        fmt.Println(t)
        if (part2) {
            t = convertToDigits(t)
        }
        fmt.Println(t)

		for i := 0; i < len(t); i++ {
			if _, err := strconv.Atoi(string(t[i])); err == nil {
				lineCalibrationValue += string(t[i])
				break
			}
		}
		if lineCalibrationValue != "" {
			for i := len(t) - 1; i >= 0; i-- {
				if _, err := strconv.Atoi(string(t[i])); err == nil {
					lineCalibrationValue += string(t[i])
					break
				}
			}
		}
        fmt.Println("line calib: " + lineCalibrationValue)
		if n, err := strconv.Atoi(lineCalibrationValue); err == nil {
			totalCalibrationValue += n
		}
	}

	fmt.Println("------ The Answer my friend, is -------")
	fmt.Println(totalCalibrationValue)
}

// This is for part 2 of day 1, where written numbers need to be converted to digits
func convertToDigits(s string) string {
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
