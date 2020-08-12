package prompt

import (
	. "github.com/randy3k/rango/prompt/key"
)

func (p *Prompt) Bindings() *KeyBindings {
	b := KeyBindings{
		{
			Keys: []Key{Up},
			Handler: func(event *Event) {
				event.Buffer.Document.MoveCursorUp()
			},
		},
		{
			Keys: []Key{Down},
			Handler: func(event *Event) {
				event.Buffer.Document.MoveCursorDown()
			},
		},
		{
			Keys: []Key{Right},
			Handler: func(event *Event) {
				event.Buffer.Document.MoveCursorRight()
			},
		},
		{
			Keys: []Key{Left},
			Handler: func(event *Event) {
				event.Buffer.Document.MoveCursorLeft()
			},
		},
		{
			Keys: []Key{Enter},
			Handler: func(event *Event) {
				event.Buffer.Document.InsertLine()
			},
		},
		{
			Keys: []Key{BackSpace},
			Handler: func(event *Event) {
				event.Buffer.Document.DeleteLeftRune()
			},
		},
		{
			Keys: []Key{Delete},
			Handler: func(event *Event) {
				event.Buffer.Document.DeleteRightRune()
			},
		},
		{
			Keys: []Key{Escape, BackSpace},
			Handler: func(event *Event) {
				event.Buffer.Document.DeleteWord()
			},
		},
		{
			Keys: []Key{Escape, "b"},
			Handler: func(event *Event) {
				event.Buffer.Document.MoveCursorByWordLeft()
			},
		},
		{
			Keys: []Key{Escape, "f"},
			Handler: func(event *Event) {
				event.Buffer.Document.MoveCursorByWordRight()
			},
		},
		{
			Keys: []Key{ControlA},
			Handler: func(event *Event) {
				event.Buffer.Document.MoveCursorToBeginningOfLine()
			},
		},
		{
			Keys: []Key{ControlE},
			Handler: func(event *Event) {
				event.Buffer.Document.MoveCursorToEndOfLine()
			},
		},
		{
			Keys: []Key{"q"},
			Handler: func(event *Event) {
				// print("bye!\r\n")
				p.Quit()
			},
		},
		{
			Keys: []Key{Any},
			Handler: func(event *Event) {
				if len(event.Data) == 1 {
					event.Buffer.InsertText(string(event.Data[0]))
				}
			},
		},
	}

	return &b
}
