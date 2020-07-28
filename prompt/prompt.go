package prompt

// ultimately, we will move this package out of rango

import (
	// "fmt"
	// "runtime"
	. "github.com/randy3k/rango/prompt/key"
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

	kParser := NewKeyParser()
	keyPress := kParser.Start()
	defer kParser.Stop()

	kProcessor := NewKeyProcessor(p.Bindings())
	kbDispatch := kProcessor.Start()
	defer kProcessor.Stop()

	// loop:
	for !p.quit {
		// caution: the case handler must not block
		select {
		case dispatch := <-kbDispatch:
			hand := dispatch.Binding.Handler.(func(*Event))
			hand(&Event{Keys: dispatch.Binding.Keys, Data: dispatch.Data, Prompt: p})
		case kp := <-keyPress:
			kProcessor.Feed(kp)
		case input := <-kbInput:
			kParser.Feed(input)
		}
	}
}
