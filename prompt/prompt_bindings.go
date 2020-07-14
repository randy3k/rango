package prompt

func (p *Prompt) Bindings() *KeyBindings {
	b := KeyBindings{
		{
			Keys: []Key{Escape},
			Handler: func(event *Event) {
				print("[escape]\r\n")
			},
		},
		{
			Keys: []Key{Escape, "b"},
			Handler: func(event *Event) {
				printf("%v\r\n", event.Keys)
			},
		},
		{
			Keys: []Key{Escape, "f"},
			Handler: func(event *Event) {
				printf("%v\r\n", event.Keys)
			},
		},
		{
			Keys: []Key{ControlA},
			Handler: func(event *Event) {
				print("[c-a]\r\n")
			},
		},
		{
			Keys: []Key{Enter},
			Handler: func(event *Event) {
				print("\r\n")
			},
		},
		{
			Keys: []Key{Left},
			Handler: func(event *Event) {
				printf("%v\r\n", event.Keys)
			},
		},
		{
			Keys: []Key{Right},
			Handler: func(event *Event) {
				printf("%v\r\n", event.Keys)
			},
		},
		{
			Keys: []Key{"q"},
			Handler: func(event *Event) {
				print("bye!\r\n")
				p.quit = true
			},
		},
		// {
		// 	Keys: []Key{Any},
		// 	Handler: func(event *Event) {
		// 		for _, k := range key {
		// 			print(k)
		// 		}
		// 	},
		// },
	}

	return &b
}
