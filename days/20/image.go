package main

import "fmt"

//Image struct
type Image struct {
	id                 int
	n, e, s, w         string
	nID, eID, sID, wID int
	spaceX, spaceY     int
}

//CreateImage creates image struct
func CreateImage(id int, input []string) *Image {
	e, w := "", ""
	n := input[0]
	s := input[len(input)-1]
	lastCharIndex := len(n) - 1
	for _, row := range input {
		w += string([]rune(row)[0])
		e += string([]rune(row)[lastCharIndex])
	}

	return &Image{id: id, n: n, e: e, s: s, w: w}
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
func (i *Image) Rotate() {
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
}

//VerticalFlip vertically flip the image
func (i *Image) VerticalFlip() {
	temp := i.s
	tempID := i.sID

	i.s = i.n
	i.sID = i.nID

	i.n = temp
	i.nID = tempID

	i.e = reverseString(i.e)
	i.w = reverseString(i.w)
}

//HorizontalFlip horizontally flip the image
func (i *Image) HorizontalFlip() {
	temp := i.w
	tempID := i.wID

	i.w = i.e
	i.wID = i.eID

	i.e = temp
	i.eID = tempID

	i.n = reverseString(i.n)
	i.s = reverseString(i.s)
}

func reverseString(in string) string {
	reverse := ""

	runes := []rune(in)
	for i := len(in) - 1; i >= 0; i-- {
		reverse += string(runes[i])
	}

	return reverse
}
