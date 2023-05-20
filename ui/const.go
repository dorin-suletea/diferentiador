package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var (
	Green    = &color.NRGBA{R: 64, G: 192, B: 64, A: 128}
	Red      = &color.NRGBA{R: 192, G: 64, B: 64, A: 128}
	Gray     = &color.NRGBA{R: 96, G: 96, B: 96, A: 255}
	FontMono = fyne.TextStyle{Monospace: true}
	Focus    = Green
)

func NewBorderLine() *canvas.Line {
	return &canvas.Line{
		StrokeColor: Green,
		StrokeWidth: 2,
	}
}
