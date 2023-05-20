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

func main() {
	app := app.NewWithID("xdiff")
	app.SetIcon(theme.FyneLogo())
	window := app.NewWindow("xdiff")

	fileStatus := status.GetStatusForFiles()
	scrollableDiffWidget := container.NewScroll(container.NewVBox())
	onMutatedHandler := func(selectedFile string) {
		diffContent := diff.GetDiffForFile(selectedFile)
		scrollableDiffWidget.Content = diff.NewDiffWidget(diffContent)
		scrollableDiffWidget.Refresh()
	}
	onDeletedHandler := func(selectedFile string) {
		headContent := diff.GetHeadForFile(selectedFile)
		markedHeadContent := diff.MarkLines(headContent, '+')
		scrollableDiffWidget.Content = diff.NewDiffWidget(markedHeadContent)
		scrollableDiffWidget.Refresh()
	}
	onUntrackedHandler := func(selectedFile string) {
		diffContent := diff.GetRawFileContents(selectedFile)
		markedDiffContent := diff.MarkLines(diffContent, '+')
		scrollableDiffWidget.Content = diff.NewDiffWidget(markedDiffContent)
		scrollableDiffWidget.Refresh()
	}

	statusWidget := status.NewFilesStatusWidget(fileStatus, onMutatedHandler, onDeletedHandler, onUntrackedHandler)

	split := container.NewHSplit(statusWidget, scrollableDiffWidget)
	split.Offset = 0.2

	window.SetContent(split)

	window.Resize(fyne.NewSize(1040, 460))
	window.ShowAndRun()
}
