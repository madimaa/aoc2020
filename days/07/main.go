package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

type bags struct {
	bags []bag
}

type bag struct {
	name  string
	count int
}

func main() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("07.txt")
	bagMap := make(map[string]string)
	for _, fileRow := range input {
		keyValue := strings.Split(fileRow, " bags contain ")
		leftBag := keyValue[0]
		rightBags := keyValue[1]
		rightBags = rightBags[:len(rightBags)-1]
		rightBags = strings.ReplaceAll(rightBags, "bags", "bag")
		rightBags = strings.ReplaceAll(rightBags, "bag", "")
		matcher := regexp.MustCompile("[0-9]{1,2}")
		rightBags = matcher.ReplaceAllString(rightBags, "")
		for _, bag := range strings.Split(rightBags, ",") {
			bag = strings.TrimSpace(bag)
			if len(bagMap[bag]) > 0 {
				bagMap[bag] += ","
			}
			bagMap[bag] += leftBag
		}
	}

	canContainShinyGoldShouldCheck := make([]string, 0)
	for _, bag := range strings.Split(bagMap["shiny gold"], ",") {
		canContainShinyGoldShouldCheck = append(canContainShinyGoldShouldCheck, bag)
	}

	canContainShinyGoldChecked := make(map[string]bool)
	canContainShinyGoldCount := len(canContainShinyGoldShouldCheck)

	for {
		if len(canContainShinyGoldShouldCheck) == 0 {
			break
		}

		actual := canContainShinyGoldShouldCheck[0]
		canContainShinyGoldShouldCheck[0] = ""
		canContainShinyGoldShouldCheck = canContainShinyGoldShouldCheck[1:len(canContainShinyGoldShouldCheck)]
		canContainShinyGoldChecked[actual] = true

		for _, nextBag := range strings.Split(bagMap[actual], ",") {
			if len(nextBag) > 0 {
				if !lib.ContainsStr(canContainShinyGoldShouldCheck, nextBag) && !canContainShinyGoldChecked[nextBag] {
					canContainShinyGoldShouldCheck = append(canContainShinyGoldShouldCheck, nextBag)
					canContainShinyGoldCount++
				}
			}
		}
	}

	fmt.Println("Result: ", canContainShinyGoldCount)
	lib.Elapsed()

	lib.Start()
	fmt.Println("Part 2")

	bagception := make(map[string]bags)
	for _, fileRow := range input {
		keyValue := strings.Split(fileRow, " bags contain ")
		leftBag := keyValue[0]
		rightBags := keyValue[1]
		rightBags = rightBags[:len(rightBags)-1]
		rightBags = strings.ReplaceAll(rightBags, "bags", "bag")
		rightBags = strings.ReplaceAll(rightBags, "bag", "")
		bagContainer := make([]bag, 0)
		for _, actBag := range strings.Split(rightBags, ",") {
			actBag = strings.TrimSpace(actBag)
			actBagNumber, _ := strconv.Atoi(actBag[:1])
			actBag = strings.TrimSpace(actBag[1:])

			bag := bag{name: actBag, count: actBagNumber}
			bagContainer = append(bagContainer, bag)
		}

		bagception[leftBag] = bags{bags: bagContainer}
	}

	fmt.Println("Result: ", recursiveSum(bagception, "shiny gold", 1)-1) //minus the first shiny gold bag

	lib.Elapsed()
	os.Exit(0)
}

func recursiveSum(bagception map[string]bags, name string, cnt int) int {
	numberOfBags := 0
	numberOfBags += cnt

	if _, ok := bagception[name]; ok {
		for _, nextBag := range bagception[name].bags {
			if nextBag.count != 0 {
				numberOfBags += recursiveSum(bagception, nextBag.name, cnt*nextBag.count)
			} else {
				return cnt
			}
		}
	}

	return numberOfBags
}
