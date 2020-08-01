package prompt

const (
	AttrOff = "\x1b[0m"
)

type Renderer struct {
	terminal *Terminal
	cursor ScreenCursor
	maxLine int
}

func NewRenderer(terminal *Terminal) *Renderer {
	return &Renderer{
		terminal: terminal,
	}
}

func (r *Renderer) Render(scr *Screen) {
	// TODO: diff previous screen

	t := r.terminal

	t.HideCursor()
	t.MoveCursorUp(r.cursor.Line)
	t.WriteString("\r")

	for i := 0; i <= r.maxLine; i++ {
		t.WriteString("\x1b[2K") // EL2
		if i < r.maxLine {
			t.MoveCursorDown(1)
		}
	}
	t.MoveCursorUp(r.maxLine)

	t.WriteString(AttrOff)
	cursorAttr := DefaultAttributes

	// find the last non-empty line
	var lastLine int
	for lastLine = scr.Lines - 1; lastLine >= 0; lastLine-- {
		if !scr.IsLineEmpty(lastLine) {
			break
		}
	}
	lastLine = max(lastLine, scr.Line)

	// for clearing screen in next rendering
	r.maxLine = lastLine
	r.cursor = scr.ScreenCursor

	for i := 0; i <= lastLine; i++ {
		for j := 0; j < scr.Columns; j++ {
			pos := i * scr.Columns + j
			c := scr.chars[pos]
			if c.Attributes != cursorAttr {
				cursorAttr = c.Attributes
				t.WriteString(AttrOff)
				t.WriteString(t.ColorSequence(c))
			}
			t.WriteString(string(c.Value))
		}
		if i + 1 <= lastLine && scr.eol[i] {
			t.WriteString("\r\n")
		}
	}
	t.WriteString(AttrOff)
	t.MoveCursorUp(lastLine)
	t.WriteString("\r")
	t.MoveCursorDown(r.cursor.Line)
	t.MoveCursorLeft(r.cursor.Column)
	t.ShowCursor()
}
