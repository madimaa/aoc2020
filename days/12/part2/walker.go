package main

import (
	"fmt"
)

//Walker walker struct
type Walker struct {
	waypointX, waypointY int
	posX, posY           int
}

//CreateWalker create the walker objest
func CreateWalker(x, y, waypointX, waypointY int) *Walker {
	return &Walker{posX: x, posY: y, waypointX: waypointX, waypointY: waypointY}
}

//Move move the waypoint or the ship
func (w *Walker) Move(direction rune, unit int) {
	switch direction {
	case 'N':
		w.waypointY += unit
	case 'S':
		w.waypointY -= unit
	case 'E':
		w.waypointX += unit
	case 'W':
		w.waypointX -= unit
	case 'F':
		w.posX += unit * w.waypointX
		w.posY += unit * w.waypointY
	}
}

//Turn turn the waypoint around the ship
func (w *Walker) Turn(direciton rune, unit int) {
	steps := unit / 90
	wayX := w.waypointX
	wayY := w.waypointY
	if direciton == 'R' {
		for steps > 0 {
			temp := wayY
			wayY = wayX * -1
			wayX = temp
			steps--
		}
	} else if direciton == 'L' {
		for steps > 0 {
			temp := wayX
			wayX = wayY * -1
			wayY = temp
			steps--
		}
	}

	w.waypointX = wayX
	w.waypointY = wayY
}

//Position ships position
func (w *Walker) Position() (int, int) {
	return w.posX, w.posY
}

//Status print ship status
func (w *Walker) Status() {
	fmt.Println(w.posX, w.posY, w.waypointX, w.waypointY)
}
