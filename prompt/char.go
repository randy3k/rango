package prompt

import (
	"github.com/alecthomas/chroma"
	"github.com/mattn/go-runewidth"
)

type Attributes struct {
	Foreground int // color in hex + 1, so 0 is default
	Background int // color in hex + 1, so 0 is default
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
	Attributes
}

func NewChar(r rune, attr Attributes) Char {
	return Char{
		Value: r,
		Width: runewidth.RuneWidth(r),
		Attributes: attr,
	}
}

func (c Char) String() string {
	return string(c.Value)
}

var DefaultAttributes Attributes
var DefaultChar Char

func init() {
	DefaultAttributes = Attributes{}
	DefaultChar = NewChar(' ', DefaultAttributes)
}


func ChromaStyleToAttributes(sty chroma.StyleEntry) Attributes {
	return Attributes{
		Foreground: int(sty.Colour),
		Background: int(sty.Background),
		Bold:       sty.Bold == chroma.Yes,
		Italic:     sty.Italic == chroma.Yes,
		Underline:  sty.Underline == chroma.Yes,
	}
}

type Chars []Char

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


func (chars Chars) SplitAt(at int) []Chars {
	if (at <= 0) {
		panic("`at` should be >= 1")
	}
	lines := make([]Chars, 0)
	last := 0
	w := 0
	for i, c := range chars {
		if w >= at {
			lines = append(lines, chars[last:i])
			last = i
			w = 0
		}
		w += c.Width
	}
	lines = append(lines, chars[last:])
	return lines
}
