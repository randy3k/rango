package prompt


type WindowRenderInfo struct {
	Width int
	Height int
	ScrollOffset int
}

type Window struct {
	Buffer       *Buffer
	Info *WindowRenderInfo
}

func NewWindow(buffer *Buffer) *Window {
	return &Window{
		Buffer: buffer,
		Info: &WindowRenderInfo{
			ScrollOffset: 1e8,
		},
	}
}

func (win *Window) WriteToScreen(scr *Screen) {
	content := win.Buffer.CreateContent(scr.Columns, scr.Lines)
	previousOffset := win.Info.ScrollOffset
	lines, eol, offset, cursor := content.GetLines(scr.Columns, scr.Lines, previousOffset)

	for i, l := range lines {
		scr.SetCharsAt(i, 0, l, eol[i])
	}
	scr.Cursor = cursor

	win.Info = &WindowRenderInfo{
		Width: scr.Columns,
		Height: len(lines),
		ScrollOffset: offset,
	}
}
