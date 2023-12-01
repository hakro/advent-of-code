package main

// Link : https://adventofcode.com/2023/day/1
// Input : https://adventofcode.com/2023/day/1/input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total_calibration_value := 0
	// Read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line_calibration_value := ""
		t := scanner.Text()
		for i := 0; i < len(t); i++ {
			if _, err := strconv.Atoi(string(t[i])); err == nil {
				line_calibration_value += string(t[i])
				break
			}
		}
		if line_calibration_value != "" {
			for i := len(t) - 1; i >= 0; i-- {
				if _, err := strconv.Atoi(string(t[i])); err == nil {
					line_calibration_value += string(t[i])
					break
				}
			}
		}

		if n, err := strconv.Atoi(line_calibration_value); err == nil {
			total_calibration_value += n
		}
	}

    fmt.Println("------ The Answer my friend, is -------")
	fmt.Println(total_calibration_value)
}
