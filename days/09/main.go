package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	position := part1()
	part2(position)
	os.Exit(0)
}

func part1() int {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("09.txt")
	numbers := amble{numbers: make([]int, 0), numMap: make(map[int]bool)}
	var res int
	var pos int
	for i, fileRow := range input {
		num, _ := strconv.Atoi(fileRow)

		if numbers.length() == 25 {
			if numbers.hasOperands(num) {
				numbers.add(num)
			} else {
				res = num
				pos = i
				break
			}
		} else {
			numbers.add(num)
		}
	}

	fmt.Println("Result: ", res)
	lib.Elapsed()
	return pos
}

func part2(pos int) {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("09.txt")
	sum := 0
	numList := make([]int, 0)
	target, _ := strconv.Atoi(input[pos])
	for i := pos - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(input[i])
		sum += num
		numList = append(numList, num)

		if sum == target {
			break
		} else if sum > target {
			sum -= numList[0]
			numList = numList[1:]
		}
	}

	min, max := numList[0], numList[0]

	for _, num := range numList {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	fmt.Println("Result: ", min+max)
	lib.Elapsed()
}

type amble struct {
	numbers []int
	numMap  map[int]bool
}

func (a *amble) length() int {
	return len(a.numbers)
}

func (a *amble) add(num int) {
	a.numbers = append(a.numbers, num)
	a.numMap[num] = true

	if a.length() > 25 {
		delete(a.numMap, a.numbers[0]) //what about duplicates?
		a.numbers = a.numbers[1:]
	}
}

func (a *amble) contains(num int) bool {
	return a.numMap[num]
}

func (a *amble) hasOperands(target int) bool {
	for _, num := range a.numbers {
		if a.contains(target - num) {
			return true
		}
	}

	return false
}
