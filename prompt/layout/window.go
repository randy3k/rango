package layout

type WindowRenderInfo struct {
	Width        int
	Height       int
	ScreenWidth int
	ScreenHeight int
	ScrollOffset int
	ScreenCursor ScreenCursor
}

type Window struct {
	Buffer *Buffer
	RenderInfo   *WindowRenderInfo
}

func NewWindow(buffer *Buffer) *Window {
	return &Window{
		Buffer: buffer,
		RenderInfo: &WindowRenderInfo{
			ScrollOffset: 1e8,
		},
	}
}


func (win *Window) WriteToScreen(screen *Screen) {
	content := win.Buffer.CreateContent(screen.Columns, screen.Lines)

	nlines := len(content.Lines)
	totalHeight := 0
	lineHeights := make([]int, nlines)
	for i := 0; i < len(content.Lines); i++ {
		lineHeights[i] = content.GetHeightForLine(i, screen.Columns)
		totalHeight += lineHeights[i]
	}

	offset := win.RenderInfo.ScrollOffset
	height := screen.Lines

	// if the window is short
	if totalHeight < height {
		offset = 0
		height = totalHeight
	}

	// if scroll passes the last line
	if offset + height > totalHeight {
		offset = totalHeight - height
	}

	// go home
	screen.Cursor.Line = 0
	screen.Cursor.Column = 0

	bufferCursor := ScreenCursor{}

	// find display region
	iBegin := 0
	deltaBegin := 0
	for i, h := 0, lineHeights[0]; i < nlines; i++ {
		if h > offset {
			iBegin = i
			deltaBegin = offset - (h - lineHeights[i])
			break
		}
		h = h + lineHeights[i]
	}

	iEnd := 0
	deltaEnd := 0
	for i, h := 0, 0; i < nlines; i++ {
		h = h + lineHeights[i]
		if h >= offset + height {
			iEnd = i
			deltaEnd = h - (offset + height)
			break
		}
	}

	for i := iBegin; i <= iEnd; i++ {
		l := content.Lines[i]

		jBegin := 0
		if i == iBegin {
			for deltaBegin > 0 {
				w := 0
				for k := jBegin; k < len(l); k++ {
					if w >= screen.Columns {
						jBegin = k
						break
					}
					w += l[k].Width
				}
				deltaBegin--
			}
		}

		jEnd := len(l)
		if i == iEnd {
			if deltaEnd > 0 {
				jEnd = 0
				deltaDisplay := lineHeights[i] - deltaEnd
				for deltaDisplay > 0 {
					w := 0
					for k := jEnd; k < len(l); k++ {
						if w >= screen.Columns {
							jEnd = k
							break
						}
						w += l[k].Width
					}
					deltaDisplay--
				}
			}
		}

		for j := jBegin; j < jEnd; j++ {
			c := l[j]
			if content.Cursor.Line == i && content.Cursor.Character + content.PrefixWidth == j {
				bufferCursor.Line, bufferCursor.Column = screen.Cursor.Line, screen.Cursor.Column
			}
			screen.Feed(c)
		}
		if content.Cursor.Line == i && content.Cursor.Character + content.PrefixWidth >= jEnd {
			bufferCursor.Line, bufferCursor.Column = screen.Cursor.Line, screen.Cursor.Column
		}
		if i < iEnd {
			screen.LineFeed()
		}
	}

	// TODO: set cursor to the focused component
	screen.Cursor = bufferCursor

	// render the cursor in next line
	if screen.Cursor.Column >= screen.Columns {
		if screen.Cursor.Line == screen.Lines - 1 {
			offset++
		}
		screen.LineFeed()
	}

	win.RenderInfo = &WindowRenderInfo{
		Width:        screen.Columns,
		Height:       height,
		ScreenWidth: screen.Columns,
		ScreenHeight: screen.Lines,
		ScrollOffset: offset,
		ScreenCursor: screen.Cursor,
	}
}
