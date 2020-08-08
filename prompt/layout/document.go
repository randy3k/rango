package layout

import (
	"regexp"
	"strings"
	// "github.com/mattn/go-runewidth"
)

var wordPatternStart = regexp.MustCompile(`(\pL+|\d+|[[:punct:]]+)\s*$`)
var wordPatternEnd = regexp.MustCompile(`^\s*(\pL+|\d+|[[:punct:]]+)`)


type DocumentCursor struct {
	Line      int
	Character int
}

type Document struct {
	Lines  [][]rune
	Cursor DocumentCursor
}

func NewDocument() *Document {
	return &Document{
		Lines: make([][]rune, 1),
	}
}

func (doc *Document) SetText(t string) {
	lines := strings.Split(t, "\n")
	doc.Lines = make([][]rune, len(lines))
	for i, l := range lines {
		doc.Lines[i] = []rune(l)
	}
}

func (doc *Document) GetText() string {
	lines := make([]string, len(doc.Lines))
	for i, l := range doc.Lines {
		lines[i] = string(l)
	}
	return strings.Join(lines, "\n")
}

func (doc *Document) MoveCursorUp() {
	if doc.Cursor.Line > 0 {
		doc.Cursor.Line -= 1
		text := doc.Lines[doc.Cursor.Line]
		if doc.Cursor.Character > len(text) {
			doc.Cursor.Character = len(text)
		}
	}
}

func (doc *Document) MoveCursorDown() {
	if doc.Cursor.Line < len(doc.Lines)-1 {
		doc.Cursor.Line += 1
		text := doc.Lines[doc.Cursor.Line]
		if doc.Cursor.Character > len(text) {
			doc.Cursor.Character = len(text)
		}
	}
}

func (doc *Document) MoveCursorLeft() {
	if doc.Cursor.Character > 0 {
		doc.Cursor.Character -= 1
	} else if doc.Cursor.Line > 0 {
		line := doc.Cursor.Line - 1
		doc.Cursor.Line = line
		doc.Cursor.Character = len([]rune(doc.Lines[line]))
	}
}

func (doc *Document) MoveCursorRight() {
	line := doc.Cursor.Line
	if doc.Cursor.Character < len([]rune(doc.Lines[line])) {
		doc.Cursor.Character += 1
	} else if doc.Cursor.Line+1 < len(doc.Lines) {
		doc.Cursor.Character = 0
		doc.Cursor.Line++
	}
}

func (doc *Document) MoveCursorByWordLeft() {
	line := doc.Cursor.Line
	character := doc.Cursor.Character
	if character == 0 {
		doc.MoveCursorLeft()
	} else {
		text := doc.Lines[line]
		stext := strings.TrimRight(string(text[:character]), " ")
		allIndexes := wordPatternStart.FindAllStringIndex(stext, -1)
		if len(allIndexes) > 0 {
			loc := allIndexes[len(allIndexes)-1]
			textbefore := stext[:loc[0]]
			doc.Cursor.Character = len([]rune(textbefore))
		}
	}
}

func (doc *Document) MoveCursorByWordRight() {
	line := doc.Cursor.Line
	character := doc.Cursor.Character
	text := doc.Lines[line]
	if character < len(text) {
		text := doc.Lines[line]
		stext := strings.TrimRight(string(text[character:]), " ")
		allIndexes := wordPatternEnd.FindAllStringIndex(stext, -1)
		if len(allIndexes) > 0 {
			loc := allIndexes[0]
			textafter := stext[:loc[1]]
			doc.Cursor.Character += len([]rune(textafter))
		}
	} else {
		doc.MoveCursorRight()
	}
}

func (doc *Document) MoveCursorToBeginningOfLine() {
	doc.Cursor.Character = 0
}

func (doc *Document) MoveCursorToEndOfLine() {
	doc.Cursor.Character = len([]rune(doc.Lines[doc.Cursor.Line]))
}

func (doc *Document) InsertText(t string) {
	line := doc.Cursor.Line
	character := doc.Cursor.Character
	text := doc.Lines[line]
	doc.Lines[line] = []rune(string(text[:character]) + t + string(text[character:]))
	doc.Cursor.Character += len([]rune(t))
}

func (doc *Document) DeleteLeftRune() {
	line := doc.Cursor.Line
	character := doc.Cursor.Character
	if character > 0 {
		text := doc.Lines[line]
		doc.Lines[line] = []rune(string(text[:character-1]) + string(text[character:]))
		doc.MoveCursorLeft()
	} else if line > 0 {
		n := len([]rune(doc.Lines[line-1]))
		doc.Lines[line-1] = []rune(string(doc.Lines[line-1]) + string(doc.Lines[line]))
		doc.Lines = append(doc.Lines[:line], doc.Lines[line+1:]...)
		doc.Cursor.Line = line - 1
		doc.Cursor.Character = n
	}
}

func (doc *Document) DeleteRightRune() {
	line := doc.Cursor.Line
	character := doc.Cursor.Character
	text := doc.Lines[line]
	if character < len(text) {
		text := doc.Lines[line]
		doc.Lines[line] = []rune(string(text[:character]) + string(text[character+1:]))
	} else if line+1 < len(doc.Lines) {
		doc.Lines[line] = []rune(string(doc.Lines[line]) + string(doc.Lines[line+1]))
		doc.Lines = append(doc.Lines[:line+1], doc.Lines[line+2:]...)
	}
}

func (doc *Document) DeleteWord() {
	line := doc.Cursor.Line
	character := doc.Cursor.Character
	if character == 0 {
		doc.DeleteLeftRune()
	} else {
		text := doc.Lines[line]
		stext := strings.TrimRight(string(text[:character]), " ")
		allIndexes := wordPatternStart.FindAllStringIndex(stext, -1)
		if len(allIndexes) > 0 {
			loc := allIndexes[len(allIndexes)-1]
			textbefore := stext[:loc[0]]
			doc.Lines[line] = []rune(textbefore + string(text[character:]))
			doc.Cursor.Character = len([]rune(textbefore))
		}
	}
}

func (doc *Document) InsertLine() {
	line := doc.Cursor.Line
	character := doc.Cursor.Character

	text := doc.Lines[line]
	doc.Lines = append(doc.Lines[:line+1], doc.Lines[line:]...)
	doc.Lines[line] = text[:character]
	doc.Lines[line+1] = text[character:]

	doc.Cursor.Line++
	doc.Cursor.Character = 0
}
