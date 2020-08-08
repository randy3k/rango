package prompt

import (
	"strings"
)

type ScreenCursor struct {
	Line   int
	Column int
}

type Screen struct {
	Lines   int
	Columns int
	Cursor  ScreenCursor
	chars   []Char
}

func NewScreen(h int, w int) *Screen {
	return &Screen{
		Columns: w,
		Lines:   h,
		chars:   make([]Char, w*h),
	}
}

func (screen *Screen) String() string {
	s := ""
	for i := 0; i < screen.Lines; i++ {
		line := make([]string, screen.Columns)
		for j := 0; j < screen.Columns; {
			pos := screen.Columns*i + j
			c := screen.chars[pos]
			if c.Value == 0 {
				line[j] = " "
				j += 1
			} else {
				line[j] = string(c.Value)
				j = j + c.Width
			}
		}
		s += "|" + strings.Join(line, "")
		if screen.IsLineEOL(i) {
			s += "|\n"
		} else {
			s += "+\n"
		}
	}
	return s
}

func (screen *Screen) Feed(c Char) (int, int) {
	line := screen.Cursor.Line
	col := screen.Cursor.Column
	if screen.Cursor.Column >= screen.Columns {
		screen.chars[screen.Columns*(line+1) - 1].EOL = true
		screen.LineFeed()
	}
	pos := screen.Columns*screen.Cursor.Line + screen.Cursor.Column
	screen.chars[pos] = c
	screen.Cursor.Column += c.Width
	return line, col
}

func (screen *Screen) IsLineEOL(line int) bool {
	return screen.chars[screen.Columns*(line+1) - 1].EOL
}

func (screen *Screen) LineFeed() {
	screen.Cursor.Column = 0
	screen.Cursor.Line += 1
	if screen.Cursor.Line == screen.Lines {
		screen.chars = append(screen.chars[screen.Columns:], make([]Char, screen.Columns)...)
		screen.Cursor.Line -= 1
	}
}

func (screen *Screen) GoTo(line int, col int) {
	screen.Cursor.Line = max(0, min(line, screen.Lines-1))
	screen.Cursor.Column = max(0, min(col, screen.Columns-1))
}

func (screen *Screen) SetCharAt(line int, col int, c Char) {
	oldline := screen.Cursor.Line
	oldcol := screen.Cursor.Column
	screen.GoTo(line, col)
	screen.Feed(c)
	screen.GoTo(oldline, oldcol)
}

func (screen *Screen) SetCharsAt(line int, col int, cs []Char) {
	oldline := screen.Cursor.Line
	oldcol := screen.Cursor.Column
	screen.GoTo(line, col)
	for _, c := range cs {
		screen.Feed(c)
	}
	screen.GoTo(oldline, oldcol)
}

func (screen *Screen) IsLineEmpty(line int) bool {
	for j := 0; j < screen.Columns; j++ {
		if screen.chars[line*screen.Columns+j].Value > 0 {
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
			if screen.chars[pos] != pScreen.chars[pos] {
				diff[i] = true
				loc[i] = j
				break
			}
		}
	}
	return
}
