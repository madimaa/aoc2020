package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

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

	input := lib.OpenFile("21.txt")
	possibilities := make(map[string]map[string]int)
	for _, row := range input {
		split := strings.Split(row, " (contains ")
		ingredients := split[0]
		allergens := split[1]
		allergens = allergens[:len(allergens)-1]
		for _, ingredient := range strings.Split(ingredients, " ") {
			for _, allergen := range strings.Split(allergens, ", ") {
				if _, ok := possibilities[allergen]; !ok {
					possibilities[allergen] = make(map[string]int)
				}
				possibilities[allergen][ingredient]++
			}
		}
	}

	allergens := make(map[string]string)
	ingredients := make(map[string]string)
	for {
		ingredient := ""
		allergen := ""
		for all, ings := range possibilities {
			max := 0
			eq := false
			maxIng := ""

			for ing, num := range ings {
				if max == num {
					eq = true
				} else if max < num {
					eq = false
					max = num
					maxIng = ing
				}
			}

			if max > 0 && !eq {
				ingredient = maxIng
				allergen = all
				break
			}
		}

		if ingredient != "" && allergen != "" {
			allergens[allergen] = ingredient
			ingredients[ingredient] = allergen
			delete(possibilities, allergen)
			for all, ings := range possibilities {
				if _, ok := ings[ingredient]; ok {
					delete(possibilities[all], ingredient)
				}
			}

			allEmpty := true
			for _, alls := range possibilities {
				if len(alls) > 0 {
					allEmpty = false
					break
				}
			}

			if allEmpty {
				break
			}
		} else {
			panic("There must be an error.")
		}
	}

	notAllergen := 0
	for _, row := range input {
		split := strings.Split(row, " (contains ")
		ings := split[0]
		for _, ing := range strings.Split(ings, " ") {
			if _, ok := ingredients[ing]; !ok {
				notAllergen++
			}
		}
	}

	fmt.Println("Result: ", notAllergen)
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	input := lib.OpenFile("21.txt")
	possibilities := make(map[string]map[string]int)
	for _, row := range input {
		split := strings.Split(row, " (contains ")
		ingredients := split[0]
		allergens := split[1]
		allergens = allergens[:len(allergens)-1]
		for _, ingredient := range strings.Split(ingredients, " ") {
			for _, allergen := range strings.Split(allergens, ", ") {
				if _, ok := possibilities[allergen]; !ok {
					possibilities[allergen] = make(map[string]int)
				}
				possibilities[allergen][ingredient]++
			}
		}
	}

	allergens := make(map[string]string)
	ingredients := make(map[string]string)
	for {
		ingredient := ""
		allergen := ""
		for all, ings := range possibilities {
			max := 0
			eq := false
			maxIng := ""

			for ing, num := range ings {
				if max == num {
					eq = true
				} else if max < num {
					eq = false
					max = num
					maxIng = ing
				}
			}

			if max > 0 && !eq {
				ingredient = maxIng
				allergen = all
				break
			}
		}

		if ingredient != "" && allergen != "" {
			allergens[allergen] = ingredient
			ingredients[ingredient] = allergen
			delete(possibilities, allergen)
			for all, ings := range possibilities {
				if _, ok := ings[ingredient]; ok {
					delete(possibilities[all], ingredient)
				}
			}

			allEmpty := true
			for _, alls := range possibilities {
				if len(alls) > 0 {
					allEmpty = false
					break
				}
			}

			if allEmpty {
				break
			}
		} else {
			panic("There must be an error.")
		}
	}

	allergenKeys := make([]string, 0, len(allergens))
	for k := range allergens {
		allergenKeys = append(allergenKeys, k)
	}

	sort.Strings(allergenKeys)

	result := ""
	for _, k := range allergenKeys {
		result += allergens[k] + ","
	}

	result = result[:len(result)-1]

	fmt.Println("Result: ", result)
	lib.Elapsed()
}
