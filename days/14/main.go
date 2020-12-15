package main

import (
	"fmt"
	"math"
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

	input := lib.OpenFile("14.txt")
	var mask []rune
	memory := make(map[int]int64)
	for _, row := range input {
		splitRow := strings.Split(row, " = ")
		if splitRow[0] == "mask" {
			mask = []rune(splitRow[1])
		} else {
			memAddr, _ := strconv.Atoi(splitRow[0][4 : len(splitRow[0])-1])
			value, _ := strconv.Atoi(splitRow[1])
			binary := strconv.FormatInt(int64(value), 2)
			for len(mask) > len(binary) {
				binary = "0" + binary
			}

			binaryRunes := []rune(binary)
			for index, maskValue := range mask {
				if maskValue != 'X' {
					binaryRunes[index] = maskValue
				}
			}

			memValue, _ := strconv.ParseInt(string(binaryRunes), 2, 64)

			memory[memAddr] = memValue
		}
	}

	var sum int64 = 0
	for _, val := range memory {
		sum += val
	}

	fmt.Println("Result: ", sum)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	var combinationMap map[int64][]rune
	xPositions := make(map[int]int)

	input := lib.OpenFile("14.txt")
	var mask []rune
	memory := make(map[int64]int64)
	for _, row := range input {
		splitRow := strings.Split(row, " = ")
		if splitRow[0] == "mask" {
			mask = []rune(splitRow[1])
			counter := 0
			for i, r := range mask {
				if r == 'X' {
					xPositions[counter] = i
					counter++
				}
			}

			combinationMap = generateCombinations(counter)
		} else {
			memAddr, _ := strconv.Atoi(splitRow[0][4 : len(splitRow[0])-1])
			value, _ := strconv.Atoi(splitRow[1])
			memAddrBinary := strconv.FormatInt(int64(memAddr), 2)
			for len(mask) > len(memAddrBinary) {
				memAddrBinary = "0" + memAddrBinary
			}

			binaryRunes := []rune(memAddrBinary)
			for index, maskValue := range mask {
				switch maskValue {
				case '1':
					binaryRunes[index] = '1'
				case 'X':
					binaryRunes[index] = 'X'
				}
			}

			for _, runes := range combinationMap {
				copy := copyRunes(binaryRunes)
				for index, r := range runes {
					copy[xPositions[index]] = r
				}

				addr, _ := strconv.ParseInt(string(copy), 2, 64)
				memory[addr] = int64(value)
			}
		}
	}

	var sum int64 = 0
	for _, val := range memory {
		sum += val
	}

	fmt.Println("Result: ", sum)
	lib.Elapsed()
}

func generateCombinations(length int) map[int64][]rune {
	result := make(map[int64][]rune)
	target := int64(math.Pow(2, float64(length)))
	for i := int64(0); i < target; i++ {
		binary := strconv.FormatInt(i, 2)
		for length > len(binary) {
			binary = "0" + binary
		}

		result[i] = []rune(binary)
	}

	return result
}

func copyRunes(cp []rune) []rune {
	ret := make([]rune, 0)
	for _, r := range cp {
		ret = append(ret, r)
	}
	return ret
}
