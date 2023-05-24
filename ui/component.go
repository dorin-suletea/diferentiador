package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type IComponent interface {
	SetFocus(focus bool)
	GetDrawable() *fyne.CanvasObject
	OnArrowDown()
	OnArrowUp()
	OnArrowRight()
	OnArrowLeft()
}

type FocusComponent struct {
	hasFocus  bool
	focused   fyne.CanvasObject
	unfocused fyne.CanvasObject
}

func NewFocusComponent(nested fyne.CanvasObject) *FocusComponent {
	focused := container.NewBorder(NewFocusLine(), NewFocusLine(), NewFocusLine(), NewFocusLine(), nested)
	unfocused := container.NewBorder(NewUnfocusLine(), NewUnfocusLine(), NewUnfocusLine(), NewUnfocusLine(), nested)
	return &FocusComponent{false, focused, unfocused}
}

func (f *FocusComponent) GetDrawable() *fyne.CanvasObject {
	if f.hasFocus {
		return &f.focused
	} else {
		return &f.unfocused
	}
}

func (f *FocusComponent) SetFocus(focus bool) {
	f.hasFocus = focus
}

type StatusWidget struct {
	*FocusComponent
}

func NewStatusWidget(nested *container.Scroll) *StatusWidget {
	return &StatusWidget{NewFocusComponent(nested)}
}

func (dw *StatusWidget) OnArrowDown() {
	fmt.Println("Down")
}

func (dw *StatusWidget) OnArrowUp() {
	fmt.Println("Up")
}

func (dw *StatusWidget) OnArrowRight() {
	fmt.Println("right")
}

func (dw *StatusWidget) OnArrowLeft() {
	fmt.Println("left")
}

// ----------------------
// Main application loop
// ----------------------
type App struct {
	components   []IComponent
	window       fyne.Window
	currentFocus int
}

func NewApp() *App {
	topLevelComponents := []IComponent{}
	app := app.NewWithID("xdiff")
	app.SetIcon(theme.FyneLogo())
	app.Settings().SetTheme(theme.DefaultTheme())

	window := app.NewWindow("xdiff")
	window.Resize(fyne.NewSize(1920, 1080))
	return &App{topLevelComponents, window, 0}
}

func (a *App) AddComponent(c IComponent) IComponent {
	a.components = append(a.components, c)
	// focus the first added component on startup
	if len(a.components) == 1 {
		a.components[0].SetFocus(true)
	}
	a.window.Canvas().FocusPrevious()
	return a.components[len(a.components)-1]
}

func (a *App) ShowAndRun() {
	a.Refresh()
	a.window.ShowAndRun()
}

func (a *App) CycleFocus() {
	for _, globalComonents := range a.components {
		globalComonents.SetFocus(false)
	}

	a.currentFocus = (a.currentFocus + 1) % len(a.components)
	a.components[a.currentFocus].SetFocus(true)
	a.Refresh()
}

func (a *App) Refresh() {
	split := container.NewHSplit(*a.components[0].GetDrawable(), *a.components[1].GetDrawable())
	split.Offset = SplitOffset

	a.window.SetContent(split)
	a.window.Canvas().Refresh(split)
}

func (a *App) AddShortcut(shortcut fyne.Shortcut, handle func()) {
	a.window.Canvas().AddShortcut(shortcut, func(shortcut fyne.Shortcut) {
		handle()
	})
}

func (a *App) OnArrowDown() {
	a.components[a.currentFocus].OnArrowDown()
}

func (a *App) OnArrowUp() {
	a.components[a.currentFocus].OnArrowUp()
}

func (a *App) OnArrowRight() {
	a.components[a.currentFocus].OnArrowRight()
}

func (a *App) OnArrowLeft() {
	a.components[a.currentFocus].OnArrowLeft()
}
