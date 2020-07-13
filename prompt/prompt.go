package prompt

// ultimately, we will move this package out of rango

import (
	// "fmt"
	// "runtime"
)

type PromptApp struct {
	quit bool
}

func (p *PromptApp) Run(message string) {
	p.quit = false
	// go func() {
	// 	fmt.Printf("%v\r\n", runtime.NumGoroutine())
	// }()

	t := NewTerminal()
	kbInput := t.Start()
	defer t.Fini()

	parser := NewParser()
	keyPress := parser.Start()
	defer parser.Fini()

	kproc := NewKeyProcessor(DefaultBindings())
	kproc.Start()
	defer kproc.Fini()

	// loop:
	for !p.quit {
		select {
		case input := <-kbInput:
			parser.Feed(input)
		case kp := <-keyPress:
			kproc.Feed(kp)
		}
	}
}
