package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	W int
	H int
}

func (m Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.W, m.H)
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func TryExerciseImage() {
	m := Image{250, 250}
	pic.ShowImage(m)
}
