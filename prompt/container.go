package prompt

type Container struct {
	buffer       *Buffer
	scrollOffset int
}

func NewContainer(buffer *Buffer) *Container {
	return &Container{
		buffer:       buffer,
		scrollOffset: -1,
	}
}

func perpareContent(content *Content, width, maxheight, offset int) ([]Chars, []bool, ScreenCursor) {
	lineFragments := make([]Chars, 0)
	eol := make([]bool, 0)

	var screenCursor ScreenCursor

	for i, l := range content.Lines {
		wl := l.SplitAt(width)
		// get screen cursor position from document cursor
		if content.Cursor.Line == i {
			k := 0
			for j, lf := range wl {
				w := 0
				for _, c := range lf {
					if content.Cursor.Character <= k {
						screenCursor.Line = len(lineFragments) + j
						screenCursor.Column = w
						goto found_cursor
					}
					k++
					w += c.Width
				}
				if content.Cursor.Character <= k {
					screenCursor.Line = len(lineFragments) + j
					screenCursor.Column = w
					goto found_cursor
				}
			}
		found_cursor:
			if screenCursor.Column >= width {
				screenCursor.Line++
				screenCursor.Column = 0
			}
		}
		for j, lf := range wl {
			lineFragments = append(lineFragments, lf)
			eol = append(eol, j+1 == len(wl))
		}
	}

	totalHeight := len(lineFragments)

	if totalHeight < maxheight {
		offset = 0
		maxheight = totalHeight
	}

	if offset < 0 || offset+maxheight > totalHeight {
		offset = totalHeight - maxheight
	}

	if offset > 0 {
		screenCursor.Line -= offset
	}

	return lineFragments[offset:(offset + maxheight)], eol[offset:(offset + maxheight)], screenCursor
}

func (c *Container) WriteToScreen(scr *Screen) {
	content := c.buffer.CreateContent(scr.Columns, scr.Lines)
	lines, eol, cursor := perpareContent(content, scr.Columns, scr.Lines, c.scrollOffset)

	scr.Reset()
	for i, l := range lines {
		scr.SetCharsAt(i, 0, l, eol[i])
	}
	scr.Cursor = cursor
}
