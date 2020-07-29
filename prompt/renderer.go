package prompt

type Renderer struct {
	terminal *Terminal
	posRelOrgin ScreenCursor
}

func NewRenderer(terminal *Terminal) *Renderer {
	return &Renderer{
		terminal: terminal,
	}
}

func (r *Renderer) Render(scr *Screen) {
	t := r.terminal
	t.WriteString(t.ti.AttrOff)
	cursorAttr := DefaultAttributes
	for i := 0; i < scr.h; i++ {
		for j := 0; j < scr.w; j++ {
			pos := i * scr.w + j
			c := scr.chars[pos]
			if c.Attributes != cursorAttr {
				cursorAttr = c.Attributes
				t.WriteString(t.ti.AttrOff)
				t.WriteString(t.ti.TParm(t.ti.TColor(
					c.Foreground.Code8Bits(), c.Background.Code8Bits())))
			}
			t.WriteString(string(c.Value))
		}
		t.WriteString("\r\n")
	}
	t.WriteString(t.ti.AttrOff)
	// myprint(scr.chars[0].Foreground.Code4Bits())
	// myprint(scr.chars[1].Attributes == scr.chars[2].Attributes)

	// fmt.Printf("%x\n", scr.chars[148].Attributes.Foreground - 1)
}
