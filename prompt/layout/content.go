package layout

import (
	. "github.com/randy3k/rango/prompt/char"
)

type heightCacheKey struct {
	lineno int
	width int
}


type Content struct {
	Lines  []Chars
	Cursor DocumentCursor
	PrefixLen int
	heightCache map[heightCacheKey]int
}

func NewContent(lines []Chars, cursor DocumentCursor, prefixLen int) *Content {
	return &Content{
		Lines:  lines,
		Cursor: cursor,
		PrefixLen: prefixLen,
		heightCache: map[heightCacheKey]int{},
	}
}


func (content *Content) GetHeightForLine(lineno, width int) int {
	if h, ok := content.heightCache[heightCacheKey{lineno, width}]; ok {
		return h
	}
	h := 1;
	w := 0
	for _, c := range content.Lines[lineno] {
		w += c.Width
		if w > width {
			h += 1
			w = c.Width
		}
	}
	content.heightCache[heightCacheKey{lineno, width}] = h
	return h
}



func (content *Content) GetAbsoluteCursorPosition(width int) (int, int) {
	cumHeight := 0
	for i, l := range content.Lines {
		h := 1
		w := 0
		for j, c := range l {
			w += c.Width
			if w > width {
				h += 1
				w = c.Width
			}
			if content.Cursor.Line == i && content.Cursor.Character + content.PrefixLen == j {
				return cumHeight, w - c.Width
			}
		}
		if content.Cursor.Line == i && content.Cursor.Character + content.PrefixLen >= len(l) {
			return cumHeight, len(l)
		}
		cumHeight += h
	}
	return -1, -1
}
