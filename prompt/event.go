package prompt

import (
	. "github.com/randy3k/rango/prompt/key"
	. "github.com/randy3k/rango/prompt/layout"
)

type Event struct {
	Keys   []Key
	Data   []rune
	Prompt *Prompt
	Window *Window
	Buffer *Buffer
}


func NewEvent(keys []Key, data []rune, p *Prompt) *Event {
	return &Event{
		Keys: keys,
		Data: data,
		Prompt: p,
		Window: p.window,
		Buffer: p.window.Buffer,
	}
}
