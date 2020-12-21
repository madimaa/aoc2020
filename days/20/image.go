package main

import "fmt"

//Image struct
type Image struct {
	id                 int
	n, e, s, w         string
	nID, eID, sID, wID int
	spaceX, spaceY     int
	fullImage          []string
}

//CreateImage creates image struct
func CreateImage(id int, input []string, withFullImage bool) *Image {
	e, w := "", ""
	n := input[0]
	s := input[len(input)-1]
	lastCharIndex := len(n) - 1
	fullImage := make([]string, 0)
	for _, row := range input {
		if withFullImage {
			fullImage = append(fullImage, row[1:len(row)-1])
		}

		w += string([]rune(row)[0])
		e += string([]rune(row)[lastCharIndex])
	}

	if withFullImage {
		fullImage = fullImage[1 : len(fullImage)-1]
	}

	return &Image{id: id, n: n, e: e, s: s, w: w, fullImage: fullImage}
}

//Print console friendly output
func (i *Image) Print() {
	fmt.Println("id:", i.id, i.IsDone(), "n", i.n, "e", i.e, "s", i.s, "w", i.w, "nID", i.nID, "eID", i.eID, "sID", i.sID, "wID", i.wID)
}

//IsDone check finished tile
func (i *Image) IsDone() bool {
	return i.nID != 0 && i.eID != 0 && i.sID != 0 && i.wID != 0
}

//IsCorner check tile is a corner
func (i *Image) IsCorner() bool {
	corners := 0
	if i.nID != 0 {
		corners++
	}

	if i.eID != 0 {
		corners++
	}

	if i.sID != 0 {
		corners++
	}

	if i.wID != 0 {
		corners++
	}

	return corners == 2
}

//Rotate rotate the image
func (i *Image) Rotate(withFullImage bool) {
	temp := i.w
	tempID := i.wID

	i.w = i.s
	i.wID = i.sID

	i.s = reverseString(i.e)
	i.sID = i.eID

	i.e = i.n
	i.eID = i.nID

	i.n = reverseString(temp)
	i.nID = tempID

	if withFullImage {
		temp := make([]string, len(i.fullImage))
		for _, row := range i.fullImage {
			for runeIndex, r := range []rune(row) {
				temp[runeIndex] = string(r) + temp[runeIndex]
			}
		}

		i.fullImage = temp
	}
}

//VerticalFlip vertically flip the image
func (i *Image) VerticalFlip(withFullImage bool) {
	temp := i.s
	tempID := i.sID

	i.s = i.n
	i.sID = i.nID

	i.n = temp
	i.nID = tempID

	i.e = reverseString(i.e)
	i.w = reverseString(i.w)

	if withFullImage {
		length := len(i.fullImage)
		temp := make([]string, length)
		for index := 0; index < length/2; index++ {
			temp[index] = i.fullImage[length-1-index]
			temp[length-1-index] = i.fullImage[index]
		}
	}
}

//HorizontalFlip horizontally flip the image
func (i *Image) HorizontalFlip(withFullImage bool) {
	temp := i.w
	tempID := i.wID

	i.w = i.e
	i.wID = i.eID

	i.e = temp
	i.eID = tempID

	i.n = reverseString(i.n)
	i.s = reverseString(i.s)

	if withFullImage {
		for index, row := range i.fullImage {
			i.fullImage[index] = reverseString(row)
		}
	}
}

func reverseString(in string) string {
	reverse := ""

	runes := []rune(in)
	for i := len(in) - 1; i >= 0; i-- {
		reverse += string(runes[i])
	}

	return reverse
}
