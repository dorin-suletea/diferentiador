package diff

import (
	"bytes"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/dorin-suletea/diferentiador~/ui"
)

// *fyne.Container
// box := container.NewVBox(asCanvas...)
func ContentAsLabels(content string) []fyne.CanvasObject {

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
	return asCanvas
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

// ----------------------
// DiffWidget
// ----------------------
type DiffWidget struct {
	*ui.FocusComponent
	scroll   *container.Scroll
	labelBox *fyne.Container
}

func NewDiffWidget(lineLabels []fyne.CanvasObject) *DiffWidget {
	labelBox := container.NewVBox(lineLabels...)
	scroll := container.NewScroll(labelBox)
	return &DiffWidget{ui.NewFocusComponent(scroll), scroll, labelBox} //scroll
}

func (dw *DiffWidget) SetContent(lineLabels []fyne.CanvasObject) {
	dw.labelBox = container.NewVBox(lineLabels...)
	dw.scroll.Content = dw.labelBox
	dw.scroll.Refresh()
}

func (dw *DiffWidget) OnArrowDown() {
	dw.scroll.Offset.Y = dw.scroll.Offset.Y + ui.FontHeight
	println(dw.scroll.Offset.Y)
	dw.scroll.Refresh()
}

func (dw *DiffWidget) OnArrowUp() {
	dw.scroll.Offset.Y = dw.scroll.Offset.Y - ui.FontHeight
	dw.scroll.Refresh()
}

func (dw *DiffWidget) OnArrowRight() {
	dw.scroll.Offset.X = dw.scroll.Offset.X + ui.FontWidth
	dw.scroll.Refresh()
}

func (dw *DiffWidget) OnArrowLeft() {
	dw.scroll.Offset.X = dw.scroll.Offset.X - ui.FontWidth
	dw.scroll.Refresh()
}
