package main

import "fmt"

//Space to store images
type Space struct {
	s map[string]*Image
}

//CreateSpace to store images
func CreateSpace() *Space {
	return &Space{s: make(map[string]*Image)}
}

//Put into space
func (s *Space) Put(x, y int, img *Image) {
	img.spaceX = x
	img.spaceY = y
	s.s[convert(x, y)] = img
}

//Get from space
func (s *Space) Get(x, y int) *Image {
	return s.s[convert(x, y)]
}

//HasImageAt x y cords
func (s *Space) HasImageAt(x, y int) bool {
	if _, ok := s.s[convert(x, y)]; ok {
		return true
	}

	return false
}

func convert(x, y int) string {
	return fmt.Sprint(x, y)
}
