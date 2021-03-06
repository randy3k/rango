package layout

import (
	"strings"
	. "github.com/randy3k/rango/prompt/char"
)

type ScreenCursor struct {
	Line   int
	Column int
}

type Cell struct {
	Continuation bool
	Char
}

type Screen struct {
	Lines   int
	Columns int
	Cursor  ScreenCursor
	Cells   []Cell
}

func NewScreen(h int, w int) *Screen {
	return &Screen{
		Columns: w,
		Lines:   h,
		Cells:   make([]Cell, w*h),
	}
}

func (screen *Screen) String() string {
	s := ""
	for i := 0; i < screen.Lines; i++ {
		line := make([]string, screen.Columns)
		for j := 0; j < screen.Columns; {
			pos := screen.Columns*i + j
			c := screen.Cells[pos]
			if c.Value == 0 {
				line[j] = " "
				j += 1
			} else {
				line[j] = string(c.Value)
				j = j + c.Width
			}
		}
		s += "|" + strings.Join(line, "")
		if screen.IsLineContinuation(i) {
			s += "+\n"
		} else {
			s += "|\n"
		}
	}
	return s
}

func (screen *Screen) Feed(c Char) (int, int) {
	line := screen.Cursor.Line
	col := screen.Cursor.Column
	if screen.Cursor.Column + c.Width > screen.Columns {
		screen.Cells[screen.Columns*(line+1) - 1].Continuation = true
		screen.LineFeed()
	}
	pos := screen.Columns*screen.Cursor.Line + screen.Cursor.Column
	screen.Cells[pos].Char = c
	screen.Cursor.Column += c.Width
	return line, col
}

func (screen *Screen) IsLineContinuation(line int) bool {
	return screen.Cells[screen.Columns*(line+1) - 1].Continuation
}

func (screen *Screen) LineFeed() {
	screen.Cursor.Column = 0
	screen.Cursor.Line += 1
	if screen.Cursor.Line == screen.Lines {
		screen.Cells = append(screen.Cells[screen.Columns:], make([]Cell, screen.Columns)...)
		screen.Cursor.Line -= 1
	}
}

func (screen *Screen) GoTo(line int, col int) {
	screen.Cursor.Line = max(0, min(line, screen.Lines-1))
	screen.Cursor.Column = max(0, min(col, screen.Columns-1))
}


func (screen *Screen) IsLineEmpty(line int) bool {
	for j := 0; j < screen.Columns; j++ {
		if screen.Cells[line*screen.Columns+j].Value > 0 {
			return false
		}
	}
	return true
}

func (screen *Screen) Diff(pScreen *Screen) (diff []bool, loc []int) {
	diff = make([]bool, screen.Lines)
	loc = make([]int, screen.Lines)

	if pScreen == nil || screen.Lines != pScreen.Lines || screen.Columns != pScreen.Columns {
		for i := 0; i < screen.Lines; i++ {
			diff[i] = true
			loc[i] = 0
		}
		return
	}

	for i := 0; i < screen.Lines; i++ {
		for j := 0; j < screen.Columns; j++ {
			pos := screen.Columns*i + j
			if screen.Cells[pos] != pScreen.Cells[pos] {
				diff[i] = true
				loc[i] = j
				break
			}
			w := screen.Cells[pos].Width
			if w > 0 {
				j += w - 1
			}
		}
	}
	return
}
