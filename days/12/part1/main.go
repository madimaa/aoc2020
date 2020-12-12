package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	part1()
	os.Exit(0)
}

func part1() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("12.txt")
	walker := CreateWalker(0, 0, 'E')
	for _, fileRow := range input {
		command := []rune(fileRow)[0]
		unit, _ := strconv.Atoi(fileRow[1:])
		switch command {
		case 'L', 'R':
			walker.Turn(command, unit)
		default:
			walker.Move(command, unit)
		}
	}

	endX, endY := walker.Position()

	fmt.Println("Result: ", math.Abs(float64(endX))+math.Abs(float64(endY)))
	lib.Elapsed()
}
