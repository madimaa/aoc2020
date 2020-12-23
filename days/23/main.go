package main

import (
	"fmt"
	"os"
	"strconv"

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

	input := lib.OpenFileAsString("23.txt")
	circle := parseInput(input)
	circle = play(circle, 100, 1, 9)

	index := 1
	result := ""
	for len(result) != len(input)-1 {
		result += strconv.Itoa(circle[index])
		index = circle[index]
	}

	fmt.Println("Result: ", result)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFileAsString("23.txt")
	circle := parseInput(input)
	var last int
	for k, v := range circle {
		if v == circle[0] && k != 0 {
			last = k
			break
		}
	}

	next := 10
	for len(circle) != 1000001 {
		circle[last] = next
		last = next
		next++
	}

	circle[last-1] = circle[0]

	circle = play(circle, 10000000, 1, 1000000)
	cup1 := circle[1]
	cup2 := circle[cup1]

	fmt.Println("Result: ", cup1*cup2)
	lib.Elapsed()
}

func parseInput(input string) map[int]int {
	cupMap := make(map[int]int)
	runes := []rune(input)
	for i, r := range runes {
		num, err := strconv.Atoi(string(r))
		lib.PanicOnError(err)

		if i == 0 {
			cupMap[i] = num
		}

		if i+1 != len(runes) {
			next, err := strconv.Atoi(string(runes[i+1]))
			lib.PanicOnError(err)
			cupMap[num] = next
		} else {
			cupMap[num] = cupMap[0]
		}
	}

	return cupMap
}

func play(circle map[int]int, moves, min, max int) map[int]int {
	for moves > 0 {
		current := circle[0]
		num1 := circle[current]
		num2 := circle[num1]
		num3 := circle[num2]
		circle[current] = circle[num3]

		destination := current - 1
		for destination == 0 || destination == num1 || destination == num2 || destination == num3 {
			destination--
			if destination < min {
				destination = max
			}
		}

		circle[num3] = circle[destination]
		circle[destination] = num1

		circle[0] = circle[current]
		moves--
	}

	return circle
}
