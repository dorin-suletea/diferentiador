package app

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/dorin-suletea/diferentiador~/internal"
)

func contentAsLabels(content string) []fyne.CanvasObject {

	uiLines := []*canvas.Text{}
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {

		if len(line) == 0 {
			tw := canvas.NewText(line, internal.Gray)
			uiLines = append(uiLines, tw)
			continue
		}

		switch line[0] {
		case '-':
			tw := canvas.NewText(line, internal.Red)
			tw.TextStyle = internal.FontMono
			uiLines = append(uiLines, tw)
		case '+':
			tw := canvas.NewText(line, internal.Green)
			tw.TextStyle = internal.FontMono
			uiLines = append(uiLines, tw)
		default:
			tw := canvas.NewText(line, internal.Gray)
			tw.TextStyle = internal.FontMono
			uiLines = append(uiLines, tw)
		}
	}

	asCanvas := make([]fyne.CanvasObject, len(uiLines))
	for i, val := range uiLines {
		asCanvas[i] = val
	}
	return asCanvas
}

var _ FileWidgetListener = (*DiffWidget)(nil)

// ----------------------
// DiffWidget
// ----------------------
type DiffWidget struct {
	*internal.FocusComponent
	scroll   *container.Scroll
	labelBox *fyne.Container
	data     *FileDifCache
}

func NewDiffWidget(data *FileDifCache) *DiffWidget {
	labelBox := container.NewVBox()
	scroll := container.NewScroll(labelBox)
	return &DiffWidget{internal.NewFocusComponent(scroll), scroll, labelBox, data}
}

func (dw *DiffWidget) OnFileSelected(selection FileStatus) {
	lines := dw.data.GetContent(selection)
	linesAsLabels := contentAsLabels(lines)
	dw.labelBox = container.NewVBox(linesAsLabels...)
	dw.scroll.Content = dw.labelBox
	dw.scroll.Refresh()
}

func (dw *DiffWidget) OnArrowDown() {
	dw.scroll.Offset.Y = dw.scroll.Offset.Y + internal.FontHeight
	dw.scroll.Refresh()
}

func (dw *DiffWidget) OnArrowUp() {
	dw.scroll.Offset.Y = dw.scroll.Offset.Y - internal.FontHeight
	dw.scroll.Refresh()
}

func (dw *DiffWidget) OnArrowRight() {
	dw.scroll.Offset.X = dw.scroll.Offset.X + internal.FontWidth
	dw.scroll.Refresh()
}

func (dw *DiffWidget) OnArrowLeft() {
	dw.scroll.Offset.X = dw.scroll.Offset.X - internal.FontWidth
	dw.scroll.Refresh()
}
