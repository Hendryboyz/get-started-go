package main

import (
	"golang.org/x/tour/pic"
)

func calculateImage1(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func calculateImage2(x, y int) uint8 {
	return uint8(x * y)
}

func calculateImage3(x, y int) uint8 {
	return uint8(x ^ y)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = calculateImage1(x, y)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
