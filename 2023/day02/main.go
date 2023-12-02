package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalCubes := make(map[string]int)
	totalCubes["red"] = 12
	totalCubes["green"] = 13
	totalCubes["blue"] = 14

	// This is where the final result is stored
	gameCount := 0

	scanner := bufio.NewScanner(file)
	// Foreach game (line in the file)
	for scanner.Scan() {
		gameIsPossible := true
		gameInfo := scanner.Text()
		gameId, _ := strconv.Atoi(strings.Split(strings.Split(gameInfo, ": ")[0], " ")[1])
		setsInfo := strings.Split(gameInfo, ": ")[1]
		// Foreach Set of the game
		for _, v := range strings.Split(setsInfo, "; ") {
			// Foreach color and its occurence inside the set
			for _, col := range strings.Split(v, ", ") {
				nb, _ := strconv.Atoi(strings.Split(col, " ")[0])
				colName := strings.Split(col, " ")[1]
				if nb > totalCubes[colName] {
					gameIsPossible = false
					break
				}
			}
			if !gameIsPossible {
				break
			}
		}
		if gameIsPossible {
			gameCount += gameId
		}
	}

	fmt.Println("------ The Answer My Friend is: ", gameCount, " -------")
}
