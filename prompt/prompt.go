package prompt

// ultimately, we will move this package out of rango

import (
	// "fmt"
	// "runtime"
	"github.com/alecthomas/chroma"
	. "github.com/randy3k/rango/prompt/key"
	. "github.com/randy3k/rango/prompt/layout"
	. "github.com/randy3k/rango/prompt/renderer"
	. "github.com/randy3k/rango/prompt/terminal"
)

type Prompt struct {
	quit   bool
	window *Window
	messageFunc func() string
	messageContinuationFunc func(int) string
	lexer chroma.Lexer
	style *chroma.Style
}

func NewPrompt(options ...Option) *Prompt {
	p := &Prompt{}
	for _, option := range options {
		option(p)
	}
	return p
}


func (p *Prompt) Quit() {
	p.quit = true
}

func (p *Prompt) Show() {
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

	renderer := NewRenderer(t)

	// TODO: persistent window
	p.window = NewWindow(
		NewBuffer(p.lexer, p.style),
		NewPrefix(p.messageFunc, p.messageContinuationFunc),
	)

	_redraw := func() {
		screen := NewScreen(t.Lines, t.Columns)
		p.window.WriteToScreen(screen)
		renderer.Render(screen)
	}
	_redraw()

	// loop:
	for !p.quit {
		// caution: the case handler must not block
		select {
		case dispatch := <-kbDispatch:
			handler := dispatch.Binding.Handler.(func(*Event))
			handler(NewEvent(dispatch.Binding.Keys, dispatch.Data, p))
			_redraw()
		case kp := <-keyPress:
			kProcessor.Feed(kp)
		case input := <-kbInput:
			kParser.Feed(input)
		}
	}
}
