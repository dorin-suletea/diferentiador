package diff

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func NewEmptyDiffWidget() *widget.TextGrid {
	grid := widget.NewTextGridFromString("")
	grid.ShowLineNumbers = true
	grid.ShowWhitespace = true
	return grid
}

func SetDiffContent(diffContent string, mutableDiffWidget *widget.TextGrid) {
	mutableDiffWidget.SetText(diffContent)
	appyStlingForGrid(mutableDiffWidget, diffContent)
}

func appyStlingForGrid(mutableGrid *widget.TextGrid, content string) {
	rowAdded := widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}}
	rowRemoved := widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 192, G: 64, B: 64, A: 128}}

	for index, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {
		if line[0] == '-' {
			mutableGrid.SetRowStyle(index, &rowRemoved)
		}
		if line[0] == '+' {
			mutableGrid.SetRowStyle(index, &rowAdded)
		}
	}
}
