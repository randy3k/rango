package prompt

import (
	"strings"
)

type ScreenCursor struct {
	row int
	col int
}

type Screen struct {
	lines int
	columns int
	chars []Char
	eol []bool
	dirty []bool

	ScreenCursor
}


func NewScreen(h int, w int) *Screen {
	return &Screen{
		columns: w,
		lines: h,
		chars: make([]Char, w*h),
		eol: make([]bool, h),
		dirty: make([]bool, h),
	}
}

func (scr *Screen) String() string {
	s := ""
	for i := 0; i < scr.lines; i++ {
		line := make([]string, scr.columns)
		for j := 0; j < scr.columns; {
			pos := scr.columns * i + j
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
	scr.chars = make([]Char, scr.columns * scr.lines)
	scr.eol = make([]bool, scr.lines)
	scr.dirty = make([]bool, scr.lines)
}


func (scr *Screen) Feed(c Char) (int, int) {
	row := scr.row
	col := scr.col
	if scr.col >= scr.columns {
		scr.eol[row] = true
		scr.LineFeed()
	}
	pos := scr.columns * scr.row + scr.col
	scr.chars[pos] = c
	scr.dirty[scr.row] = true
	scr.col += c.Width
	return row, col
}

func (scr *Screen) LineFeed() {
	scr.col = 0
	scr.row += 1
	if scr.row == scr.lines {
		scr.chars = append(scr.chars[scr.columns:], make([]Char, scr.columns)...)
		scr.eol = append(scr.eol[1:], false)
		scr.dirty = append(scr.dirty[1:], false)
		scr.row -= 1
	}
}

func (scr *Screen) GoTo(row int, col int) {
	scr.row = max(0, min(row, scr.lines - 1))
	scr.col = max(0, min(col, scr.columns - 1))
}

func (scr *Screen) SetCharAt(row int, col int, c Char) {
	oldrow := scr.row
	oldcol := scr.col
	scr.GoTo(row, col)
	scr.Feed(c)
	scr.GoTo(oldrow, oldcol)
}

func (scr *Screen) SetCharsAt(row int, col int, cs []Char, eol bool) {
	oldrow := scr.row
	oldcol := scr.col
	scr.GoTo(row, col)
	for _, c := range cs {
		scr.Feed(c)
	}
	scr.eol[scr.row] = eol
	scr.GoTo(oldrow, oldcol)
}


func (scr *Screen) IsRowEmpty(row int) bool {
	for j := 0; j < scr.columns; j++ {
		if scr.chars[row * scr.columns + j].Value > 0 {
			return false
		}
	}
	return true
}
