package layout

import (
	// "bytes"
	// "os"
	. "github.com/randy3k/rango/prompt/char"

	"github.com/alecthomas/chroma"
)

type Buffer struct {
	// completer
	Document *Document
	lexer    *chroma.Lexer
	style    *chroma.Style
	prefixFunc func(int) string
}

func NewBuffer(lexer chroma.Lexer, style *chroma.Style,
		prefixFunc func(int) string) *Buffer {

	buf := &Buffer{
		Document: NewDocument(),
		lexer:    &lexer,
		style:    style,
		prefixFunc: prefixFunc,
	}
	return buf
}


func (buf *Buffer) SetText(t string) {
	buf.Document.SetText(t)
}

func (buf *Buffer) InsertText(t string) {
	buf.Document.InsertText(t)
}

func (buf *Buffer) Highlight() (chars Chars) {
	chars = make(Chars, 0)

	if buf.lexer == nil || buf.style == nil {
		for _, x := range buf.Document.GetText() {
			chars = append(chars, NewChar(x, DefaultAttributes))
		}
		return
	}

	it, err := (*buf.lexer).Tokenise(nil, buf.Document.GetText())
	if err != nil {
		for _, x := range buf.Document.GetText() {
			chars = append(chars, NewChar(x, DefaultAttributes))
		}
		return
	}

	for token := it(); token != chroma.EOF; token = it() {
		value := []rune(token.Value)
		var style chroma.StyleEntry
		if buf.style.Has(token.Type.SubCategory()) {
			style = buf.style.Get(token.Type.SubCategory())
		} else {
			style = buf.style.Get(token.Type.Category())
		}
		for _, x := range value {
			chars = append(chars, NewChar(x, ChromaStyleToAttributes(style)))
		}
	}
	return
}

func (buf *Buffer) CreateContent(width, height int) *Content {
	lines := buf.Highlight().SplitBy('\n')
	prefix := ANSI(buf.prefixFunc(0))
	prefixWidth := prefix.Width()

	for i := range lines {
		thisPrefix := ANSI(buf.prefixFunc(i))
		thisPrefixWidth := thisPrefix.Width()
		if thisPrefixWidth < prefixWidth {
			for i := 0; i < prefixWidth - thisPrefixWidth; i++ {
				thisPrefix = append(thisPrefix, NewChar(' ', DefaultAttributes))
			}
		} else if thisPrefixWidth > prefixWidth {
			thisPrefix = thisPrefix[:prefixWidth]
		}
		lines[i] = append(thisPrefix, lines[i]...)
	}
	return NewContent(lines, buf.Document.Cursor, prefixWidth)
}
