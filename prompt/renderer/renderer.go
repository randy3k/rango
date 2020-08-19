package renderer

import (
	"github.com/randy3k/rango/prompt/char"
	. "github.com/randy3k/rango/prompt/layout"
	. "github.com/randy3k/rango/prompt/terminal"
)

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
	diff, pos := screen.Diff(r.previousScreen)

	t := r.terminal

	t.HideCursor()
	t.MoveCursorUp(r.cursor.Line)
	t.WriteString("\r")

	for i := 0; i <= r.maxLine; i++ {
		if diff[i] {
			t.MoveCursorRight(pos[i])
			t.WriteString("\x1b[0K")
			t.WriteString("\r")
		}
		if i < r.maxLine {
			t.MoveCursorDown(1)
		}
	}
	t.MoveCursorUp(r.maxLine)

	t.WriteString(AttrOff)
	cursorAttr := char.DefaultAttributes

	// find the last non-empty line
	lastLine := screen.Lines - 1
	for lastLine >= 0 {
		if !screen.IsLineEmpty(lastLine) {
			break
		}
		lastLine--
	}

	lastLine = max(lastLine, screen.Cursor.Line)
	// for clearing screen in next rendering
	r.maxLine = lastLine
	r.cursor = screen.Cursor

	for i := 0; i <= lastLine; i++ {
		if diff[i] {
			t.MoveCursorRight(pos[i])
			for j := pos[i]; j < screen.Columns; j++ {
				pos := i*screen.Columns + j
				c := screen.Cells[pos].Char
				if c.Attributes != cursorAttr {
					cursorAttr = c.Attributes
					t.WriteString(AttrOff)
					t.WriteString(t.ColorSequence(c))
				}
				t.WriteString(string(c.Value))
			}
			if i+1 <= lastLine && !screen.IsLineContinuation(i) {
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
	t.MoveCursorRight(r.cursor.Column)
	t.ShowCursor()
	t.Flush()

	r.previousScreen = screen
}
