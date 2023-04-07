package diff

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func NewDiffForFileWidget(content string) *widget.TextGrid {
	// diff := workingtree.GetAlterationsForFile("main.go")
	// log.Println(diff)

	grid := widget.NewTextGridFromString(content)
	appyStlingForGrid(grid, content)
	// grid := widget.NewTextGridFromString(diff)
	// grid.SetStyleRange(0, 4, 0, 7, &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 64, B: 192, A: 128}})
	// grid.SetRowStyle(1, &rowAdded)
	// grid.SetRowStyle(3, &rowRemoved)

	// white := &widget.CustomTextGridStyle{FGColor: color.White, BGColor: color.Black}
	// black := &widget.CustomTextGridStyle{FGColor: color.Black, BGColor: color.White}
	// grid.Rows[2].Cells[0].Style = white
	// grid.Rows[2].Cells[1].Style = black
	// grid.Rows[2].Cells[2].Style = white
	// grid.Rows[2].Cells[3].Style = black
	// grid.Rows[2].Cells[4].Style = white

	grid.ShowLineNumbers = true
	grid.ShowWhitespace = true

	return grid
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

		fmt.Println(line)
	}
}
