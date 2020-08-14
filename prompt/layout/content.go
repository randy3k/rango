package layout

import (
	. "github.com/randy3k/rango/prompt/char"
)


type Content struct {
	Lines  []Chars
	Cursor DocumentCursor
	PrefixWidth int
}

func NewContent(lines []Chars, cursor DocumentCursor, prefixWidth int) *Content {
	return &Content{
		Lines:  lines,
		Cursor: cursor,
		PrefixWidth: prefixWidth,
	}
}


func (content *Content) GetHeightForLine(lineno, width int) int {
	h := 1;
	w := 0
	for _, c := range content.Lines[lineno] {
		w += c.Width
		if w > width {
			h += 1
			w = 0
		}
	}
	return h
}
