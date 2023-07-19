package main

//sudo apt-get install libgl1-mesa-dev xorg-dev
// DEMO code
//https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
//go run fyne.io/fyne/v2/cmd/fyne_demo@latest

import (
	"github.com/dorin-suletea/diferentiador~/app"
	"github.com/dorin-suletea/diferentiador~/internal"
)

const refreshRateSeconds int = 1

func main() {

	application := internal.NewApp()
	fileCache := app.NewChangedFilesCache(refreshRateSeconds)
	diffCache := app.DiffCache(fileCache, refreshRateSeconds)

	diffWidget := app.NewDiffWidget(diffCache)
	fileWidget := app.NewChangedFilesWidget(fileCache, diffWidget)
	fileCache.RegisterCacheListener(fileWidget)
	diffCache.RegisterCacheListener(diffWidget)

	application.AddComponent(fileWidget)
	application.AddComponent(diffWidget)

	application.AddShortcut(internal.ShCycleFocus, func() { application.CycleFocus() })
	application.AddShortcut(internal.ShArrowDown, func() { application.OnArrowDown() })
	application.AddShortcut(internal.ShArrowUp, func() { application.OnArrowUp() })
	application.AddShortcut(internal.ShArrowRight, func() { application.OnArrowRight() })
	application.AddShortcut(internal.ShArrowLeft, func() { application.OnArrowLeft() })
	application.ShowAndRun()
}
