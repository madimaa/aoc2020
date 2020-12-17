package main

import (
	"fmt"
	"os"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	part2()
	os.Exit(0)
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("17.txt")
	dim := CreatePocketDimension()

	for y, row := range input {
		for x, char := range []rune(row) {
			if char == '#' {
				dim.PutIntoDimension(CreateConwayCube(x, y, 0, 0, char))
			}
		}
	}

	cycles := 6

	for cycles > 0 {
		tempDim := CycleDimension(dim)
		dim = CreatePocketDimension()
		for x := range tempDim.content {
			for y := range tempDim.content[x] {
				for z := range tempDim.content[x][y] {
					for w, cube := range tempDim.content[x][y][z] {
						if cube.status == '#' {
							dim.PutIntoDimension(CreateConwayCube(x, y, z, w, '#'))
						}
					}
				}
			}
		}

		cycles--
	}

	actives := 0
	for x := range dim.content {
		for y := range dim.content[x] {
			for z := range dim.content[x][y] {
				for _, cube := range dim.content[x][y][z] {
					if cube.status == '#' {
						actives++
					}
				}
			}
		}
	}

	fmt.Println("Result: ", actives)
	lib.Elapsed()
}
