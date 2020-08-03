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
	eol     []bool
}

func NewScreen(h int, w int) *Screen {
	return &Screen{
		Columns: w,
		Lines:   h,
		chars:   make([]Char, w*h),
		eol:     make([]bool, h),
	}
}

func (scr *Screen) String() string {
	s := ""
	for i := 0; i < scr.Lines; i++ {
		line := make([]string, scr.Columns)
		for j := 0; j < scr.Columns; {
			pos := scr.Columns*i + j
			c := scr.chars[pos]
			if c.Value == 0 {
				line[j] = " "
				j += 1
			} else {
				line[j] = string(c.Value)
				j = j + c.Width
			}
		}
		s += "|" + strings.Join(line, "")
		if scr.eol[i] {
			s += "|\n"
		} else {
			s += "+\n"
		}
	}
	return s
}

func (scr *Screen) Feed(c Char) (int, int) {
	line := scr.Cursor.Line
	col := scr.Cursor.Column
	if scr.Cursor.Column >= scr.Columns {
		scr.eol[line] = true
		scr.LineFeed()
	}
	pos := scr.Columns*scr.Cursor.Line + scr.Cursor.Column
	scr.chars[pos] = c
	scr.Cursor.Column += c.Width
	return line, col
}

func (scr *Screen) LineFeed() {
	scr.Cursor.Column = 0
	scr.Cursor.Line += 1
	if scr.Cursor.Line == scr.Lines {
		scr.chars = append(scr.chars[scr.Columns:], make([]Char, scr.Columns)...)
		scr.eol = append(scr.eol[1:], false)
		scr.Cursor.Line -= 1
	}
}

func (scr *Screen) GoTo(line int, col int) {
	scr.Cursor.Line = max(0, min(line, scr.Lines-1))
	scr.Cursor.Column = max(0, min(col, scr.Columns-1))
}

func (scr *Screen) SetCharAt(line int, col int, c Char) {
	oldline := scr.Cursor.Line
	oldcol := scr.Cursor.Column
	scr.GoTo(line, col)
	scr.Feed(c)
	scr.GoTo(oldline, oldcol)
}

func (scr *Screen) SetCharsAt(line int, col int, cs []Char, eol bool) {
	oldline := scr.Cursor.Line
	oldcol := scr.Cursor.Column
	scr.GoTo(line, col)
	for _, c := range cs {
		scr.Feed(c)
	}
	scr.eol[scr.Cursor.Line] = eol
	scr.GoTo(oldline, oldcol)
}

func (scr *Screen) IsLineEmpty(line int) bool {
	for j := 0; j < scr.Columns; j++ {
		if scr.chars[line*scr.Columns+j].Value > 0 {
			return false
		}
	}
	return true
}
