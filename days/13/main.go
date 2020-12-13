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

	input := lib.OpenFile("13.txt")
	targetTimestamp, _ := strconv.Atoi(input[0])
	input[1] = strings.ReplaceAll(input[1], "x,", "")
	busNum, minWait := 0, 999
	for _, in := range strings.Split(input[1], ",") {
		num, _ := strconv.Atoi(in)
		waitTime := num - (targetTimestamp % num)
		if waitTime < minWait {
			minWait = waitTime
			busNum = num
		}
	}

	fmt.Println("Result: ", minWait*busNum)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("13.txt")
	busNumbers := make([]int, 0)
	reminders := make([]int, 0)

	for i, in := range strings.Split(input[1], ",") {
		if in == "x" {
			continue
		}

		number, _ := strconv.Atoi(in)
		busNumbers = append(busNumbers, number)
		reminders = append(reminders, number-i)
	}

	fmt.Println("Result: ", crt(reminders, busNumbers))
	lib.Elapsed()
}

func crt(rems, mods []int) int {
	product := 1
	for _, mod := range mods {
		product *= mod
	}

	result := 0
	for i, mod := range mods {
		pp := product / mod
		result += rems[i] * inverseModulo(pp, mod) * pp
	}

	return result % product
}

func inverseModulo(pp, mod int) int {
	if mod == 1 {
		return 0
	}

	x0, x1 := 0, 1
	tempMod := mod

	for pp > 1 {
		q := pp / tempMod
		temp := tempMod

		tempMod = pp % tempMod
		pp = temp

		temp = x0
		x0 = x1 - q*x0
		x1 = temp
	}

	if x1 < 0 {
		x1 += mod
	}

	return x1
}
