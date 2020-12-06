package main

import (
	"fmt"
	"os"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("06.txt")
	abc := make(map[rune]bool)
	numberOfYeses := 0
	totalNumberOfYeses := 0
	for _, row := range input {
		if len(row) == 0 {
			totalNumberOfYeses += numberOfYeses
			abc = make(map[rune]bool)
			numberOfYeses = 0
		}

		for _, currentRune := range []rune(row) {
			if !abc[currentRune] {
				abc[currentRune] = true
				numberOfYeses++
			}
		}
	}

	totalNumberOfYeses += numberOfYeses

	fmt.Println("Result: ", totalNumberOfYeses)
	lib.Elapsed()

	lib.Start()
	fmt.Println("Part 2")
	abcInt := make(map[rune]int)
	totalNumberOfYeses = 0
	groupSize := 0
	for _, row := range input {
		if len(row) == 0 {
			totalNumberOfYeses += countYeses(groupSize, abcInt)
			groupSize = 0
			abcInt = make(map[rune]int)
		} else {
			for _, currentRune := range []rune(row) {
				abcInt[currentRune]++
			}

			groupSize++
		}
	}

	totalNumberOfYeses += countYeses(groupSize, abcInt)

	fmt.Println("Result: ", totalNumberOfYeses)

	lib.Elapsed()
	os.Exit(0)
}

func countYeses(groupSize int, abcInt map[rune]int) int {
	numberOfYeses := 0
	for _, numberOfCurrentRune := range abcInt {
		if numberOfCurrentRune == groupSize {
			numberOfYeses++
		}
	}

	return numberOfYeses
}
