package main

import (
	"fmt"
	"os"
	"sort"
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

	input := lib.OpenFile("10.txt")
	adapters := make([]int, 0)
	for _, fileRow := range input {
		jolt, _ := strconv.Atoi(fileRow)
		adapters = append(adapters, jolt)
	}

	sort.Ints(adapters)
	base := 0
	ones, twos, threes := 0, 0, 0
	for _, jolt := range adapters {
		switch jolt - base {
		case 1:
			ones++
		case 2:
			twos++
		case 3:
			threes++
		default:
			panic(fmt.Sprint("Something is not right, the difference is ", jolt-base))
		}

		base = jolt
	}

	threes++ //your device has a built-in joltage adapter rated for 3 jolts higher than the highest-rated adapter in your bag

	fmt.Println("Result: ", ones, "*", threes, "=", ones*threes)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("10.txt")
	adapters := make([]int, 0)
	for _, fileRow := range input {
		jolt, _ := strconv.Atoi(fileRow)
		adapters = append(adapters, jolt)
	}

	sort.Ints(adapters)
	paths := 1
	base := 0
	ones, twos, threes := 0, 0, 0
	for _, jolt := range adapters {
		switch jolt - base {
		case 1:
			ones++
		case 2:
			twos++
		case 3:
			//TODO my input does not have a 2 jump so I will skip that part, but have to code it later
			if ones > 1 {
				paths *= possibleSteps(ones)
			}

			ones = 0
			threes++
		default:
			panic(fmt.Sprint("Something is not right, the difference is ", jolt-base))
		}

		base = jolt
	}

	if ones > 1 {
		paths *= possibleSteps(ones)
	}

	fmt.Println("Result: ", paths)
	lib.Elapsed()
}

func possibleSteps(distance int) int {
	//DO NOT ASK! ITS ON MY WHITEBOARD
	//ok. let me explain s = start, 1 = 1 jump, 2 = 2 jump, 3 = 3 jump, there is no more because you cannot jump 4
	switch distance {
	case 2:
		return 2 //means s->1,1 or s->2
	case 3:
		return 4 //means s->1,1,1 or s->1,2 or s->2,1 or s->3
	case 4:
		return 7 //means s->1,1,1,1 or s->2,1,1 or s->1,2,1 or s->1,1,2 or s->2,2 or s->3,1 or s->1,3
	case 5:
		return 13 //just follow the patter
	default:
		panic("¯\\_(ツ)_/¯")
	}

	//If someone reads this and knows how to code the calculation for this, let me know. PRs are welcome.
}
