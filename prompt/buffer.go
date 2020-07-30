package prompt

import (
	// "bytes"
	// "os"
	"github.com/alecthomas/chroma"
)

type Buffer struct {
	// completer
	Document   *Document
	Lexer *chroma.Lexer
	Style *chroma.Style
}

func NewBuffer(lexer chroma.Lexer, style *chroma.Style) *Buffer {
	buf := &Buffer{
		Document: NewDocument(),
		Lexer: &lexer,
		Style: style,
	}
	return buf
}

func (buf *Buffer) SetText(t string) {
	buf.Document.SetText(t)
}

func (buf *Buffer) InsertText(t string) {
	buf.Document.InsertText(t)
}

func (buf *Buffer) GetChars() Chars {
	it, err := (*buf.Lexer).Tokenise(nil, buf.Document.GetText())
	chars := make(Chars, 0)
	if err != nil {
		for _, x := range buf.Document.GetText() {
			chars = append(chars, NewChar(x, DefaultAttributes))
		}
	} else {
		for token := it(); token != chroma.EOF; token = it() {
			value := []rune(token.Value)
			var style chroma.StyleEntry
			if buf.Style.Has(token.Type.SubCategory()) {
				style = buf.Style.Get(token.Type.SubCategory())
			} else {
				style = buf.Style.Get(token.Type.Category())
			}
			for _, x := range value {
				chars = append(chars, NewChar(x, ChromaStyleToAttributes(style)))
			}
		}
	}
	return chars
}


func (buf *Buffer) CreateContent(width, height int) *Content {
	lines := buf.GetChars().SplitBy('\n')
	return NewContent(lines, buf.Document.Cursor)
}
