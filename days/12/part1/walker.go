package main

import (
	"bytes"
	"fmt"
)

//Walker todo
type Walker struct {
	facing     rune
	posX, posY int
}

//CreateWalker todo
func CreateWalker(x, y int, facing rune) *Walker {
	return &Walker{posX: x, posY: y, facing: facing}
}

//Move todo
func (w *Walker) Move(direction rune, unit int) {
	if direction == 'F' {
		direction = w.facing
	}

	switch direction {
	case 'N':
		w.posY += unit
	case 'S':
		w.posY -= unit
	case 'E':
		w.posX += unit
	case 'W':
		w.posX -= unit
	}
}

//Turn todo
func (w *Walker) Turn(direciton rune, unit int) {
	steps := unit / 90
	if direciton == 'L' {
		steps *= -1
	}

	direcitons := []byte{'N', 'E', 'S', 'W'}
	facing := bytes.IndexRune(direcitons, w.facing)
	facing += steps + 4
	facing %= 4

	w.facing = rune(direcitons[facing])
}

//Position todo
func (w *Walker) Position() (int, int) {
	return w.posX, w.posY
}

//Status todo
func (w *Walker) Status() {
	fmt.Println(w.posX, w.posY, string(w.facing))
}
