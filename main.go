package main

//sudo apt-get install libgl1-mesa-dev xorg-dev
// DEMO code
//https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
//go run fyne.io/fyne/v2/cmd/fyne_demo@latest

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/dorin-suletea/diferentiador~/internal/diff"
	"github.com/dorin-suletea/diferentiador~/internal/status"
)

type myTheme struct{}

// TODO : TextGrid is terribly slow as it can style each character
// create a custom one with styling per row.
func main() {
	app := app.NewWithID("diferentiador")
	app.SetIcon(theme.FyneLogo())
	window := app.NewWindow("diferentiador")

	fileStatus := status.GetStatusForFiles()
	diffWidget := diff.NewEmptyDiffWidget()
	scrollableDiffWidget := container.NewVScroll(diffWidget)

	selectionHandler := func(selectedFile string) {
		diffContent := diff.GetDiffForFile(selectedFile)
		diff.SetDiffContent(diffContent, diffWidget)
	}

	statusWidget := status.NewFilesStatusWidget(fileStatus, selectionHandler)

	split := container.NewHSplit(statusWidget, scrollableDiffWidget)
	split.Offset = 0.2

	window.SetContent(split)

	window.Resize(fyne.NewSize(1040, 460))
	window.ShowAndRun()
}
