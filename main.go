package main

//sudo apt-get install libgl1-mesa-dev xorg-dev
// DEMO code
//https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
//go run fyne.io/fyne/v2/cmd/fyne_demo@latest

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dorin-suletea/diferentiador~/internal/changelog"
)

type myTheme struct{}

func main() {
	app := app.NewWithID("diferentiador")
	app.SetIcon(theme.FyneLogo())
	window := app.NewWindow("diferentiador")
	changedContent := widget.NewTextGridFromString("potato poteason")

	mods := changelog.GetGitModififications()
	modsWidget := changelog.NewModifiedItemsWidget(mods)

	split := container.NewHSplit(modsWidget, changedContent)
	split.Offset = 0.2

	window.SetContent(split)

	window.Resize(fyne.NewSize(1040, 460))
	window.ShowAndRun()
}

//git diff --no-prefix -U1000 main.go
func main3() {
	app := app.NewWithID("diferentiador")
	// app.Settings().SetTheme(ui.NewCustomTheme())
	app.SetIcon(theme.FyneLogo())
	window := app.NewWindow("diferentiador")
	changedContent := widget.NewTextGridFromString("potato poteason")

	var data = []string{"a", "string", "list"}

	iconFile, err := os.Open("/home/dsu/Workspace/diferentiador/res/red_a.png")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(iconFile)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	name := widget.NewLabel("ssssss")
	label := widget.NewIcon(fyne.NewStaticResource("icon", b))
	button := widget.NewButtonWithIcon(">", theme.ConfirmIcon(), func() { fmt.Println("tapped text & leading icon button") })
	// button.Resize(fyne.NewSize(150, 30))
	// label.Resize(fyne.NewSize(150, 30))

	// grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(40, 40)),
	// 	label, widget.NewLabel(""), widget.NewLabel(""), button)
	// grid := container.New(layout.NewGridLayout(2),
	// 	label, button)
	// grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 40)),
	// 	container.New(layout.NewHBoxLayout(), label, layout.NewSpacer(), button),
	// 	container.New(layout.NewHBoxLayout(), label, name, layout.NewSpacer(), button))

	// row := container.NewBorder(nil, nil, label, button)

	// row.Resize(fyne.NewSize(150, 30))

	// list := widget.NewList(
	// 	func() int {
	// 		return len(data)
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("template")
	// 	},
	// 	func(i widget.ListItemID, o fyne.CanvasObject) {

	// 		o.(*widget.Label).SetText(data[i])
	// 	})
	// list.Select(1000)
	// changedFiles := container.NewBorder(nil, nil, nil, nil, widget.NewList(
	// 	func() int {
	// 		return len(data)
	// 	},
	// 	func() fyne.CanvasObject {
	// 		row := container.NewCenter(
	// 			widget.NewButton("", nil),
	// 			widget.NewLabel("Will be replaced"))
	// 		row.Resize(fyne.Size{row.Size().Width, 10})
	// 		return row
	// 	},
	// 	func(id widget.ListItemID, item fyne.CanvasObject) {
	// 		item.(*fyne.Container).Resize(fyne.Size{item.(*fyne.Container).Size().Width, 200})

	// 		item.(*fyne.Container).Objects[0].(*widget.Button).SetText(data[id])
	// 		item.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = func() {
	// 			fmt.Println("I am button " + data[id])
	// 			changedContent.SetText(data[id])
	// 		}
	// 	}))

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.New(layout.NewHBoxLayout(), label, name, layout.NewSpacer(), button)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// o.(*widget.Label).SetText(data[i])
		})

	// split := container.NewHSplit(grid, changedContent)
	split := container.NewHSplit(list, changedContent)
	split.Offset = 0.2

	window.SetContent(split)

	window.Resize(fyne.NewSize(1040, 460))
	window.ShowAndRun()
}
