package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// borderBox := container.NewBorder(canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), box)

type Focusable struct {
	hasFocus bool
	focused  *fyne.Container
	nested   *fyne.CanvasObject
}

func NewFocusable(nested *fyne.CanvasObject) *Focusable {
	focused := container.NewBorder(canvas.NewLine(Focus), canvas.NewLine(Focus), canvas.NewLine(Focus), canvas.NewLine(Focus), nested)
	return &Focusable{false, focused, nested}
}

// func (f *Focusable) DrawableComponent() {
// 	borderBox := container.NewBorder(canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), box)
// }

// func (f *Focusable) NestedComponent() {

// 	borderBox := container.NewBorder(canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), canvas.NewLine(ui.Green), box)
// }
