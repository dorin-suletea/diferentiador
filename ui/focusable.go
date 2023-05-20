package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// borderBox := container.NewBorder(canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), box)

type Focusable struct {
	hasFocus  bool
	focused   fyne.CanvasObject
	unfocused fyne.CanvasObject
}

func NewFocusable(nested fyne.CanvasObject) *Focusable {
	focused := container.NewBorder(NewBorderLine(), NewBorderLine(), NewBorderLine(), NewBorderLine(), nested)
	return &Focusable{false, focused, nested}
}

func (f *Focusable) DrawableComponent() *fyne.CanvasObject {
	if f.hasFocus {
		return &f.focused
	} else {
		return &f.unfocused
	}
}

func (f *Focusable) Forcus(globalFocusables []*Focusable) {
	print(globalFocusables)
	for _, other := range globalFocusables {
		other.hasFocus = false
	}
	f.hasFocus = true
}
