package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	darkGray   = &color.NRGBA{R: 96, G: 96, B: 96, A: 255}
	lightGray  = &color.NRGBA{R: 160, G: 160, B: 160, A: 255}
	darkOrange = &color.NRGBA{R: 204, G: 102, B: 0, A: 255}
	purple     = &color.NRGBA{R: 255, G: 102, B: 255, A: 255}
)

// customTheme is a simple demonstration of a bespoke theme loaded by a Fyne app.
type customTheme struct {
}

func (customTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	// case theme.ColorNameButton:
	// 	return darkGray
	// case theme.ColorNameForeground: // button text color
	// 	return darkOrange

	// case theme.ColorNameBackground:
	// 	return darkGray
	// case theme.ColorNamePrimary, theme.ColorNameHover, theme.ColorNameFocus:
	// 	return purple

	// case theme.ColorNameDisabled:
	// 	return color.Black
	// case theme.ColorNamePlaceHolder, theme.ColorNameScrollBar:
	// 	return grey

	// case theme.ColorNameShadow:
	// 	return &color.RGBA{R: 0xcc, G: 0xcc, B: 0xcc, A: 0xcc}

	case "scrollBar":
		return purple
	default:
		return theme.LightTheme().Color(c, v)
	}
}

func (customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DarkTheme().Font(style)
}

func (customTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DarkTheme().Icon(n)
}

func (customTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	// case theme.SizeNamePadding:
	// 	return 8
	// case theme.SizeNameInlineIcon:
	// 	return 20
	// case theme.SizeNameScrollBar:
	// 	return 10
	// case theme.SizeNameScrollBarSmall:
	// 	return 5
	case theme.SizeNameText:
		return 14
	// case theme.SizeNameHeadingText:
	// 	return 30
	// case theme.SizeNameSubHeadingText:
	// 	return 25
	// case theme.SizeNameCaptionText:
	// 	return 15
	// case theme.SizeNameInputBorder:
	// 	return 1
	default:
		return 0
	}
}

func NewCustomTheme() fyne.Theme {
	return &customTheme{}
}
