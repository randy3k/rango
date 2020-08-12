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

func (content *Content) GetHeightForLine(lineno, width int, prefixWidth int) int {
	h := 1;
	w := prefixWidth
	for _, c := range content.Lines[lineno] {
		if w >= width {
			h += 1
			w = 0
		}
		w += c.Width
	}
	return h
}
