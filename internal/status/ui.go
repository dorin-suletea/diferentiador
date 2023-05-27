package status

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dorin-suletea/diferentiador~/ui"
)

type SelectionHandler func(FileStatus)

func newFilesStatusList(data []FileStatus, onSelected SelectionHandler) *widget.List {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			placeholderStatusIcon := widget.NewIcon(theme.ConfirmIcon())
			placeholderFileName := widget.NewLabel("")
			placeholderFileName.TextStyle = ui.FontMono

			return container.New(layout.NewHBoxLayout(), placeholderStatusIcon, placeholderFileName, layout.NewSpacer())
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			container := o.(*fyne.Container)
			item := data[i]
			// icon
			resource, err := loadIcon(pickIconPath(item))
			if err != nil {
				fmt.Println(err)
				resource = theme.ContentRemoveIcon()
			}
			(container.Objects[0].(*widget.Icon)).SetResource(resource)
			(container.Objects[1].(*widget.Label)).SetText(data[i].FilePath)
		})
	return list
}

func loadIcon(relativePath string) (fyne.Resource, error) {
	res, err := fyne.LoadResourceFromPath(relativePath)
	return res, err
}

func pickIconPath(item FileStatus) string {
	if item.Status == Added {
		return "res/green_plus.png"
	}
	if item.Status == Deleted && item.staged {
		return "res/green_minus.png"
	}
	if item.Status == Deleted && !item.staged {
		return "res/red_minus.png"
	}
	if item.Status == Modified && item.staged {
		return "res/green_m.png"
	}
	if item.Status == Modified && !item.staged {
		return "res/red_m.png"
	}
	if item.Status == Untracked {
		return "res/red_question.png"
	}
	return ""
}

// ----------------------
// StatusWidget
// ----------------------
type StatusWidget struct {
	*ui.FocusComponent
	scroll         *container.Scroll
	nestedList     *widget.List
	selectionIndex int
	data           []FileStatus
	onSelected     SelectionHandler
}

func NewStatusWidget(data []FileStatus, onSelected SelectionHandler) *StatusWidget {
	list := newFilesStatusList(data, onSelected)

	scroll := container.NewScroll(list)

	ret := &StatusWidget{ui.NewFocusComponent(scroll), scroll, list, 0, data, onSelected}
	if len(data) != 0 {
		ret.selectItem(0)
	}

	// update selectionIndex on click
	list.OnSelected = func(i widget.ListItemID) {
		onSelected(data[i])
		ret.selectionIndex = i
	}

	return ret
}

func (dw *StatusWidget) InitHandlers() {

}

func (dw *StatusWidget) OnArrowDown() {
	dw.selectionIndex = (dw.selectionIndex + 1) % len(dw.data)
	dw.selectItem(dw.selectionIndex)
}

func (dw *StatusWidget) OnArrowUp() {
	newIndex := dw.selectionIndex - 1

	if newIndex >= 0 {
		dw.selectionIndex = newIndex
	} else {
		dw.selectionIndex = len(dw.data) + newIndex
	}
	dw.selectItem(dw.selectionIndex)
}

func (dw *StatusWidget) OnArrowRight() {
	fmt.Println("right")
}

func (dw *StatusWidget) OnArrowLeft() {
	fmt.Println("left")
}

func (dw *StatusWidget) selectItem(index int) {
	dw.onSelected(dw.data[index])
	dw.nestedList.Select(index)
}
