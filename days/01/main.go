package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	lib.Start()
	fmt.Println("Part 1")

	var target = 2020
	result := lib.OpenFile("01.txt")
	expMap := make(map[int]bool)
	for _, s := range result {
		number, err := strconv.Atoi(s)
		lib.LogOnError(err)
		expMap[number] = true
	}

	x, y := find(expMap, target)
	fmt.Println("Result: ", x, " * ", y, " = ", x*y)
	lib.Elapsed()

	lib.Start()
	fmt.Println("Part 2")

	for key := range expMap {
		diff := target - key

		if val := expMap[diff]; val == false {
			a, b := find(expMap, diff)
			if a != 0 && b != 0 {
				fmt.Println("Result: ", a, " * ", b, " * ", key, " = ", a*b*key)
				break
			}
		}
	}

	lib.Elapsed()
	os.Exit(0)
}

func find(expMap map[int]bool, target int) (int, int) {
	for key := range expMap {
		diff := target - key

		if val := expMap[diff]; val == true {
			return key, diff
		}
	}

	return 0, 0
}
