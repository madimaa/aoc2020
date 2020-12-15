package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	part1()
	part2()
	os.Exit(0)
}

func part1() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("15.txt")
	occurrenceMap := make(map[int]int)
	numbers := make([]int, 0)
	for i, n := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(n)
		occurrenceMap[num] = i + 1
		numbers = append(numbers, num)
	}

	counter := len(numbers)
	lastNumber := numbers[counter-1]
	delete(occurrenceMap, lastNumber)

	for counter < 2020 {
		if occurrenceMap[lastNumber] == 0 {
			occurrenceMap[lastNumber] = counter
			lastNumber = 0
		} else {
			lastOccurrence := occurrenceMap[lastNumber]
			occurrenceMap[lastNumber] = counter
			lastNumber = counter - lastOccurrence
		}

		counter++
	}

	fmt.Println("Result: ", lastNumber)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")
	input := lib.OpenFile("15.txt")
	occurrenceMap := make(map[int]int)
	numbers := make([]int, 0)
	for i, n := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(n)
		occurrenceMap[num] = i + 1
		numbers = append(numbers, num)
	}

	counter := len(numbers)
	lastNumber := numbers[counter-1]
	delete(occurrenceMap, lastNumber)

	for counter < 30000000 {
		if occurrenceMap[lastNumber] == 0 {
			occurrenceMap[lastNumber] = counter
			lastNumber = 0
		} else {
			lastOccurrence := occurrenceMap[lastNumber]
			occurrenceMap[lastNumber] = counter
			lastNumber = counter - lastOccurrence
		}

		counter++
	}

	fmt.Println("Result: ", lastNumber)
	lib.Elapsed()
}
