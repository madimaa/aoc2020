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

	input := lib.OpenFile("16.txt")
	lastLine := 0
	temp1, ticketFields := parseFields(input)
	lastLine += temp1 + 1

	temp2, _ := parseTickets(input[lastLine:])
	lastLine += temp2 + 1

	_, nearbyTickets := parseTickets(input[lastLine:])

	errorRate := 0
	for _, numbers := range nearbyTickets {
		for _, number := range numbers {
			if !ticketFields.ContainsAny(number) {
				errorRate += number
				break
			}
		}
	}

	fmt.Println("Result: ", errorRate)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("16.txt")
	lastLine := 0
	temp1, ticketFields := parseFields(input)
	lastLine += temp1 + 1

	temp2, myTicket := parseTickets(input[lastLine:])
	lastLine += temp2 + 1

	_, nearbyTickets := parseTickets(input[lastLine:])
	validTickets := make(map[int][]int)

	for i, numbers := range nearbyTickets {
		invalidTicket := false
		for _, number := range numbers {
			if !ticketFields.ContainsAny(number) {
				invalidTicket = true
				break
			}
		}

		if !invalidTicket {
			validTickets[i] = numbers
		}
	}

	result := 1
	rightPositions := determineFieldPositions(validTickets, ticketFields)
	for name, pos := range rightPositions {
		if strings.HasPrefix(name, "departure") {
			result *= myTicket[1][pos]
		}
	}

	fmt.Println("Result: ", result)
	lib.Elapsed()
}

func parseFields(input []string) (int, fields) {
	result := fields{data: make(map[string]field)}

	lastReadLine := 0
	for i, row := range input {
		if len(row) == 0 {
			lastReadLine = i
			break
		}

		name := strings.Split(row, ": ")[0]
		values := strings.Split(strings.Split(row, ": ")[1], " or ")
		lower := strings.Split(values[0], "-")
		higher := strings.Split(values[1], "-")
		lowerLower, _ := strconv.Atoi(lower[0])
		lowerHigher, _ := strconv.Atoi(lower[1])
		higherLower, _ := strconv.Atoi(higher[0])
		higherHigher, _ := strconv.Atoi(higher[1])

		f := field{name: name, lowerBoundaries: pair{lower: lowerLower, higher: lowerHigher}, higherBoundaries: pair{lower: higherLower, higher: higherHigher}}
		result.data[name] = f
	}

	return lastReadLine, result
}

func parseTickets(input []string) (int, map[int][]int) {
	result := make(map[int][]int)

	lastReadLine := 0
	for i, row := range input {
		if i == 0 {
			continue
		}

		if len(row) == 0 {
			lastReadLine = i
			break
		}

		slice := make([]int, 0)
		for _, val := range strings.Split(row, ",") {
			num, _ := strconv.Atoi(val)
			slice = append(slice, num)
		}

		result[i] = slice
	}

	return lastReadLine, result
}

func determineFieldPositions(validTickets map[int][]int, ticketFields fields) map[string]int {
	possibilities := make(map[string][]int)

	position := 0

	var max int
	for _, numbers := range validTickets {
		max = len(numbers)
		break
	}

	for position < max {
		numbers := make([]int, 0)
		for _, nums := range validTickets {
			numbers = append(numbers, nums[position])
		}

		for name, f := range ticketFields.data {
			ok := true
			for _, num := range numbers {
				if !f.Contains(num) {
					ok = false
					break
				}
			}

			if ok {
				if _, ok := possibilities[name]; !ok {
					possibilities[name] = make([]int, 0)
				}

				possibilities[name] = append(possibilities[name], position)
			}
		}

		position++
	}

	rightPosition := make(map[string]int)
	for len(possibilities) > 0 {
		var shouldDeleteName string
		var shouldDeletePosition int
		for name, numbers := range possibilities {
			if len(numbers) == 1 {
				rightPosition[name] = numbers[0]
				shouldDeleteName = name
				shouldDeletePosition = numbers[0]
				break
			}
		}

		delete(possibilities, shouldDeleteName)
		for name, numbers := range possibilities {
			if lib.ContainsInt(numbers, shouldDeletePosition) {
				newArr := make([]int, 0)
				for _, num := range numbers {
					if num != shouldDeletePosition {
						newArr = append(newArr, num)
					}
				}

				possibilities[name] = newArr
			}
		}
	}

	return rightPosition
}
