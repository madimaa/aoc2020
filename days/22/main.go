package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

func main() {
	fmt.Println("On Windows change the input file's line endings to LF")
	part1()
	part2()
	os.Exit(0)
}

func part1() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFileAsString("22.txt")
	player1, player2 := parseInput(input)
	result := playCombat(player1, player2)

	score := 0
	multiplier := 1
	for i := len(result) - 1; i >= 0; i-- {
		score += multiplier * result[i]
		multiplier++
	}

	fmt.Println("Result: ", score)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFileAsString("22.txt")
	player1, player2 := parseInput(input)
	_, result := playRecursiveCombat(player1, player2)

	score := 0
	multiplier := 1
	for i := len(result) - 1; i >= 0; i-- {
		score += multiplier * result[i]
		multiplier++
	}

	fmt.Println("Result: ", score)
	lib.Elapsed()
}

func parseInput(input string) ([]int, []int) {
	players := strings.Split(input, "\n\n")
	player1 := parsePlayerCards(players[0])
	player2 := parsePlayerCards(players[1])

	return player1, player2
}

func parsePlayerCards(input string) []int {
	player := strings.Split(input, "\n")
	cards := make([]int, 0)
	for i := 1; i < len(player); i++ {
		num, err := strconv.Atoi(player[i])
		lib.PanicOnError(err)
		cards = append(cards, num)
	}

	return cards
}

func playCombat(player1, player2 []int) []int {
	for len(player1) > 0 && len(player2) > 0 {
		card1 := player1[0]
		player1 = player1[1:len(player1)]

		card2 := player2[0]
		player2 = player2[1:len(player2)]

		if card1 > card2 {
			player1 = append(player1, card1, card2)
		} else {
			player2 = append(player2, card2, card1)
		}
	}

	if len(player1) != 0 {
		return player1
	}

	return player2
}

func playRecursiveCombat(player1, player2 []int) (int, []int) {
	previousRounds := make(map[string]bool)

	for len(player1) > 0 && len(player2) > 0 {
		round := fmt.Sprint(player1, player2)
		if previousRounds[round] {
			return 1, player1
		}

		previousRounds[round] = true

		card1 := player1[0]
		player1 = player1[1:len(player1)]

		card2 := player2[0]
		player2 = player2[1:len(player2)]

		if card1 > len(player1) || card2 > len(player2) {
			if card1 > card2 {
				player1 = append(player1, card1, card2)
			} else {
				player2 = append(player2, card2, card1)
			}
		} else {
			newPlayer1 := make([]int, 0)
			newPlayer2 := make([]int, 0)

			for i := 0; i < card1; i++ {
				newPlayer1 = append(newPlayer1, player1[i])
			}

			for i := 0; i < card2; i++ {
				newPlayer2 = append(newPlayer2, player2[i])
			}

			player, _ := playRecursiveCombat(newPlayer1, newPlayer2)
			if player == 1 {
				player1 = append(player1, card1, card2)
			} else {
				player2 = append(player2, card2, card1)
			}
		}
	}

	if len(player1) != 0 {
		return 1, player1
	}

	return 2, player2
}
