package main

import (
	"fmt"
	"os"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	part1()
	os.Exit(0)
}

func part1() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("17.txt")
	dim := CreatePocketDimension()

	for y, row := range input {
		for x, char := range []rune(row) {
			if char == '#' {
				dim.PutIntoDimension(CreateConwayCube(x, y, 1, char))
			}
		}
	}

	cycles := 6

	for cycles > 0 {
		tempDim := CycleDimension(dim)
		dim = CreatePocketDimension()
		for x := range tempDim.content {
			for y := range tempDim.content[x] {
				for z, cube := range tempDim.content[x][y] {
					if cube.status == '#' {
						dim.PutIntoDimension(CreateConwayCube(x, y, z, '#'))
					}
				}
			}
		}

		cycles--
	}

	actives := 0
	for x := range dim.content {
		for y := range dim.content[x] {
			for _, cube := range dim.content[x][y] {
				if cube.status == '#' {
					actives++
				}
			}
		}
	}

	fmt.Println("Result: ", actives)
	lib.Elapsed()
}
