package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dorin-suletea/diferentiador~/internal"
)

type SelectionHandler func(FileStatus)

func newFilesStatusList(data []FileStatus) *widget.List {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			placeholderStatusIcon := widget.NewIcon(theme.ConfirmIcon())
			placeholderFileName := widget.NewLabel("")
			placeholderFileName.TextStyle = internal.FontMono

			return container.New(layout.NewHBoxLayout(), widget.NewLabel(" "), placeholderStatusIcon, placeholderFileName, layout.NewSpacer())
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			container := o.(*fyne.Container)
			item := data[i]
			// icon
			resource, err := loadIcon(pickIconPath(item))
			if err != nil {
				resource = theme.ContentRemoveIcon()
			}
			(container.Objects[1].(*widget.Icon)).SetResource(resource)
			(container.Objects[2].(*widget.Label)).SetText(data[i].FilePath)
		})
	return list
}

func loadIcon(relativePath string) (fyne.Resource, error) {
	res, err := fyne.LoadResourceFromPath(relativePath)
	return res, err
}

// https://icons8.com/icon/set/plus/ios-filled
func pickIconPath(item FileStatus) string {
	if item.Status == Added {
		return "res/plus.png"
	}
	if item.Status == Deleted && item.staged {
		return "res/minus_green.png"
	}
	if item.Status == Deleted && !item.staged {
		return "res/minus_red.png"
	}
	if item.Status == Modified && item.staged {
		return "res/m_green.png"
	}
	if item.Status == Modified && !item.staged {
		return "res/m_red.png"
	}
	if item.Status == Untracked {
		return "res/questionmark.png"
	}
	if item.Status == Renamed {
		return "res/r.png"
	}
	return ""
}

// ----------------------
// ChangedFilesWidget
// ----------------------
type ChangedFilesWidget struct {
	*internal.FocusComponent
	scroll         *container.Scroll
	nestedList     *widget.List
	selectionIndex int
	data           *ChangedFileCache
	onSelected     SelectionHandler
}

func NewChangedFilesWidget(data *ChangedFileCache, onSelected SelectionHandler) *ChangedFilesWidget {
	list := newFilesStatusList(data.GetAll())

	scroll := container.NewScroll(list)

	ret := &ChangedFilesWidget{internal.NewFocusComponent(scroll), scroll, list, 0, data, onSelected}
	if data.Len() != 0 {
		ret.selectItem(0)
	}

	ret.setSelectionHandlers()
	return ret
}

func (dw *ChangedFilesWidget) setSelectionHandlers() {
	dw.nestedList.OnSelected = func(i widget.ListItemID) {
		dw.onSelected(dw.data.Get(i))
		dw.selectionIndex = i
	}
}

func (dw *ChangedFilesWidget) setContent(newData []FileStatus) {
	dw.nestedList = newFilesStatusList(newData)
	dw.scroll.Content = dw.nestedList
	dw.setSelectionHandlers()
	dw.selectItem(dw.selectionIndex)
	dw.scroll.Refresh()
}

func (dw *ChangedFilesWidget) OnArrowDown() {
	dw.selectionIndex = (dw.selectionIndex + 1) % dw.data.Len()
	dw.selectItem(dw.selectionIndex)
}

func (dw *ChangedFilesWidget) OnArrowUp() {
	newIndex := dw.selectionIndex - 1

	if newIndex >= 0 {
		dw.selectionIndex = newIndex
	} else {
		dw.selectionIndex = dw.data.Len() + newIndex
	}
	dw.selectItem(dw.selectionIndex)
}

func (dw *ChangedFilesWidget) OnArrowRight() {
	//no-op
}

func (dw *ChangedFilesWidget) OnArrowLeft() {
	//no-op
}

func (dw *ChangedFilesWidget) selectItem(i int) {
	dw.onSelected(dw.data.Get(i))
	dw.nestedList.Select(i)
}

func (dw *ChangedFilesWidget) GetSelected() FileStatus {
	return dw.data.Get(dw.selectionIndex)
}
