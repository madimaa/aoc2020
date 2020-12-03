package main

import (
	"fmt"
	"os"

	"github.com/madimaa/aoc2020/lib"
	"github.com/madimaa/aoc2020/lib/array2d"
)

var width, height int

type record struct {
	lowerPos, higherPos int
	key                 rune
	password            string
	valid               bool
}

func main() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("03.txt")
	width, height = 0, 0
	s := ""
	for height, s = range input {
		if width == 0 {
			width = len(s)
		}
	}

	height++

	pattern := array2d.Create(width, height)
	for y, s := range input {
		for x, val := range []rune(s) {
			pattern.Put(x, y, string(val))
		}
	}

	res13 := slideDown(1, 3, 0, 0, pattern, height)
	fmt.Println("Result: ", res13)
	lib.Elapsed()

	lib.Start()
	fmt.Println("Part 2")

	res11 := slideDown(1, 1, 0, 0, pattern, height)
	res15 := slideDown(1, 5, 0, 0, pattern, height)
	res17 := slideDown(1, 7, 0, 0, pattern, height)
	res21 := slideDown(2, 1, 0, 0, pattern, height)

	fmt.Println("Result: ", res11*res13*res15*res17*res21)

	lib.Elapsed()
	os.Exit(0)
}

func slideDown(vertical, horizontal, startX, startY int, pattern *array2d.Array2D, height int) int {
	trees := 0
	actualX, actualY := startX, startY

	for true {
		actualX = (actualX + horizontal) % width
		actualY = actualY + vertical
		if isTree(actualX, actualY, pattern) {
			trees++
		}

		if actualY >= height {
			break
		}
	}

	return trees
}

func isTree(x, y int, pattern *array2d.Array2D) bool {
	if pattern.Get(x, y) == "#" {
		return true
	}

	return false
}
