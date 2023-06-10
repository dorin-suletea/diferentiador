package main

//sudo apt-get install libgl1-mesa-dev xorg-dev
// DEMO code
//https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
//go run fyne.io/fyne/v2/cmd/fyne_demo@latest

import (
	"fyne.io/fyne/v2"
	"github.com/dorin-suletea/diferentiador~/internal/diff"
	"github.com/dorin-suletea/diferentiador~/internal/status"
	"github.com/dorin-suletea/diferentiador~/ui"
)

//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

const refreshRateSeconds int = 10

func main() {
	app := ui.NewApp()

	// diff widget, center
	diffWidget := diff.NewDiffWidget([]fyne.CanvasObject{})
	statusCache := status.NewGitDiffCache(refreshRateSeconds)

	diffCache := diff.NewGitDiffCache(statusCache.GetChangedFiles(), refreshRateSeconds)

	selectionHandler := func(f status.FileStatus) {
		//TODO : must receive params, this is not safe
		content := diffCache.GetContent(f)
		diffWidget.SetContent(content)
	}

	statusWidget := status.NewStatusWidget(statusCache.GetChangedFiles(), selectionHandler)

	// Update the active diff windown whenever the diff cache is refreshed s
	diffCache.SetOnRefreshHandler(func() {
		selection := statusWidget.GetSelected()
		diffContent := diffCache.GetContent(selection)
		diffWidget.SetContent(diffContent)
	})
	statusCache.SetOnRefreshHandler(func() {
		statusWidget.SetContent(statusCache.GetChangedFiles())
	})

	app.AddComponent(statusWidget)
	app.AddComponent(diffWidget)

	app.AddShortcut(ui.ShCycleFocus, func() { app.CycleFocus() })
	app.AddShortcut(ui.ShArrowDown, func() { app.OnArrowDown() })
	app.AddShortcut(ui.ShArrowUp, func() { app.OnArrowUp() })
	app.AddShortcut(ui.ShArrowRight, func() { app.OnArrowRight() })
	app.AddShortcut(ui.ShArrowLeft, func() { app.OnArrowLeft() })
	app.ShowAndRun()
}

// var focus []*ui.Focusable = []*ui.Focusable{}
// var currentFocusIndex = 0
// func main2() {
// 	app := app.NewWithID("xdiff")
// 	app.SetIcon(theme.FyneLogo())
// 	// app.Settings().SetTheme(ui.NewCustomTheme())
// 	app.Settings().SetTheme(theme.DefaultTheme())

// 	window := app.NewWindow("xdiff")
// 	// TODO : open maximised
// 	window.Resize(fyne.NewSize(1920, 1080))

// 	// TODO : start a cron to eagerly pre-create all these widgets
// 	fileStatusContent := status.GetStatusForFiles()
// 	diffWidget := container.NewScroll(container.NewVBox())

// 	genericSelectionHandler := func(content string, scrollContainer *container.Scroll) {
// 		contentBox := diff.NewDiffWidget(content)
// 		diffWidget.Content = contentBox

// 		diffWidget.Refresh()
// 		contentBox.Refresh()
// 	}

// 	onMutatedHandler := func(selectedFile string) {
// 		content := diff.GetDiffForFile(selectedFile)
// 		genericSelectionHandler(content, diffWidget)
// 	}
// 	onDeletedHandler := func(selectedFile string) {
// 		content := diff.MarkLines(diff.GetHeadForFile(selectedFile), '-')
// 		genericSelectionHandler(content, diffWidget)
// 	}
// 	onUntrackedHandler := func(selectedFile string) {
// 		content := diff.GetRawFileContents(selectedFile)
// 		genericSelectionHandler(content, diffWidget)
// 	}

// 	// fileStatusWidget := status.NewFilesStatusWidget(fileStatusContent, onMutatedHandler, onDeletedHandler, onUntrackedHandler)
// 	// scrollableFileStatusWidget := container.NewScroll(fileStatusWidget)
// 	// Auto-select first file
// 	if len(fileStatusContent) != 0 {
// 		status.HandleSelection(fileStatusContent[0], onMutatedHandler, onDeletedHandler, onUntrackedHandler)
// 	}

// 	// fStatus := ui.NewFocusable(scrollableFileStatusWidget)
// 	// fDiff := ui.NewFocusable(diffWidget)
// 	// focus = append(focus, fStatus, fDiff)

// 	focus[currentFocusIndex].Forcus(focus)
// 	setupSplit(window)

// 	initShortcuts(window)
// 	// diffWidget.ScrollToBottom()

// 	window.ShowAndRun()
// }

// func initShortcuts(w fyne.Window) {
// 	w.Canvas().AddShortcut(ui.ShCycleFocus, func(shortcut fyne.Shortcut) {
// 		currentFocusIndex = (currentFocusIndex + 1) % len(focus)
// 		focus[currentFocusIndex].Forcus(focus)
// 		setupSplit(w)

// 		fmt.Println(focus[currentFocusIndex].HasFocus)
// 	})

// 	w.Canvas().AddShortcut(ui.ShScrollDown, func(shortcut fyne.Shortcut) {
// 		// diffWidget.Offset.Y = 100
// 	})

// }

// func setupSplit(w fyne.Window) {
// 	split := container.NewHSplit(*focus[0].DrawableComponent(), *focus[1].DrawableComponent())
// 	split.Offset = ui.SplitOffset
// 	w.SetContent(split)
// 	w.Canvas().Refresh(split)
// }
