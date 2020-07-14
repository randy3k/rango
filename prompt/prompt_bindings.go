package prompt

func PromptBindings(p *Prompt) *KeyBindings {
	b := KeyBindings{
		Bindings: []*Binding{
			{
				Keys: []Key{Escape},
				Handler: func(key []Key, data []rune) {
					print("esc\r\n")
				},
			},
			{
				Keys: []Key{Escape, Any},
				Handler: func(key []Key, data []rune) {
					printf("%v\r\n", key)
				},
			},
			{
				Keys: []Key{ControlA},
				Handler: func(key []Key, data []rune) {
					print("[c-a]\r\n")
				},
			},
			{
				Keys: []Key{Enter},
				Handler: func(key []Key, data []rune) {
					print("\r\n")
				},
			},
			{
				Keys: []Key{"q"},
				Handler: func(key []Key, data []rune) {
					print("bye!\r\n")
					p.quit = true
				},
			},
			// {
			// 	Keys: []Key{Any},
			// 	Handler: func(key []Key, data []rune) {
			// 		for _, k := range key {
			// 			print(k)
			// 		}
			// 	},
			// },
		},
	}

	return &b
}
