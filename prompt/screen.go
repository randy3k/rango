package prompt

import (
	"strings"
)

type ScreenCursor struct {
	row int
	col int
}

type Screen struct {
	h int
	w int
	chars []Char
	eol []bool
	Dirty []bool

	ScreenCursor
}


func NewScreen(h int, w int) *Screen {
	return &Screen{
		w: w,
		h: h,
		chars: make([]Char, w*h),
		eol: make([]bool, h),
		Dirty: make([]bool, h),
	}
}

func (scr *Screen) String() string {
	s := ""
	for i := 0; i < scr.h; i++ {
		line := make([]string, scr.w)
		for j := 0; j < scr.w; {
			pos := scr.w * i + j
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


func (scr *Screen) Clear() {
	scr.chars = make([]Char, scr.w * scr.h)
	scr.eol = make([]bool, scr.h)
	for i := 0; i < scr.h; i++ {
		scr.Dirty[i] = true
	}
}


func (scr *Screen) Feed(c Char) (int, int) {
	row := scr.row
	col := scr.col
	if scr.col >= scr.w {
		scr.eol[row] = true
		scr.LineFeed()
	}
	pos := scr.w * scr.row + scr.col
	scr.chars[pos] = c
	scr.Dirty[scr.row] = true
	scr.col += c.Width
	return row, col
}

func (scr *Screen) LineFeed() {
	scr.col = 0
	scr.row += 1
	if scr.row == scr.h {
		scr.chars = append(scr.chars[scr.w:], make([]Char, scr.w)...)
		scr.eol = append(scr.eol[1:], false)
		scr.Dirty = append(scr.Dirty[1:], false)
		scr.row -= 1
	}
}

func (scr *Screen) GoTo(row int, col int) {
	scr.row = max(0, min(row, scr.h - 1))
	scr.col = max(0, min(col, scr.w - 1))
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
