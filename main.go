package main

//sudo apt-get install libgl1-mesa-dev xorg-dev
// DEMO code
//https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
//go run fyne.io/fyne/v2/cmd/fyne_demo@latest

import (
	"fmt"
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_demo/tutorials"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

//git diff --no-prefix -U1000 main.go
func main() {
	app := app.NewWithID("diferentiador")
	app.SetIcon(theme.FyneLogo())
	window := app.NewWindow("diferentiador")

	changedContent := widget.NewTextGridFromString("potato poteason")

	var data = []string{"a", "string", "list"}
	changedFiles := container.NewBorder(nil, nil, nil, nil, widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewPadded(
				widget.NewButton("", nil),
			)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Button).SetText(data[id])
			item.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = func() {
				fmt.Println("I am button " + data[id])
				changedContent.SetText(data[id])
			}
		}))

	split := container.NewHSplit(changedFiles, changedContent)
	split.Offset = 0.2

	window.SetContent(split)

	window.Resize(fyne.NewSize(1040, 460))
	window.ShowAndRun()
}

//
//
//
//
//

func main2() {
	a := app.NewWithID("io.fyne.demo")
	a.SetIcon(theme.FyneLogo())
	a.Settings().SetTheme(theme.LightTheme())
	makeTray(a)
	w := a.NewWindow("Fyne Demo")
	topWindow = w

	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()

	content := container.NewMax()
	title := widget.NewLabel("Component name blah")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	grid := widget.NewTextGridFromString("potato poteason")

	intro.Wrapping = fyne.TextWrapWord
	setTutorial := func(t tutorials.Tutorial) {
		title.SetText(t.Title)
		intro.SetText(t.Intro)

		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro, grid), nil, nil, nil, content)

	// file tab
	split := container.NewHSplit(makeNav(setTutorial, true), tutorial)
	split.Offset = 0.2
	w.SetContent(split)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	checkedItem := fyne.NewMenuItem("Checked", nil)
	checkedItem.Checked = true
	disabledItem := fyne.NewMenuItem("Disabled", nil)
	disabledItem.Disabled = true
	otherItem := fyne.NewMenuItem("Other", nil)
	mailItem := fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") })
	mailItem.Icon = theme.MailComposeIcon()
	otherItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
		mailItem,
	)
	fileItem := fyne.NewMenuItem("File", func() { fmt.Println("Menu New->File") })
	fileItem.Icon = theme.FileIcon()
	dirItem := fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") })
	dirItem.Icon = theme.FolderIcon()
	newItem.ChildMenu = fyne.NewMenu("",
		fileItem,
		dirItem,
		otherItem,
	)

	openSettings := func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Settings", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})

	cutShortcut := &fyne.ShortcutCut{Clipboard: w.Clipboard()}
	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(cutShortcut, w)
	})
	cutItem.Shortcut = cutShortcut
	copyShortcut := &fyne.ShortcutCopy{Clipboard: w.Clipboard()}
	copyItem := fyne.NewMenuItem("Copy", func() {
		shortcutFocused(copyShortcut, w)
	})
	copyItem.Shortcut = copyShortcut
	pasteShortcut := &fyne.ShortcutPaste{Clipboard: w.Clipboard()}
	pasteItem := fyne.NewMenuItem("Paste", func() {
		shortcutFocused(pasteShortcut, w)
	})
	pasteItem.Shortcut = pasteShortcut
	performFind := func() { fmt.Println("Menu Find") }
	findItem := fyne.NewMenuItem("Find", performFind)
	findItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyF, Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierAlt | fyne.KeyModifierShift | fyne.KeyModifierControl | fyne.KeyModifierSuper}
	w.Canvas().AddShortcut(findItem.Shortcut, func(shortcut fyne.Shortcut) {
		performFind()
	})

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://developer.fyne.io")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://fyne.io/support/")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Sponsor", func() {
			u, _ := url.Parse("https://fyne.io/sponsor/")
			_ = a.OpenURL(u)
		}))

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", newItem, checkedItem, disabledItem)
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}
	main := fyne.NewMainMenu(
		file,
		fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem),
		helpMenu,
	)
	checkedItem.Action = func() {
		checkedItem.Checked = !checkedItem.Checked
		main.Refresh()
	}
	return main
}

func makeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		h.Icon = theme.HomeIcon()
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			log.Println("System tray menu tapped")
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}

func unsupportedTutorial(t tutorials.Tutorial) bool {
	return !t.SupportWeb && fyne.CurrentDevice().IsBrowser()
}

func makeNav(setTutorial func(tutorial tutorials.Tutorial), loadPrevious bool) fyne.CanvasObject {
	var data = []string{"a", "string", "list"}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewPadded(
				widget.NewButton("", nil),
			)
			// return widget.NewLabel("template")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Button).SetText(data[id])
			item.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = func() {
				fmt.Println("I am button " + data[id])
			}
		})
	return container.NewBorder(nil, nil, nil, nil, list)

	// a := fyne.CurrentApp()

	// tree := &widget.Tree{
	// 		ChildUIDs: func(uid string) []string {
	// 			return tutorials.TutorialIndex[uid]
	// 		},
	// 	IsBranch: func(uid string) bool {
	// 		children, ok := tutorials.TutorialIndex[uid]

	// 		return ok && len(children) > 0
	// 	},
	// 	CreateNode: func(branch bool) fyne.CanvasObject {
	// 		return widget.NewLabel("Collection Widgets")
	// 	},
	// 	UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
	// 		t, ok := tutorials.Tutorials[uid]
	// 		if !ok {
	// 			fyne.LogError("Missing tutorial panel: "+uid, nil)
	// 			return
	// 		}
	// 		obj.(*widget.Label).SetText(t.Title)
	// 		if unsupportedTutorial(t) {
	// 			obj.(*widget.Label).TextStyle = fyne.TextStyle{Italic: true}
	// 		} else {
	// 			obj.(*widget.Label).TextStyle = fyne.TextStyle{}
	// 		}
	// 	},
	// 	OnSelected: func(uid string) {
	// 		if t, ok := tutorials.Tutorials[uid]; ok {
	// 			if unsupportedTutorial(t) {
	// 				return
	// 			}
	// 			a.Preferences().SetString(preferenceCurrentTutorial, uid)
	// 			setTutorial(t)
	// 		}
	// 	},

	// if loadPrevious {
	// 	currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
	// 	tree.Select(currentPref)
	// }

	// return container.NewBorder(nil, nil, nil, nil, tree)

}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	switch sh := s.(type) {
	case *fyne.ShortcutCopy:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutCut:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutPaste:
		sh.Clipboard = w.Clipboard()
	}
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}
