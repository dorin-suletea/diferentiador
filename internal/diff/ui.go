package diff

import (
	"bytes"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewDiffWidget(content string) *fyne.Container {
	green := &color.NRGBA{R: 64, G: 192, B: 64, A: 128}
	red := &color.NRGBA{R: 192, G: 64, B: 64, A: 128}
	gray := &color.NRGBA{R: 96, G: 96, B: 96, A: 255}
	font := fyne.TextStyle{Italic: true, Monospace: true}

	uiLines := []*canvas.Text{}
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {

		if len(line) == 0 {
			tw := canvas.NewText(line, gray)
			uiLines = append(uiLines, tw)
			continue
		}

		switch line[0] {
		case '-':
			tw := canvas.NewText(line, red)
			tw.TextStyle = font
			uiLines = append(uiLines, tw)
		case '+':
			tw := canvas.NewText(line, green)
			tw.TextStyle = font
			uiLines = append(uiLines, tw)
		default:
			tw := canvas.NewText(line, gray)
			tw.TextStyle = font
			uiLines = append(uiLines, tw)
		}
	}

	asCanvas := make([]fyne.CanvasObject, len(uiLines))
	for i, val := range uiLines {
		asCanvas[i] = val
	}
	box := container.NewVBox(asCanvas...)
	return box
}

/*
	Prefixed all lines of a given \n separated string.
	This is useful to mark all lines of a deleted file with '-' and with '+' for an unstaged file.
	While these are not technically line changes from a GIT perspective this provides useful feedback to the user.
*/
func MarkLines(content string, prefix byte) string {
	prefixed := bytes.Buffer{}
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {
		prefixed.WriteByte(prefix)
		prefixed.WriteString(line)
		prefixed.WriteString("\n")
	}
	return prefixed.String()
}
