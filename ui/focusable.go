package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// borderBox := container.NewBorder(canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), box)

type Focusable struct {
	HasFocus  bool
	focused   fyne.CanvasObject
	unfocused fyne.CanvasObject
}

func NewFocusable(nested fyne.CanvasObject) *Focusable {
	focused := container.NewBorder(NewFocusLine(), NewFocusLine(), NewFocusLine(), NewFocusLine(), nested)
	unfocused := container.NewBorder(NewUnfocusLine(), NewUnfocusLine(), NewUnfocusLine(), NewUnfocusLine(), nested)
	return &Focusable{false, focused, unfocused}
}

func (f *Focusable) DrawableComponent() *fyne.CanvasObject {
	if f.HasFocus {
		return &f.focused
	} else {
		return &f.unfocused
	}
}

func (f *Focusable) Forcus(globalFocusables []*Focusable) {
	for _, other := range globalFocusables {
		other.HasFocus = false
	}
	f.HasFocus = true
}
