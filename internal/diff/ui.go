package diff

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/dorin-suletea/diferentiador~/ui"
)

func contentAsLabels(content string) []fyne.CanvasObject {

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

func (dw *DiffWidget) SetContent(diff string) {
	dw.setContent(contentAsLabels(diff))
}

func (dw *DiffWidget) setContent(lineLabels []fyne.CanvasObject) {
	dw.labelBox = container.NewVBox(lineLabels...)
	dw.scroll.Content = dw.labelBox
	dw.scroll.Refresh()
}

func (dw *DiffWidget) OnArrowDown() {
	dw.scroll.Offset.Y = dw.scroll.Offset.Y + ui.FontHeight
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
