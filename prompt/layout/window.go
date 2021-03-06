package layout

import (
	. "github.com/randy3k/rango/prompt/char"
)

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

	// move cursor if it is out of the viewport
	absCursorY, absCursorX := content.GetAbsoluteCursorPosition(screen.Columns)
	if absCursorY < offset {
		offset = absCursorY
	} else if height == screen.Lines && absCursorY + 1 > offset + height {
		offset = absCursorY - height + 1
	}
	bufferCursor := ScreenCursor{Line: absCursorY - offset, Column: absCursorX}

	win.writeBufferToScreen(screen, offset, height, lineHeights, content.Lines)

	screen.Cursor = bufferCursor

	if screen.Cursor.Column >= screen.Columns {
		// render the cursor in next line
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


func (win *Window) writeBufferToScreen(
			screen *Screen, offset, height int, lineHeights []int, contentLines []Chars) {
	// go home
	screen.Cursor.Line = 0
	screen.Cursor.Column = 0

	nlines := len(lineHeights)

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
		l := contentLines[i]

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
			screen.Feed(c)
		}
		if i < iEnd {
			screen.LineFeed()
		}
	}
}
