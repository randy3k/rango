package prompt

type Content struct {
	Lines []Chars
	Cursor Position
	cursor_map map[Position]Position
}


func NewContent(lines []Chars, cursor Position) *Content {
	return &Content{
		Lines: lines,
		Cursor: cursor,
	}
}


func (c *Content) GetLines(width, maxheight, offset int) ([]Chars, []bool) {
	lineFragments := make([]Chars, 0)
	lineno := make([]int, 0)
	eol := make([]bool, 0)


	for i := range c.Lines {
		wl := c.Lines[i].SplitAt(width)
		for j, lf := range wl {
			lineFragments = append(lineFragments, lf)
			lineno = append(lineno, i)
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

	return lineFragments[offset:(offset+maxheight)], eol[offset:(offset+maxheight)]
}
