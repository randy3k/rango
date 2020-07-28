package prompt

import (
	"strings"
)

type Document struct {
	Lines []string
	N int
	Cursor Position
}

func (doc *Document) SetText(t string) {
	doc.Lines = strings.Split(t, "\n")
	doc.N = len(doc.Lines)
}

func (doc *Document) GetText() string {
	return strings.Join(doc.Lines, "\n")
}
