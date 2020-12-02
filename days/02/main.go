package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

type record struct {
	lowerPos, higherPos int
	key                 rune
	password            string
	valid               bool
}

func main() {
	lib.Start()
	fmt.Println("Part 1")

	validPasswords := 0
	input := lib.OpenFile("02.txt")
	records := make([]record, 0)
	for _, s := range input {
		content := strings.Split(s, " ")
		minMax := strings.Split(content[0], "-")
		min, err := strconv.Atoi(minMax[0])
		lib.LogOnError(err)
		max, err := strconv.Atoi(minMax[1])
		lib.LogOnError(err)

		rec := record{lowerPos: min, higherPos: max, key: []rune(content[1])[0], password: content[2]}
		if isValidOldJob(rec) {
			rec.valid = true
			validPasswords++
		}
		records = append(records, rec)
	}

	fmt.Println("Result: ", validPasswords)
	lib.Elapsed()

	lib.Start()
	fmt.Println("Part 2")

	validPasswords = 0
	for _, rec := range records {
		if isValidOTCPolicy(rec) {
			rec.valid = true
			validPasswords++
		} else {
			rec.valid = false
		}
	}

	fmt.Println("Result: ", validPasswords)

	lib.Elapsed()
	os.Exit(0)
}

func isValidOldJob(rec record) bool {
	identical := 0
	for _, r := range []rune(rec.password) {
		if r == rec.key {
			identical++
		}
	}

	return identical >= rec.lowerPos && identical <= rec.higherPos
}

func isValidOTCPolicy(rec record) bool {
	passwordRunes := []rune(rec.password)
	//Toboggan Corporate Policies have no concept of "index zero"! => subtract 1
	lower := passwordRunes[rec.lowerPos-1] == rec.key
	higher := passwordRunes[rec.higherPos-1] == rec.key

	return lower != higher
}
