package char

import (
	"github.com/alecthomas/chroma"
	"github.com/mattn/go-runewidth"
)

type Attributes struct {
	Foreground Color // color in hex + 1, so 0 is default
	Background Color // color in hex + 1, so 0 is default
	Bold       bool
	Italic     bool
	Underline  bool
	Blink      bool
	Reverse    bool
	Hidden     bool
}

type Char struct {
	Value rune
	Width int
	Continuation   bool
	Attributes
}

var DefaultAttributes Attributes = Attributes{}
var DefaultChar Char = Char{Value: ' ', Width: 1, Attributes: DefaultAttributes}

func NewChar(r rune, attr Attributes) Char {
	return Char{
		Value:      r,
		Width:      runewidth.RuneWidth(r),
		Attributes: attr,
	}
}

func (c Char) String() string {
	return string(c.Value)
}

func ChromaStyleToAttributes(sty chroma.StyleEntry) Attributes {
	return Attributes{
		Foreground: Color(sty.Colour),
		Bold:       sty.Bold == chroma.Yes,
		Italic:     sty.Italic == chroma.Yes,
		Underline:  sty.Underline == chroma.Yes,
	}
}

type Chars []Char

func (chars Chars) Width() int {
	w := 0
	for _, c := range chars {
		w += c.Width
	}
	return w
}

func (chars Chars) SplitBy(r rune) []Chars {
	lines := make([]Chars, 0)
	last := 0
	for i, c := range chars {
		if c.Value == r {
			lines = append(lines, chars[last:i])
			last = i + 1
		}
	}
	lines = append(lines, chars[last:])
	return lines
}

func (chars Chars) SplitEvery(width int) []Chars {
	if width <= 0 {
		panic("`at` should be >= 1")
	}
	lines := make([]Chars, 0)
	last := 0
	w := 0
	for i, c := range chars {
		w += c.Width
		if w > width {
			lines = append(lines, chars[last:i])
			last = i
			w = 0
		}
	}
	lines = append(lines, chars[last:])
	return lines
}
