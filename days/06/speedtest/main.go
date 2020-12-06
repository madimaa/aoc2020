package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("../06.txt")
	yesString := ""
	numberOfYeses := 0
	totalNumberOfYeses := 0
	for _, row := range input {
		if len(row) == 0 {
			totalNumberOfYeses += numberOfYeses
			yesString = ""
			numberOfYeses = 0
		}

		for _, currentRune := range []rune(row) {
			if !strings.ContainsRune(yesString, currentRune) {
				yesString += string(currentRune)
				numberOfYeses++
			}
		}
	}

	totalNumberOfYeses += numberOfYeses

	fmt.Println("Result: ", totalNumberOfYeses)
	lib.Elapsed()
	os.Exit(0)
}
