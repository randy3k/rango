package prompt

import (
	"fmt"
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
	t := r.terminal

	t.WriteString(t.ti.HideCursor)
	t.WriteString(fmt.Sprintf("\x1b[%dA\r", r.cursor.row))  // CUU and CR
	for i := 0; i <= r.maxRow; i++ {
		t.WriteString("\x1b[2K") // EL2
	}

	t.WriteString(t.ti.AttrOff)
	cursorAttr := DefaultAttributes

	// find the last non-empty row
	lastRow := scr.lines - 1
	for i := scr.lines - 1; i >= 0; i-- {
		if !scr.IsRowEmpty(i) {
			lastRow = i
			break
		}
	}

	// for clearing screen in next rendering
	r.maxRow = lastRow
	r.cursor = scr.ScreenCursor

	for i := 0; i <= lastRow; i++ {
		for j := 0; j < scr.columns; j++ {
			pos := i * scr.columns + j
			c := scr.chars[pos]
			if c.Attributes != cursorAttr {
				cursorAttr = c.Attributes
				t.WriteString(t.ti.AttrOff)
				t.WriteString(t.TColor(c))
			}
			t.WriteString(string(c.Value))
		}
		if i + 1 <= lastRow {
			t.WriteString("\r\n")
		}
	}
	t.WriteString(t.ti.AttrOff)
	t.WriteString(t.Goto(r.cursor.row, r.cursor.col))
	t.WriteString(t.ti.ShowCursor)
}
