package layout

type WindowRenderInfo struct {
	Width        int
	Height       int
	ScrollOffset int
}

type Window struct {
	Buffer *Buffer
	Margin *Margin
	Info   *WindowRenderInfo
}

func NewWindow(buffer *Buffer, margin *Margin) *Window {
	return &Window{
		Buffer: buffer,
		Margin: margin,
		Info: &WindowRenderInfo{
			ScrollOffset: 1e8,
		},
	}
}

func (win *Window) WriteToScreen(screen *Screen) {
	content := win.Buffer.CreateContent(screen.Columns, screen.Lines)
	previousOffset := win.Info.ScrollOffset
	lines, eol, offset, cursor := content.Format(screen.Columns, screen.Lines, previousOffset)

	for i, l := range lines {
		screen.SetCharsAt(i, 0, l)
		if eol[i] {
			screen.Chars[screen.Columns*(i+1)-1].EOL = true
		}
	}
	screen.Cursor = cursor

	win.Info = &WindowRenderInfo{
		Width:        screen.Columns,
		Height:       len(lines),
		ScrollOffset: offset,
	}
}
