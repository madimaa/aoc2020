package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

type replace func(regex, old, new string) string

func main() {
	part1()
	part2()
	os.Exit(0)
}

func part1() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("19.txt")
	lastLine, regex := regexpRules(input, replacePart1)
	messages := input[lastLine+1:]

	fmt.Println("Result: ", countCorrectMessages(messages, regex))
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("19.txt")
	lastLine, regex := regexpRules(input, replacePart2HACK)
	messages := input[lastLine+1:]

	//recursive workaround
	correctMessages := 0
	for i := 1; i < 10; i++ {
		r := strings.ReplaceAll(regex, "{X}", "{"+strconv.Itoa(i)+"}")
		correctMessages += countCorrectMessages(messages, r)
	}

	fmt.Println("Result: ", correctMessages)
	lib.Elapsed()
}

func regexpRules(input []string, r replace) (int, string) {
	rules := make(map[int]string)
	var lastLine int
	for index, row := range input {
		if len(row) == 0 {
			lastLine = index
			break
		}

		splitRow := strings.Split(row, ": ")
		key, _ := strconv.Atoi(splitRow[0])
		rules[key] = splitRow[1]
	}

	regex := "^ " + rules[0] + " "
	replacedInts := parseInts(rules[0])
	for len(replacedInts) > 0 {
		replaces := make([]int, 0)
		replaced := make(map[int]bool)
		for _, num := range replacedInts {
			strNum := strconv.Itoa(num)
			if !replaced[num] {
				replaced[num] = true
				regex = r(regex, strNum, rules[num])
				nums := parseInts(rules[num])
				for _, number := range nums {
					replaced[number] = false
					replaces = append(replaces, number)
				}
			}
		}

		replacedInts = replaces
	}

	regex = strings.ReplaceAll(regex, " ", "")
	regex = strings.ReplaceAll(regex, "\"", "")
	regex += "$"

	return lastLine, regex
}

func replacePart1(regex, old, new string) string {
	return strings.ReplaceAll(regex, " "+old+" ", " (?: "+new+" ) ")
}

func replacePart2(regex, old, new string) string {
	switch old {
	case "8":
		s := strings.ReplaceAll(regex, " 8 ", " (?: 42 )+ ")
		return s
	case "11":
		return strings.ReplaceAll(regex, " 11 ", " (?:(?: 42 )+ (?: 31 )+) ")
	default:
		return strings.ReplaceAll(regex, " "+old+" ", " (?: "+new+" ) ")
	}
}

func replacePart2HACK(regex, old, new string) string {
	switch old {
	case "8":
		s := strings.ReplaceAll(regex, " 8 ", " (?: 42 )+ ")
		return s
	case "11":
		return strings.ReplaceAll(regex, " 11 ", " (?:(?: 42 ){X} (?: 31 ){X}) ")
	default:
		return strings.ReplaceAll(regex, " "+old+" ", " (?: "+new+" ) ")
	}
}

func parseInts(input string) []int {
	nums := make(map[int]bool)
	for _, item := range strings.Split(input, " ") {
		num, err := strconv.Atoi(item)
		if err == nil {
			nums[num] = true
		}
	}

	result := make([]int, 0)
	for num := range nums {
		result = append(result, num)
	}

	return result
}

func countCorrectMessages(messages []string, regex string) int {
	correctMessages := 0

	for _, message := range messages {
		match, err := regexp.MatchString(regex, message)
		if err == nil && match {
			correctMessages++
		}
	}

	return correctMessages
}
