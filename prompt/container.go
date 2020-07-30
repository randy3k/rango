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

func (c *Container) WriteToScreen(scr *Screen) {
	content := c.buffer.CreateContent(scr.columns, scr.lines)
	lines, eol := content.GetRows(scr.columns, scr.lines, c.scrollOffset)

	scr.Reset()
	for i, l := range lines {
		scr.SetCharsAt(i, 0, l, eol[i])
	}
	// FIXME: handle char width and line wrap
	scr.row = content.Cursor.Line
	scr.col = content.Cursor.Character
}
