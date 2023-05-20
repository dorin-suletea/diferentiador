package status

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SelectionHandler func(string)

func NewFilesStatusWidget(data []FileStatus, onSelectMutated SelectionHandler, onSelectDeleted SelectionHandler, onSelectUntracked SelectionHandler) *widget.List {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			placeholderStatusIcon := widget.NewIcon(theme.ConfirmIcon())
			placeholderFileName := widget.NewLabel("")
			testButton := widget.NewButton("Select", func() {})
			return container.New(layout.NewHBoxLayout(), placeholderStatusIcon, placeholderFileName, layout.NewSpacer(), testButton)
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
			// file path
			(container.Objects[1].(*widget.Label)).SetText(data[i].fileName)
			(container.Objects[3].(*widget.Button)).OnTapped = func() {
				switch data[i].status {
				case Deleted:
					onSelectDeleted(data[i].fileName)
				case Untracked:
					onSelectUntracked(data[i].fileName)
				default:
					onSelectMutated(data[i].fileName)
				}
			}
		})
	return list
}

func loadIcon(relativePath string) (fyne.Resource, error) {
	res, err := fyne.LoadResourceFromPath(relativePath)
	return res, err
}

func pickIconPath(item FileStatus) string {
	if item.status == Added {
		return "res/green_plus.png"
	}
	if item.status == Deleted && item.staged {
		return "res/green_minus.png"
	}
	if item.status == Deleted && !item.staged {
		return "res/red_minus.png"
	}
	if item.status == Modified && item.staged {
		return "res/green_m.png"
	}
	if item.status == Modified && !item.staged {
		return "res/red_m.png"
	}
	if item.status == Untracked {
		return "res/red_question.png"
	}
	return ""
}
