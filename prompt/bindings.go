package prompt

import (
	. "github.com/randy3k/rango/prompt/key"
)

func (p *Prompt) Bindings() *KeyBindings {
	b := KeyBindings{
		// {
		// 	Keys: []Key{Escape},
		// 	Handler: func(event *Event) {
		// 		print("[escape]\r\n")
		// 	},
		// },
		// {
		// 	Keys: []Key{Escape, "b"},
		// 	Handler: func(event *Event) {
		// 		myprint(event.Keys)
		// 	},
		// },
		// {
		// 	Keys: []Key{Escape, "f"},
		// 	Handler: func(event *Event) {
		// 		myprint(event.Keys)
		// 	},
		// },
		// {
		// 	Keys: []Key{ControlA},
		// 	Handler: func(event *Event) {
		// 		print("[c-a]\r\n")
		// 	},
		// },
		{
			Keys: []Key{Enter},
			Handler: func(event *Event) {
				event.Prompt.Buffer.Document.InsertLine()
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
				if (len(event.Data) == 1) {
					event.Prompt.Buffer.InsertText(string(event.Data[0]))
				}
			},
		},
	}

	return &b
}
