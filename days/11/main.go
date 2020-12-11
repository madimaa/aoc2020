package main

import (
	"fmt"
	"os"

	"github.com/madimaa/aoc2020/lib"
	"github.com/madimaa/aoc2020/lib/array2d"
)

func main() {
	part1()
	part2()
	os.Exit(0)
}

func part1() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("11.txt")
	xMax := len(input[0])
	yMax := len(input)
	matrix := array2d.Create(xMax, yMax)
	for y, fileRow := range input {
		for x, act := range []rune(fileRow) {
			matrix.Put(x, y, act)
		}
	}

	original := matrix.Copy()
	for {
		for y := 0; y < yMax; y++ {
			for x := 0; x < xMax; x++ {
				switch original.Get(x, y).(rune) {
				case '#':
					if getAdjacentTakenSeats(original, x, y, xMax, yMax) >= 4 {
						matrix.Put(x, y, 'L')
					}
				case 'L':
					if getAdjacentTakenSeats(original, x, y, xMax, yMax) == 0 {
						matrix.Put(x, y, '#')
					}
				case '.':
					//nothing
				default:
					panic("Something went wrong!")
				}
			}
		}

		//fmt.Println("---------------------------------------------------------------------")
		//drawToScreen(matrix, xMax, yMax)
		//fmt.Println("---------------------------------------------------------------------")
		//drawToScreen(original, xMax, yMax)
		//fmt.Println("---------------------------------------------------------------------")

		if original.Equals(matrix) {
			break
		} else {
			original = matrix.Copy()
		}
	}

	//drawToScreen(matrix, xMax, yMax)
	seats := 0
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if matrix.Get(x, y).(rune) == '#' {
				seats++
			}
		}
	}

	fmt.Println("Result: ", seats)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("11.txt")
	xMax := len(input[0])
	yMax := len(input)
	matrix := array2d.Create(xMax, yMax)
	for y, fileRow := range input {
		for x, act := range []rune(fileRow) {
			matrix.Put(x, y, act)
		}
	}

	original := matrix.Copy()
	for {
		for y := 0; y < yMax; y++ {
			for x := 0; x < xMax; x++ {
				switch original.Get(x, y).(rune) {
				case '#':
					if getTakensSeatsAtSight(original, x, y, xMax, yMax) >= 5 {
						matrix.Put(x, y, 'L')
					}
				case 'L':
					if getTakensSeatsAtSight(original, x, y, xMax, yMax) == 0 {
						matrix.Put(x, y, '#')
					}
				case '.':
					//nothing
				default:
					panic("Something went wrong!")
				}
			}
		}

		//fmt.Println("---------------------------------------------------------------------")
		//drawToScreen(matrix, xMax, yMax)
		//fmt.Println("---------------------------------------------------------------------")
		//drawToScreen(original, xMax, yMax)
		//fmt.Println("---------------------------------------------------------------------")

		if original.Equals(matrix) {
			break
		} else {
			original = matrix.Copy()
		}
	}

	//drawToScreen(matrix, xMax, yMax)
	seats := 0
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if matrix.Get(x, y).(rune) == '#' {
				seats++
			}
		}
	}

	fmt.Println("Result: ", seats)
	lib.Elapsed()
}

func getAdjacentTakenSeats(a2d *array2d.Array2D, x, y, xMax, yMax int) int {
	taken := 0
	startX, endX := x-1, x+1
	startY, endY := y-1, y+1

	if startX < 0 {
		startX = 0
	}

	if endX >= xMax {
		endX = xMax - 1
	}

	if startY < 0 {
		startY = 0
	}

	if endY >= yMax {
		endY = yMax - 1
	}

	for i := startY; i <= endY; i++ {
		for j := startX; j <= endX; j++ {
			if i == y && j == x {
				continue
			}

			if a2d.Get(j, i).(rune) == '#' {
				taken++
			}
		}
	}

	return taken
}

func getTakensSeatsAtSight(a2d *array2d.Array2D, x, y, xMax, yMax int) int {
	taken := 0
	taken += takenSeatAtSightDirection(a2d, x, y, xMax, yMax, 0, -1)
	taken += takenSeatAtSightDirection(a2d, x, y, xMax, yMax, 1, -1)
	taken += takenSeatAtSightDirection(a2d, x, y, xMax, yMax, 1, 0)
	taken += takenSeatAtSightDirection(a2d, x, y, xMax, yMax, 1, 1)
	taken += takenSeatAtSightDirection(a2d, x, y, xMax, yMax, 0, 1)
	taken += takenSeatAtSightDirection(a2d, x, y, xMax, yMax, -1, 1)
	taken += takenSeatAtSightDirection(a2d, x, y, xMax, yMax, -1, 0)
	taken += takenSeatAtSightDirection(a2d, x, y, xMax, yMax, -1, -1)

	return taken
}

func takenSeatAtSightDirection(a2d *array2d.Array2D, x, y, xMax, yMax, xDir, yDir int) int {
	for {
		x += xDir
		y += yDir

		if x < 0 || x >= xMax || y < 0 || y >= yMax {
			return 0
		}

		switch a2d.Get(x, y).(rune) {
		case '#':
			return 1
		case 'L':
			return 0
		}
	}
}

func drawToScreen(matrix *array2d.Array2D, xMax, yMax int) {
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			fmt.Print(string(matrix.Get(x, y).(rune)))
		}
		fmt.Println()
	}
}
