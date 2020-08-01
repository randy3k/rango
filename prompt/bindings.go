package prompt

import (
	. "github.com/randy3k/rango/prompt/key"
)

func (p *Prompt) Bindings() *KeyBindings {
	b := KeyBindings{
		{
			Keys: []Key{Up},
			Handler: func(event *Event) {
				event.Prompt.Buffer.Document.MoveCursorUp()
			},
		},
		{
			Keys: []Key{Down},
			Handler: func(event *Event) {
				event.Prompt.Buffer.Document.MoveCursorDown()
			},
		},
		{
			Keys: []Key{Right},
			Handler: func(event *Event) {
				event.Prompt.Buffer.Document.MoveCursorRight()
			},
		},
		{
			Keys: []Key{Left},
			Handler: func(event *Event) {
				event.Prompt.Buffer.Document.MoveCursorLeft()
			},
		},
		{
			Keys: []Key{Escape, Enter},
			Handler: func(event *Event) {
				event.Prompt.Buffer.Document.InsertLine()
			},
		},
		{
			Keys: []Key{BackSpace},
			Handler: func(event *Event) {
				event.Prompt.Buffer.Document.DeleteLeftRune()
			},
		},
		{
			Keys: []Key{Delete},
			Handler: func(event *Event) {
				event.Prompt.Buffer.Document.DeleteRightRune()
			},
		},
		{
			Keys: []Key{Escape, BackSpace},
			Handler: func(event *Event) {
				event.Prompt.Buffer.Document.DeleteWord()
			},
		},
		{
			Keys: []Key{"q"},
			Handler: func(event *Event) {
				// print("bye!\r\n")
				p.Quit = true
			},
		},
		{
			Keys: []Key{Any},
			Handler: func(event *Event) {
				if len(event.Data) == 1 {
					event.Prompt.Buffer.InsertText(string(event.Data[0]))
				}
			},
		},
	}

	return &b
}
