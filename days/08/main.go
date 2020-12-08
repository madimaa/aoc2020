package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

type instruction struct {
	op  string
	arg int
}

func main() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("08.txt")
	instructions := make([]instruction, 0)
	for _, fileRow := range input {
		row := strings.Split(fileRow, " ")
		num, _ := strconv.Atoi(row[1])
		instructions = append(instructions, instruction{op: row[0], arg: num})
	}

	visited := make(map[int]bool)
	lineNum := 0
	accumulator := 0
	for {
		if visited[lineNum] {
			break
		} else {
			visited[lineNum] = true
		}

		instr := instructions[lineNum]
		switch instr.op {
		case "acc":
			accumulator += instr.arg
			lineNum++
		case "jmp":
			lineNum += instr.arg
		case "nop":
			lineNum++
		default:
			panic("Something went wrong!")
		}
	}

	fmt.Println("Result: ", accumulator)
	lib.Elapsed()

	lib.Start()
	fmt.Println("Part 2")

	visited = make(map[int]bool)
	lineNum = 0
	accumulator = 0
	change := false
	originalOp := ""
	originalLineNum := 0
	originalAccumulator := 0
	originalVisited := make(map[int]bool)
	changed := make(map[int]bool)
	for lineNum < len(instructions) {
		if visited[lineNum] {
			visited = copyMap(originalVisited)
			lineNum = originalLineNum
			accumulator = originalAccumulator
			instructions[lineNum].op = originalOp
			change = false
		} else {
			visited[lineNum] = true
		}

		instr := instructions[lineNum]
		switch instr.op {
		case "acc":
			accumulator += instr.arg
			lineNum++
		case "jmp":
			if !change && !changed[lineNum] {
				changed[lineNum] = true
				instructions[lineNum].op = "nop"
				originalOp = "jmp"
				originalAccumulator = accumulator
				originalLineNum = lineNum
				originalVisited = copyMap(visited)
				change = true
				visited[lineNum] = false
			} else {
				lineNum += instr.arg
			}
		case "nop":
			if !change && !changed[lineNum] {
				changed[lineNum] = true
				instructions[lineNum].op = "jmp"
				originalOp = "nop"
				originalAccumulator = accumulator
				originalLineNum = lineNum
				originalVisited = copyMap(visited)
				change = true
				visited[lineNum] = false
			} else {
				lineNum++
			}
		default:
			panic("Something went wrong!")
		}
	}

	fmt.Println("Result: ", accumulator)

	lib.Elapsed()
	os.Exit(0)
}

func copyMap(original map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for k, v := range original {
		result[k] = v
	}

	return result
}
