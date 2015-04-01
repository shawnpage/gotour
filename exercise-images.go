//https://gist.github.com/zyxar/2317744#file-exercise-tour-go-L154

package main

import (
	"golang.org/x/tour/pic"	
	"image"
	"image/color"
)

type Image struct{
	Width, Height int
	colr uint8
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0,0, i.Width, i.Height)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	
	return color.RGBA{i.colr+uint8(x), i.colr+uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 128}

	pic.ShowImage(&m)
}
