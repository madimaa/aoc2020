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

	input := lib.OpenFile("19.txt")
	lastLine, rules := parseRules(input)
	possibleMesages := generatePossibleMessages(rules)
	messages := input[lastLine+1:]

	fmt.Println("Result: ", countCorrectMessages(messages, possibleMesages))
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	fmt.Println("Result: ", 1)
	lib.Elapsed()
}

func parseRules(input []string) (int, map[int]Rule) {
	rules := make(map[int]Rule)
	for index, row := range input {
		if len(row) == 0 {
			return index, rules
		}

		splitRow := strings.Split(row, ": ")
		key, _ := strconv.Atoi(splitRow[0])
		subrules := make([]Subrule, 0)
		switch {
		case strings.Contains(splitRow[1], "\""):
			rules[key] = Rule{letter: []rune(splitRow[1])[1]}
		case strings.Contains(splitRow[1], "|"):
			splitSubrules := strings.Split(splitRow[1], " | ")
			for _, subruleString := range splitSubrules {
				subrule := Subrule{subrules: parseInts(subruleString)}
				subrules = append(subrules, subrule)
			}

			rules[key] = Rule{subrule: subrules}
		default:
			subrule := Subrule{subrules: parseInts(splitRow[1])}
			subrules = append(subrules, subrule)
			rules[key] = Rule{subrule: subrules}
		}
	}

	panic("Something wrong with the input, the should have never happened.")
}

func parseInts(input string) []int {
	result := make([]int, 0)
	for _, item := range strings.Split(input, " ") {
		num, _ := strconv.Atoi(item)
		result = append(result, num)
	}

	return result
}

func generatePossibleMessages(rules map[int]Rule) map[string]bool {
	result := make(map[string]bool)

	return result
}

func countCorrectMessages(messages []string, possibleMesages map[string]bool) int {
	correctMessages := 0

	for _, message := range messages {
		if _, ok := possibleMesages[message]; ok {
			correctMessages++
		}
	}

	return correctMessages
}
