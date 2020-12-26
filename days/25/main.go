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

	input := lib.OpenFile("25.txt")
	cardPubKey, err := strconv.Atoi(input[0])
	lib.PanicOnError(err)
	doorPubKey, err := strconv.Atoi(input[1])
	lib.PanicOnError(err)

	cardLoopSize, doorLoopSize := determineLoopSize(cardPubKey, doorPubKey)
	result := calculateEncryptionKey(doorPubKey, cardLoopSize)
	if result != calculateEncryptionKey(cardPubKey, doorLoopSize) {
		panic("Encryption key not found.")
	}

	fmt.Println("Result: ", result)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	lib.Elapsed()
}

func determineLoopSize(cardPubKey, doorPubKey int) (int, int) {
	cardLoopSize, doorLoopSize := 0, 0

	value := 1
	subjectNumber := 7
	loop := 1
	for cardLoopSize == 0 || doorLoopSize == 0 {
		value = transform(value, subjectNumber)

		if value == cardPubKey {
			cardLoopSize = loop
		}

		if value == doorPubKey {
			doorLoopSize = loop
		}

		loop++
	}

	return cardLoopSize, doorLoopSize
}

func transform(value, subjectNumber int) int {
	value *= subjectNumber
	value %= 20201227
	return value
}

func calculateEncryptionKey(subjectNumber, loop int) int {
	value := 1
	for loop > 0 {
		value = transform(value, subjectNumber)
		loop--
	}

	return value
}
