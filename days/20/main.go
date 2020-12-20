package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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

	input := lib.OpenFile("20.txt")
	xBound, yBound := determineBoundaries(input[1:])
	images := readImages(input, xBound, yBound)
	pictureSize := math.Sqrt(float64(len(images)))

	storage := make(map[string]*Image)
	for _, image := range images {
		if _, ok := storage[image.n]; ok {
			image.nID = storage[image.n].id
			registerPair(storage[image.n], image.id, image.n)
			delete(storage, image.n)
		} else {
			storage[image.n] = image
		}

		if _, ok := storage[image.e]; ok {
			image.eID = storage[image.e].id
			registerPair(storage[image.e], image.id, image.e)
			delete(storage, image.e)
		} else {
			storage[image.e] = image
		}

		if _, ok := storage[image.s]; ok {
			image.sID = storage[image.s].id
			registerPair(storage[image.s], image.id, image.s)
			delete(storage, image.s)
		} else {
			storage[image.s] = image
		}

		if _, ok := storage[image.w]; ok {
			image.wID = storage[image.w].id
			registerPair(storage[image.w], image.id, image.w)
			delete(storage, image.w)
		} else {
			storage[image.w] = image
		}
	}

	for _, img := range images {
		img.Print()
	}

	fmt.Println("Result: ", pictureSize, len(images))
	lib.Elapsed()
}

func part2() {
	lib.Start()
	fmt.Println("Part 2")

	fmt.Println("Result: ", 1)
	lib.Elapsed()
}

func determineBoundaries(input []string) (int, int) {
	x, y := 0, 0
	for index, row := range input {
		if len(row) == 0 {
			x = len(input[index-1])
			y = index
			break
		}
	}

	return x, y
}

func readImages(input []string, xBound, yBound int) []*Image {
	images := make([]*Image, 0)

	for i := 0; i < len(input); i++ {
		keyString := strings.TrimLeft(input[i], "Title ")
		keyString = keyString[:len(keyString)-1]
		key, _ := strconv.Atoi(keyString)
		i++
		images = append(images, CreateImage(key, input[i:i+yBound]))
		i += yBound
	}

	return images
}

func registerPair(img *Image, id int, str string) {
	switch str {
	case img.n:
		img.nID = id
	case img.e:
		img.eID = id
	case img.s:
		img.sID = id
	case img.w:
		img.wID = id
	}
}
