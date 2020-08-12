package layout

type WindowRenderInfo struct {
	Width        int
	Height       int
	ScrollOffset int
}

type Window struct {
	Buffer *Buffer
	Info   *WindowRenderInfo
	prefix *Prefix
}

func NewWindow(buffer *Buffer, prefix *Prefix) *Window {
	return &Window{
		Buffer: buffer,
		Info: &WindowRenderInfo{
			ScrollOffset: 1e8,
		},
		prefix: prefix,
	}
}

func (win *Window) WriteToScreen(screen *Screen) {
	content := win.Buffer.CreateContent(screen.Columns, screen.Lines)
	prefixWidth := win.prefix.Width()

	nlines := len(content.Lines)
	totalHeight := 0
	lineHeights := make([]int, nlines)
	for i := 0; i < len(content.Lines); i++ {
		lineHeights[i] = content.GetHeightForLine(i, screen.Columns, prefixWidth)
		totalHeight += lineHeights[i]
	}

	offset := win.Info.ScrollOffset
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
		l := append(win.prefix.GetPrefix(i), content.Lines[i]...)

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
			if content.Cursor.Line == i && content.Cursor.Character + prefixWidth == j {
				bufferCursor.Line, bufferCursor.Column = screen.Cursor.Line, screen.Cursor.Column
			}
			screen.Feed(c)
		}
		if content.Cursor.Line == i && content.Cursor.Character + prefixWidth >= jEnd {
			bufferCursor.Line, bufferCursor.Column = screen.Cursor.Line, screen.Cursor.Column
		}
		if i < iEnd {
			screen.LineFeed()
		}
	}

	// TODO: set cursor to the focused component
	screen.Cursor = bufferCursor

	if screen.Cursor.Column == screen.Columns {
		if screen.Cursor.Line == screen.Lines - 1 {
			// render the cursor in next line
			offset++
		}
		screen.LineFeed()
	}

	win.Info = &WindowRenderInfo{
		Width:        screen.Columns,
		Height:       0,
		ScrollOffset: offset,
	}
}


// func (content *Content) Format(width, maxHeight, offset int) ([]Chars, []bool, int, ScreenCursor) {
// 	lineFragments := make([]Chars, 0)
// 	eol := make([]bool, 0)

// 	cursorLine := 0
// 	cursorColumn := 0

// 	for i, l := range content.Lines {
// 		wl := l.SplitAt(width)
// 		// get screen cursor position from document cursor
// 		if content.Cursor.Line == i {
// 			k := 0
// 			for j, lf := range wl {
// 				w := 0
// 				for _, c := range lf {
// 					if content.Cursor.Character <= k {
// 						cursorLine = len(lineFragments) + j
// 						cursorColumn = w
// 						goto found_cursor
// 					}
// 					k++
// 					w += c.Width
// 				}
// 				if content.Cursor.Character <= k {
// 					cursorLine = len(lineFragments) + j
// 					cursorColumn = w
// 					goto found_cursor
// 				}
// 			}
// 		found_cursor:
// 			if cursorColumn >= width {
// 				cursorLine++
// 				cursorColumn = 0
// 			}
// 		}
// 		for j, lf := range wl {
// 			lineFragments = append(lineFragments, lf)
// 			eol = append(eol, j+1 == len(wl))
// 		}
// 	}

// 	totalHeight := len(lineFragments)

// 	height := maxHeight

// 	// if the window is short
// 	if totalHeight < height {
// 		offset = 0
// 		height = totalHeight
// 	}

// 	// if scroll passes the last line
// 	if offset+height > totalHeight {
// 		offset = totalHeight - height
// 	}

// 	// if cursor is outside of the viewport
// 	if cursorLine < offset {
// 		offset = cursorLine
// 	} else if height == maxHeight && cursorLine + 1 > offset + height {
// 		offset = cursorLine - height + 1
// 	}

// 	return lineFragments[offset:(offset + height)],
// 		eol[offset:(offset + height)],
// 		offset,
// 		ScreenCursor{Line: cursorLine - offset, Column: cursorColumn}
// }
