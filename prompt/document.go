package prompt

import (
	"strings"
)

type DocumentCursor struct {
    line int
    character int
}


type Document struct {
	Lines []string
	N int
	Cursor DocumentCursor
}

func (doc *Document) SetText(t string) {
	doc.Lines = strings.Split(t, "\n")
	doc.N = len(doc.Lines)
}

func (doc *Document) GetText() string {
	return strings.Join(doc.Lines, "\n")
}
