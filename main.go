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
	"github.com/dorin-suletea/diferentiador~/ui"
)

var focus []*ui.Focusable = []*ui.Focusable{}

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
	fileStatus := status.GetStatusForFiles()
	scrollableDiffWidget := container.NewScroll(container.NewVBox())

	genericSelectionHandler := func(content string, scrollContainer *container.Scroll) {
		contentBox := diff.NewDiffWidget(content)
		scrollableDiffWidget.Content = contentBox

		scrollableDiffWidget.Refresh()
		contentBox.Refresh()
	}

	onMutatedHandler := func(selectedFile string) {
		content := diff.GetDiffForFile(selectedFile)
		genericSelectionHandler(content, scrollableDiffWidget)
	}
	onDeletedHandler := func(selectedFile string) {
		content := diff.MarkLines(diff.GetHeadForFile(selectedFile), '-')
		genericSelectionHandler(content, scrollableDiffWidget)
	}
	onUntrackedHandler := func(selectedFile string) {
		content := diff.GetRawFileContents(selectedFile)
		genericSelectionHandler(content, scrollableDiffWidget)
	}

	statusWidget := status.NewFilesStatusWidget(fileStatus, onMutatedHandler, onDeletedHandler, onUntrackedHandler)

	// Auto-select first file
	if len(fileStatus) != 0 {
		status.HandleSelection(fileStatus[0], onMutatedHandler, onDeletedHandler, onUntrackedHandler)
	}

	// TODO : figure out focus and tabbing betwen containers
	// borderedContent := container.NewBorder(ui.NewBorderLine(), ui.NewBorderLine(), ui.NewBorderLine(), ui.NewBorderLine(), scrollableDiffWidget)
	fStatus := ui.NewFocusable(statusWidget)
	fDiff := ui.NewFocusable(scrollableDiffWidget)
	focus = append(focus, fStatus, fDiff)

	fStatus.Forcus(focus)

	split := container.NewHSplit(*fStatus.DrawableComponent(), *fDiff.DrawableComponent())
	split.Offset = 0.2
	window.SetContent(split)

	// window.Canvas().Refresh(window.Content())
	window.ShowAndRun()
}
