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

	buildPicture(images, int(pictureSize)*4, storage)
	result := 1
	for _, img := range images {
		if img.IsCorner() {
			result *= img.id
		}
	}

	fmt.Println("Result: ", result)

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

func readImages(input []string, xBound, yBound int) map[int]*Image {
	images := make(map[int]*Image)

	for i := 0; i < len(input); i++ {
		keyString := strings.TrimLeft(input[i], "Title ")
		keyString = keyString[:len(keyString)-1]
		key, _ := strconv.Atoi(keyString)
		i++
		images[key] = CreateImage(key, input[i:i+yBound])
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

func buildPicture(images map[int]*Image, numberOfEdges int, storage map[string]*Image) *Space {
	foundEdges := 0

	space := CreateSpace()
	recentlyPutIntoSpace := make([]*Image, 0)
	recentHelper := make(map[int]bool)
	start := findCompletedSegment(images)
	space.Put(0, 0, start)
	recentlyPutIntoSpace = append(recentlyPutIntoSpace, start)
	recentHelper[start.id] = true
	placed := make(map[int]bool)
	for foundEdges != numberOfEdges {
		actual := recentlyPutIntoSpace[0]
		placed[actual.id] = true
		recentlyPutIntoSpace = recentlyPutIntoSpace[1:]
		delete(recentHelper, actual.id)
		switch actual.nID {
		case 0:
			rev := reverseString(actual.n)
			if img, ok := storage[rev]; ok && img != actual {
				actual.nID = img.id
			} else if img, ok := storage[actual.n]; ok && img != actual {
				actual.nID = img.id
			} else {
				foundEdges++
				break
			}

			fallthrough
		default:
			other := images[actual.nID]
			if !placed[other.id] {
				space.Put(start.spaceX, start.spaceY-1, other)
				if _, ok := recentHelper[other.id]; !ok {
					recentlyPutIntoSpace = append(recentlyPutIntoSpace, other)
					recentHelper[other.id] = true
				}
				counter := 0
				for other.s != actual.n {
					switch counter % 5 {
					case 0:
						other.Rotate(false)
					case 1, 2:
						other.HorizontalFlip(false)
					case 3, 4:
						other.VerticalFlip(false)
					}
					counter++
				}

				other.sID = actual.id
			}
		}
		switch actual.eID {
		case 0:
			rev := reverseString(actual.e)
			if img, ok := storage[rev]; ok && img != actual {
				actual.eID = img.id
			} else if img, ok := storage[actual.e]; ok && img != actual {
				actual.eID = img.id
			} else {
				foundEdges++
				break
			}

			fallthrough
		default:
			other := images[actual.eID]
			if !placed[other.id] {
				space.Put(start.spaceX+1, start.spaceY, other)
				if _, ok := recentHelper[other.id]; !ok {
					recentlyPutIntoSpace = append(recentlyPutIntoSpace, other)
					recentHelper[other.id] = true
				}
				counter := 0
				for other.w != actual.e {
					switch counter % 5 {
					case 0:
						other.Rotate(false)
					case 1, 2:
						other.HorizontalFlip(false)
					case 3, 4:
						other.VerticalFlip(false)
					}
					counter++
				}

				other.wID = actual.id
			}
		}
		switch actual.sID {
		case 0:
			rev := reverseString(actual.s)
			if img, ok := storage[rev]; ok && img != actual {
				actual.sID = img.id
			} else if img, ok := storage[actual.s]; ok && img != actual {
				actual.sID = img.id
			} else {
				foundEdges++
				break
			}

			fallthrough
		default:
			other := images[actual.sID]
			if !placed[other.id] {
				space.Put(start.spaceX, start.spaceY+1, other)
				if _, ok := recentHelper[other.id]; !ok {
					recentlyPutIntoSpace = append(recentlyPutIntoSpace, other)
					recentHelper[other.id] = true
				}
				counter := 0
				for other.n != actual.s {
					switch counter % 5 {
					case 0:
						other.Rotate(false)
					case 1, 2:
						other.HorizontalFlip(false)
					case 3, 4:
						other.VerticalFlip(false)
					}
					counter++
				}

				other.nID = actual.id
			}
		}
		switch actual.wID {
		case 0:
			rev := reverseString(actual.w)
			if img, ok := storage[rev]; ok && img != actual {
				actual.wID = img.id
			} else if img, ok := storage[actual.w]; ok && img != actual {
				actual.wID = img.id
			} else {
				foundEdges++
				break
			}

			fallthrough
		default:
			other := images[actual.wID]
			if !placed[other.id] {
				space.Put(start.spaceX-1, start.spaceY, other)
				if _, ok := recentHelper[other.id]; !ok {
					recentlyPutIntoSpace = append(recentlyPutIntoSpace, other)
					recentHelper[other.id] = true
				}
				counter := 0
				for other.e != actual.w {
					switch counter % 5 {
					case 0:
						other.Rotate(false)
					case 1, 2:
						other.HorizontalFlip(false)
					case 3, 4:
						other.VerticalFlip(false)
					}
					counter++
				}

				other.eID = actual.id
			}
		}
	}

	return space
}

func findCompletedSegment(images map[int]*Image) *Image {
	for _, img := range images {
		if img.IsDone() {
			return img
		}
	}

	panic("This is awkward. :|")
}
