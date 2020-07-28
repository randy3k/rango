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
	content := c.buffer.CreateContent(scr.w, scr.h)

	lines, eol := content.GetLines(scr.w, scr.h, c.scrollOffset)
	for i, l := range lines {
		scr.SetCharsAt(i, 0, l, !eol[i])
	}
}
