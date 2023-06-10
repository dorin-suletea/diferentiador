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

func newFilesStatusList(data []FileStatus) *widget.List {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			placeholderStatusIcon := widget.NewIcon(theme.ConfirmIcon())
			placeholderFileName := widget.NewLabel("")
			placeholderFileName.TextStyle = ui.FontMono

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
	list := newFilesStatusList(data)

	scroll := container.NewScroll(list)

	ret := &StatusWidget{ui.NewFocusComponent(scroll), scroll, list, 0, data, onSelected}
	if len(data) != 0 {
		ret.selectItem(0)
	}

	ret.setSelectionHandlers()
	return ret
}

// TODO : panic here on adding new file and selecting it.
// github.com/dorin-suletea/diferentiador~/internal/status.(*StatusWidget).setSelectionHandlers.func1(0x997bee?)
//         /home/dsu/Workspace/diferentiador/internal/status/ui.go:101 +0xa6
// fyne.io/fyne/v2/widget.(*List).Select.func1()
//         /home/dsu/Tools/go/pkg/mod/fyne.io/fyne/v2@v2.3.4/widget/list.go:178 +0x5c
// fyne.io/fyne/v2/widget.(*List).Select(0xc0000eec60, 0x5)
//         /home/dsu/Tools/go/pkg/mod/fyne.io/fyne/v2@v2.3.4/widget/list.go:183 +0x178
// fyne.io/fyne/v2/widget.(*listLayout).setupListItem.func1()
//         /home/dsu/Tools/go/pkg/mod/fyne.io/fyne/v2@v2.3.4/widget/list.go:583 +0x25
// fyne.io/fyne/v2/widget.(*listItem).Tapped(0xc0004467e0, 0xc000460001)
//         /home/dsu/Tools/go/pkg/mod/fyne.io/fyne/v2@v2.3.4/widget/list.go:432 +0x37
// fyne.io/fyne/v2/internal/driver/glfw.(*window).mouseClickedHandleTapDoubleTap.func1()
//         /home/dsu/Tools/go/pkg/mod/fyne.io/fyne/v2@v2.3.4/internal/driver/glfw/window.go:634 +0x26
// fyne.io/fyne/v2/internal/driver/common.(*Window).RunEventQueue(0x0?)
//         /home/dsu/Tools/go/pkg/mod/fyne.io/fyne/v2@v2.3.4/internal/driver/common/window.go:35 +0x3e
// created by fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).createWindow.func1
//         /home/dsu/Tools/go/pkg/mod/fyne.io/fyne/v2@v2.3.4/internal/driver/glfw/window.go:946 +0x136
// exit status 2

func (dw *StatusWidget) setSelectionHandlers() {
	dw.nestedList.OnSelected = func(i widget.ListItemID) {
		dw.onSelected(dw.data[i])
		dw.selectionIndex = i
	}
}

func (dw *StatusWidget) SetContent(newData []FileStatus) {
	dw.nestedList = newFilesStatusList(newData)
	dw.setSelectionHandlers()
	dw.scroll.Content = dw.nestedList
	dw.scroll.Refresh()

	fmt.Println("setting status content")
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
	//no-op
}

func (dw *StatusWidget) OnArrowLeft() {
	//no-op
}

func (dw *StatusWidget) selectItem(index int) {
	dw.onSelected(dw.data[index])
	dw.nestedList.Select(index)
}

func (dw *StatusWidget) GetSelected() FileStatus {
	return dw.data[dw.selectionIndex]
}
