package layout

import (
	. "github.com/randy3k/rango/prompt/char"
)


type Content struct {
	Lines  []Chars
	Cursor DocumentCursor
}

func NewContent(lines []Chars, cursor DocumentCursor) *Content {
	return &Content{
		Lines:  lines,
		Cursor: cursor,
	}
}

func (content *Content) Format(width, maxHeight, offset int) ([]Chars, []bool, int, ScreenCursor) {
	lineFragments := make([]Chars, 0)
	eol := make([]bool, 0)

	cursorLine := 0
	cursorColumn := 0

	for i, l := range content.Lines {
		wl := l.SplitAt(width)
		// get screen cursor position from document cursor
		if content.Cursor.Line == i {
			k := 0
			for j, lf := range wl {
				w := 0
				for _, c := range lf {
					if content.Cursor.Character <= k {
						cursorLine = len(lineFragments) + j
						cursorColumn = w
						goto found_cursor
					}
					k++
					w += c.Width
				}
				if content.Cursor.Character <= k {
					cursorLine = len(lineFragments) + j
					cursorColumn = w
					goto found_cursor
				}
			}
		found_cursor:
			if cursorColumn >= width {
				cursorLine++
				cursorColumn = 0
			}
		}
		for j, lf := range wl {
			lineFragments = append(lineFragments, lf)
			eol = append(eol, j+1 == len(wl))
		}
	}

	totalHeight := len(lineFragments)

	height := maxHeight
	if totalHeight < height {
		offset = 0
		height = totalHeight
	}

	if offset+height > totalHeight {
		offset = totalHeight - height
	}

	// if cursor is outside of the viewport
	if cursorLine < offset {
		offset = cursorLine
	} else if height == maxHeight && cursorLine + 1 > offset + height {
		offset = cursorLine - height + 1
	}

	return lineFragments[offset:(offset + height)],
		eol[offset:(offset + height)],
		offset,
		ScreenCursor{Line: cursorLine - offset, Column: cursorColumn}
}
