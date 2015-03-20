package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	outerslice := make([][] uint8, dy)
	for i := 0; i < dy; i++ {
		outerslice[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			outerslice[i][j] = uint8((i+j)/2)
		}
	}

	return outerslice
	
}

func main() {
	pic.Show(Pic)
}
