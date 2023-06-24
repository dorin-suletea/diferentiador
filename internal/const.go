package internal

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	Transparent = &color.Transparent
	Green       = &color.NRGBA{R: 64, G: 192, B: 64, A: 128}
	Red         = &color.NRGBA{R: 192, G: 64, B: 64, A: 128}
	Gray        = &color.NRGBA{R: 96, G: 96, B: 96, A: 255}
	FontMono    = fyne.TextStyle{Monospace: true}
	FontHeight  = float32(21)
	FontWidth   = float32(10)
	Focus       = Green
)

var SplitOffset = 0.2

func NewFocusLine() *canvas.Line {
	return &canvas.Line{
		StrokeColor: Green,
		StrokeWidth: 4,
	}
}

func NewUnfocusLine() *canvas.Line {
	return &canvas.Line{
		StrokeColor: Transparent,
		StrokeWidth: 4,
	}
}

var (
	ShCycleFocus = &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	ShArrowDown  = &desktop.CustomShortcut{KeyName: fyne.KeyDown, Modifier: fyne.KeyModifierControl}
	ShArrowUp    = &desktop.CustomShortcut{KeyName: fyne.KeyUp, Modifier: fyne.KeyModifierControl}
	ShArrowRight = &desktop.CustomShortcut{KeyName: fyne.KeyRight, Modifier: fyne.KeyModifierControl}
	ShArrowLeft  = &desktop.CustomShortcut{KeyName: fyne.KeyLeft, Modifier: fyne.KeyModifierControl}
)
