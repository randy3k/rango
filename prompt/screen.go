package prompt

import (
	"strings"
)

type ScreenCursor struct {
	Line int
	Column int
}

type Screen struct {
	Lines int
	Columns int
	chars []Char
	eol []bool
	dirty []bool

	ScreenCursor
}


func NewScreen(h int, w int) *Screen {
	return &Screen{
		Columns: w,
		Lines: h,
		chars: make([]Char, w*h),
		eol: make([]bool, h),
		dirty: make([]bool, h),
	}
}

func (scr *Screen) String() string {
	s := ""
	for i := 0; i < scr.Lines; i++ {
		line := make([]string, scr.Columns)
		for j := 0; j < scr.Columns; {
			pos := scr.Columns * i + j
			c := scr.chars[pos]
			if c.Value == 0 {
				line[j] = " "
				j += 1
			} else  {
				line[j] = string(c.Value)
				j = j + c.Width
			}
		}
		s += "|" + strings.Join(line, "")
		if scr.eol[i] {
			s += "|\n"
		} else  {
			s += "+\n"
		}
	}
	return s
}


func (scr *Screen) Reset() {
	scr.chars = make([]Char, scr.Columns * scr.Lines)
	scr.eol = make([]bool, scr.Lines)
	scr.dirty = make([]bool, scr.Lines)
}


func (scr *Screen) Feed(c Char) (int, int) {
	line := scr.Line
	col := scr.Column
	if scr.Column >= scr.Columns {
		scr.eol[line] = true
		scr.LineFeed()
	}
	pos := scr.Columns * scr.Line + scr.Column
	scr.chars[pos] = c
	scr.dirty[scr.Line] = true
	scr.Column += c.Width
	return line, col
}

func (scr *Screen) LineFeed() {
	scr.Column = 0
	scr.Line += 1
	if scr.Line == scr.Lines {
		scr.chars = append(scr.chars[scr.Columns:], make([]Char, scr.Columns)...)
		scr.eol = append(scr.eol[1:], false)
		scr.dirty = append(scr.dirty[1:], false)
		scr.Line -= 1
	}
}

func (scr *Screen) GoTo(line int, col int) {
	scr.Line = max(0, min(line, scr.Lines - 1))
	scr.Column = max(0, min(col, scr.Columns - 1))
}

func (scr *Screen) SetCharAt(line int, col int, c Char) {
	oldline := scr.Line
	oldcol := scr.Column
	scr.GoTo(line, col)
	scr.Feed(c)
	scr.GoTo(oldline, oldcol)
}

func (scr *Screen) SetCharsAt(line int, col int, cs []Char, eol bool) {
	oldline := scr.Line
	oldcol := scr.Column
	scr.GoTo(line, col)
	for _, c := range cs {
		scr.Feed(c)
	}
	scr.eol[scr.Line] = eol
	scr.GoTo(oldline, oldcol)
}


func (scr *Screen) IsLineEmpty(line int) bool {
	for j := 0; j < scr.Columns; j++ {
		if scr.chars[line * scr.Columns + j].Value > 0 {
			return false
		}
	}
	return true
}
