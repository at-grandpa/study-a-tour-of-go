package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for yi := range s {
		s[yi] = make([]uint8, dx)
		for xi := range s[yi] {
			s[yi][xi] = uint8((xi + yi) / 2)
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
