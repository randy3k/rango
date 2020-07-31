package prompt

const (
	AttrOff = "\x1b[0m"
)

type Renderer struct {
	terminal *Terminal
	cursor ScreenCursor
	maxRow int
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
	t.MoveCursorUp(r.cursor.row)
	t.WriteString("\r")
	t.WriteString("\x1b7")  // DECSC

	for i := 0; i <= r.maxRow; i++ {
		t.WriteString("\x1b[2K") // EL2
		if i < r.maxRow {
			t.MoveCursorDown(1)
		}
	}
	t.WriteString("\x1b8")

	t.WriteString(AttrOff)
	cursorAttr := DefaultAttributes

	// find the last non-empty row
	var lastRow int
	for lastRow = scr.lines - 1; lastRow >= 0; lastRow-- {
		if !scr.IsRowEmpty(lastRow) {
			break
		}
	}
	lastRow = max(lastRow, scr.row)

	// for clearing screen in next rendering
	r.maxRow = lastRow
	r.cursor = scr.ScreenCursor

	for i := 0; i <= lastRow; i++ {
		for j := 0; j < scr.columns; j++ {
			pos := i * scr.columns + j
			c := scr.chars[pos]
			if c.Attributes != cursorAttr {
				cursorAttr = c.Attributes
				t.WriteString(AttrOff)
				t.WriteString(t.ColorSequence(c))
			}
			t.WriteString(string(c.Value))
		}
		if i + 1 <= lastRow && scr.eol[i] {
			t.WriteString("\r\n")
		}
	}
	t.WriteString(AttrOff)
	t.WriteString("\x1b8")  // DECRC
	t.MoveCursorDown(r.cursor.row)
	t.MoveCursorLeft(r.cursor.col)
	t.ShowCursor()
}
