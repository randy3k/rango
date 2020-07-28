package prompt

import (
	// "bytes"
	// "os"
	"github.com/alecthomas/chroma"
)

type Buffer struct {
	// completer
	Doc   *Document
	Lexer *chroma.Lexer
	Style *chroma.Style
}

func NewBuffer(lexer chroma.Lexer, style *chroma.Style) *Buffer {
	buf := &Buffer{
		Doc:   &Document{},
		Lexer: &lexer,
		Style: style,
	}
	return buf
}

func (buf *Buffer) SetText(t string) {
	buf.Doc.SetText(t)
}

func (buf *Buffer) GetChars() Chars {
	it, err := (*buf.Lexer).Tokenise(nil, buf.Doc.GetText())
	chars := make(Chars, 0)
	if err != nil {
		for _, x := range buf.Doc.GetText() {
			chars = append(chars, NewChar(x, DefaultAttributes))
		}
	} else {
		for token := it(); token != chroma.EOF; token = it() {
			value := []rune(token.Value)
			style := buf.Style.Get(token.Type.Category())
			for _, x := range value {
				chars = append(chars, NewChar(x, ChromaStyleToAttributes(style)))
			}
		}
	}
	return chars
}


func (buf *Buffer) CreateContent(width, height int) *Content {
	lines := buf.GetChars().SplitBy('\n')
	return NewContent(lines, buf.Doc.Cursor)
}
