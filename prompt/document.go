package prompt

import (
	"strings"
	// "github.com/mattn/go-runewidth"
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

func (doc *Document) MoveCursorLeft() {
	if doc.Cursor.Character > 0 {
		doc.Cursor.Character -= 1
	}
}

func (doc *Document) MoveCursorRight() {
	line := doc.Cursor.Line
	if doc.Cursor.Character < len([]rune(doc.Lines[line])) {
		doc.Cursor.Character += 1
	}
}

func (doc *Document) InsertText(t string) {
	line := doc.Cursor.Line
	character := doc.Cursor.Character
	text := []rune(doc.Lines[line])
	doc.Lines[line] = string(text[:character]) + t + string(text[character:])
	doc.Cursor.Character += len([]rune(t))
}

func (doc *Document) InsertLine() {
	line := doc.Cursor.Line
	character := doc.Cursor.Character

	text := []rune(doc.Lines[line])
	doc.Lines[line] = string(text[:character])
	afterLines := doc.Lines[line+1:]
	doc.Lines = append(append(doc.Lines[:line+1], string(text[character:])), afterLines...)
	doc.Cursor.Line++
	doc.Cursor.Character = 0
}
