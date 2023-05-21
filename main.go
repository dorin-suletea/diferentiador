package main

//sudo apt-get install libgl1-mesa-dev xorg-dev
// DEMO code
//https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
//go run fyne.io/fyne/v2/cmd/fyne_demo@latest

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/dorin-suletea/diferentiador~/internal/diff"
	"github.com/dorin-suletea/diferentiador~/internal/status"
	"github.com/dorin-suletea/diferentiador~/ui"
)

var focus []*ui.Focusable = []*ui.Focusable{}
var currentFocusIndex = 0

// make([]*ui.Focusable, 0)

func main() {
	app := app.NewWithID("xdiff")
	app.SetIcon(theme.FyneLogo())
	// app.Settings().SetTheme(ui.NewCustomTheme())
	app.Settings().SetTheme(theme.DefaultTheme())

	window := app.NewWindow("xdiff")
	// TODO : open maximised
	window.Resize(fyne.NewSize(1920, 1080))

	// TODO : start a cron to eagerly pre-create all these widgets
	fileStatusContent := status.GetStatusForFiles()
	diffWidget := container.NewScroll(container.NewVBox())

	genericSelectionHandler := func(content string, scrollContainer *container.Scroll) {
		contentBox := diff.NewDiffWidget(content)
		diffWidget.Content = contentBox

		diffWidget.Refresh()
		contentBox.Refresh()
	}

	onMutatedHandler := func(selectedFile string) {
		content := diff.GetDiffForFile(selectedFile)
		genericSelectionHandler(content, diffWidget)
	}
	onDeletedHandler := func(selectedFile string) {
		content := diff.MarkLines(diff.GetHeadForFile(selectedFile), '-')
		genericSelectionHandler(content, diffWidget)
	}
	onUntrackedHandler := func(selectedFile string) {
		content := diff.GetRawFileContents(selectedFile)
		genericSelectionHandler(content, diffWidget)
	}

	fileStatusWidget := status.NewFilesStatusWidget(fileStatusContent, onMutatedHandler, onDeletedHandler, onUntrackedHandler)
	scrollableFileStatusWidget := container.NewScroll(fileStatusWidget)
	// Auto-select first file
	if len(fileStatusContent) != 0 {
		status.HandleSelection(fileStatusContent[0], onMutatedHandler, onDeletedHandler, onUntrackedHandler)
	}

	fStatus := ui.NewFocusable(scrollableFileStatusWidget)
	fDiff := ui.NewFocusable(diffWidget)
	focus = append(focus, fStatus, fDiff)

	focus[currentFocusIndex].Forcus(focus)
	setupSplit(window)

	initShortcuts(window)
	// diffWidget.ScrollToBottom()

	window.ShowAndRun()
}

func initShortcuts(w fyne.Window) {
	w.Canvas().AddShortcut(ui.ShCycleFocus, func(shortcut fyne.Shortcut) {
		currentFocusIndex = (currentFocusIndex + 1) % len(focus)
		focus[currentFocusIndex].Forcus(focus)
		setupSplit(w)

		fmt.Println(focus[currentFocusIndex].HasFocus)
	})

	w.Canvas().AddShortcut(ui.ShScrollDown, func(shortcut fyne.Shortcut) {
		// diffWidget.Offset.Y = 100
	})

}

func setupSplit(w fyne.Window) {
	split := container.NewHSplit(*focus[0].DrawableComponent(), *focus[1].DrawableComponent())
	split.Offset = ui.SplitOffset
	w.SetContent(split)
	w.Canvas().Refresh(split)
}
