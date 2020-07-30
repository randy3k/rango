package prompt


type Container struct {
	buffer *Buffer
	scrollOffset int
}

func NewContainer(buffer *Buffer) *Container {
	return &Container{
		buffer: buffer,
		scrollOffset: -1,
	}
}


func perpareContent(content *Content, width, maxheight, offset int) ([]Chars, []bool, ScreenCursor) {
	lineFragments := make([]Chars, 0)
	eol := make([]bool, 0)

	var scursor ScreenCursor

	for i, l := range content.Lines {
		wl := l.SplitAt(width)
		// get screen cursor position from document cursor
		if content.Cursor.Line == i {
			k := 0
			for j, lf := range wl {
				w := 0
				for _, c := range lf {
					if content.Cursor.Character <= k {
						scursor.row = len(lineFragments) + j
						scursor.col = w
						goto found_cursor
					}
					k++
					w += c.Width
				}
				if content.Cursor.Character <= k {
					scursor.row = len(lineFragments) + j
					scursor.col = w
					goto found_cursor
				}
			}
		found_cursor:
			if scursor.col >= width {
				scursor.row++
				scursor.col = 0
			}
		}
		for j, lf := range wl {
			lineFragments = append(lineFragments, lf)
			eol = append(eol, j + 1 == len(wl))
		}
	}

	totalHeight := len(lineFragments)

	if (totalHeight < maxheight) {
		offset = 0
		maxheight = totalHeight
	}

	if (offset < 0 || offset + maxheight > totalHeight) {
		offset = totalHeight - maxheight
	}

	if offset > 0 {
		scursor.row -= offset
	}

	return lineFragments[offset:(offset+maxheight)], eol[offset:(offset+maxheight)], scursor
}


func (c *Container) WriteToScreen(scr *Screen) {
	content := c.buffer.CreateContent(scr.columns, scr.lines)
	lines, eol, cursor := perpareContent(content, scr.columns, scr.lines, c.scrollOffset)

	scr.Reset()
	for i, l := range lines {
		scr.SetCharsAt(i, 0, l, eol[i])
	}
	scr.ScreenCursor = cursor
}
