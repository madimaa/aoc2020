package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

type solvePart func(eq string) int

func main() {
	part1()
	part2()
	os.Exit(0)
}

func part1() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("18.txt")
	sum := 0
	for _, row := range input {
		sum += solveEquation(row, solvePart1)
	}

	fmt.Println("Result: ", sum)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("18.txt")
	sum := 0
	for _, row := range input {
		sum += solveEquation(row, solvePart2)
	}

	fmt.Println("Result: ", sum)
	lib.Elapsed()
}

func solveEquation(equation string, solve solvePart) int {
	for strings.ContainsRune(equation, '(') {
		lastIndexOfOpenBracket := 0
	loop:
		for index, char := range []rune(equation) {
			switch char {
			case '(':
				lastIndexOfOpenBracket = index
			case ')':
				part := equation[lastIndexOfOpenBracket+1 : index]
				res := solve(part)
				equation = strings.Replace(equation, "("+part+")", strconv.Itoa(res), 1)
				break loop
			}
		}
	}

	return solve(equation)
}

func solvePart1(equation string) int {
	lastOperator := "x"
	result := 0
	for _, item := range strings.Split(equation, " ") {
		num, err := strconv.Atoi(item)
		if err != nil {
			lastOperator = item
		} else {
			switch lastOperator {
			case "x", "+":
				result += num
			case "*":
				result *= num
			}
		}
	}

	return result
}

func solvePart2(equation string) int {
	for strings.Contains(equation, "+") {
		items := strings.Split(equation, " ")
		for index, item := range items {
			if item == "+" {
				a, _ := strconv.Atoi(items[index-1])
				b, _ := strconv.Atoi(items[index+1])

				sum := a + b

				equation = strings.Replace(equation, items[index-1]+" + "+items[index+1], strconv.Itoa(sum), 1)
				break
			}
		}
	}

	result := 1
	if strings.Contains(equation, "*") {
		for _, item := range strings.Split(equation, " * ") {
			num, _ := strconv.Atoi(item)
			result *= num
		}
	} else {
		num, _ := strconv.Atoi(equation)
		result = num
	}

	return result
}
