package prompt

const (
	AttrOff = "\x1b[0m"
)

type Renderer struct {
	terminal *Terminal
	previousScreen *Screen
	cursor   ScreenCursor
	maxLine  int
}

func NewRenderer(terminal *Terminal) *Renderer {
	return &Renderer{
		terminal: terminal,
	}
}

func (r *Renderer) Render(screen *Screen) {
	// TODO: diff previous screen

	diff, _ := screen.Diff(r.previousScreen)

	t := r.terminal

	t.HideCursor()
	t.MoveCursorUp(r.cursor.Line)
	t.WriteString("\r")

	for i := 0; i <= r.maxLine; i++ {
		if diff[i] {
			t.WriteString("\x1b[0K")
		}
		if i < r.maxLine {
			t.MoveCursorDown(1)
		}
	}
	t.MoveCursorUp(r.maxLine)

	t.WriteString(AttrOff)
	cursorAttr := DefaultAttributes

	// find the last non-empty line
	var lastLine int
	for lastLine = screen.Lines - 1; lastLine >= 0; lastLine-- {
		if !screen.IsLineEmpty(lastLine) {
			break
		}
	}

	lastLine = max(lastLine, screen.Cursor.Line)
	// for clearing screen in next rendering
	r.maxLine = lastLine
	r.cursor = screen.Cursor

	for i := 0; i <= lastLine; i++ {
		if diff[i] {
			for j := 0; j < screen.Columns; j++ {
				pos := i*screen.Columns + j
				c := screen.chars[pos]
				if c.Attributes != cursorAttr {
					cursorAttr = c.Attributes
					t.WriteString(AttrOff)
					t.WriteString(t.ColorSequence(c))
				}
				t.WriteString(string(c.Value))
			}
			if i+1 <= lastLine && screen.IsLineEOL(i) {
				t.WriteString("\r\n")
			}
		} else if i+1 <= lastLine {
			t.WriteString("\r\n")
		}
	}
	t.WriteString(AttrOff)
	t.MoveCursorUp(lastLine)
	t.WriteString("\r")
	t.MoveCursorDown(r.cursor.Line)
	t.MoveCursorLeft(r.cursor.Column)
	t.ShowCursor()
	t.Flush()

	r.previousScreen = screen
}
