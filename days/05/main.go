package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("05.txt")
	var highestSeatID int64 = 0
	seatIDs := make(map[int64]bool)
	for _, fileRow := range input {
		length := len(fileRow)
		row := fileRow[:length-3]
		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")

		column := fileRow[length-3:]
		column = strings.ReplaceAll(column, "L", "0")
		column = strings.ReplaceAll(column, "R", "1")

		rowNum, _ := strconv.ParseInt(row, 2, 8)
		colNum, _ := strconv.ParseInt(column, 2, 8)

		result := rowNum*8 + colNum
		if highestSeatID < result {
			highestSeatID = result
		}

		seatIDs[result] = true
	}

	fmt.Println("Result: ", highestSeatID)
	lib.Elapsed()

	lib.Start()
	fmt.Println("Part 2")

	var mySeatID int64
	var i int64
	for i = 1; i < highestSeatID; i++ {
		if !seatIDs[i] {
			if seatIDs[i-1] && seatIDs[i+1] {
				mySeatID = i
				break
			}
		}
	}

	fmt.Println("Result: ", mySeatID)

	lib.Elapsed()
	os.Exit(0)
}
