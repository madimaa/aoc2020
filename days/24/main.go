package main

import (
	"fmt"
	"os"

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

	input := lib.OpenFile("24.txt")
	g := placeTiles(input)
	blackTiles := countBlackTiles(g)

	fmt.Println("Result: ", blackTiles)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("24.txt")
	g := placeTiles(input)
	days := 1
	for days <= 100 {
		g = g.flip()
		days++
	}

	blackTiles := countBlackTiles(g)

	fmt.Println("Result: ", blackTiles)
	lib.Elapsed()
}

func placeTiles(input []string) *grid {
	g := createGrid()
	for _, row := range input {
		t := createTile(0, 0, true)
		runes := []rune(row)
		for i := 0; i < len(runes); {
			r := runes[i]
			switch r {
			case 'n', 's':
				str := string(r) + string(runes[i+1])
				t.move(str)
				i += 2
			default:
				str := string(r)
				t.move(str)
				i++
			}
		}

		if g.get(t.x, t.y) == nil {
			g.putIntoGrid(t)
		} else {
			g.get(t.x, t.y).color = !g.get(t.x, t.y).color
		}
	}

	return g
}

func countBlackTiles(g *grid) int {
	blackTiles := 0
	for _, val := range g.content {
		for _, t := range val {
			if t.color {
				blackTiles++
			}
		}
	}

	return blackTiles
}
