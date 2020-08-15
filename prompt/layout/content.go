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
	Width int
	heightCache map[heightCacheKey]int
}

func NewContent(lines []Chars, cursor DocumentCursor, width int) *Content {
	return &Content{
		Lines:  lines,
		Cursor: cursor,
		Width: width,
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
			w = 0
		}
	}
	content.heightCache[heightCacheKey{lineno, width}] = h
	return h
}
