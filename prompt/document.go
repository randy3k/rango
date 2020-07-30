package prompt

import (
	"strings"
)

type DocumentCursor struct {
	Line      int
	Character int
}

type Document struct {
	Lines  []string
	N      int
	Cursor DocumentCursor
}

func NewDocument() *Document {
	return &Document{
		Lines: []string{""},
		N:     1,
	}
}

func (doc *Document) SetText(t string) {
	doc.Lines = strings.Split(t, "\n")
	doc.N = len(doc.Lines)
}

func (doc *Document) GetText() string {
	return strings.Join(doc.Lines, "\n")
}

func (doc *Document) InsertText(t string) {
	line := doc.Cursor.Line
	character := doc.Cursor.Character

	text := []rune(doc.Lines[line])
	insertText := []rune(t)
	doc.Lines[line] = string(append(append(text[:character], insertText...), text[character:]...))
	doc.Cursor.Character += len(insertText)
}

func (doc *Document) InsertLine() {
	line := doc.Cursor.Line
	character := doc.Cursor.Character

	text := []rune(doc.Lines[line])
	doc.Lines[line] = string(text[:character])
	doc.Lines = append(append(doc.Lines[:(line+1)], string(text[character:])), doc.Lines[line+1:]...)
	doc.Cursor.Line++
	doc.Cursor.Character = 0
}
