package prompt

// ultimately, we will move this package out of rango

import (
	// "fmt"
	// "runtime"
)

type Prompt struct {
	quit bool
}

func NewPrompt() *Prompt {
	p := &Prompt{}
	return p
}

func (p *Prompt) Show(message string) {
	p.quit = false
	// go func() {
	// 	printf("%v\r\n", runtime.NumGoroutine())
	// }()

	t := NewTerminal()
	kbInput := t.Start()
	defer t.Stop()

	parser := NewParser()
	keyPress := parser.Start()
	defer parser.Stop()

	kproc := NewKeyProcessor(PromptBindings(p))
	keyPressEvent := kproc.Start()
	defer kproc.Stop()

	// loop:
	for !p.quit {
		select {
		case ev := <-keyPressEvent:
			ev.Handler(ev.Keys, ev.Data)
		case kp := <-keyPress:
			kproc.Feed(kp)
		case input := <-kbInput:
			parser.Feed(input)
		}
	}
}
