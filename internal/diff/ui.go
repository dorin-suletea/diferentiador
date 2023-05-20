package diff

import (
	"bytes"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/dorin-suletea/diferentiador~/ui"
)

func NewDiffWidget(content string) *fyne.Container {

	uiLines := []*canvas.Text{}
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {

		if len(line) == 0 {
			tw := canvas.NewText(line, ui.Gray)
			uiLines = append(uiLines, tw)
			continue
		}

		switch line[0] {
		case '-':
			tw := canvas.NewText(line, ui.Red)
			tw.TextStyle = ui.FontMono
			uiLines = append(uiLines, tw)
		case '+':
			tw := canvas.NewText(line, ui.Green)
			tw.TextStyle = ui.FontMono
			uiLines = append(uiLines, tw)
		default:
			tw := canvas.NewText(line, ui.Gray)
			tw.TextStyle = ui.FontMono
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
