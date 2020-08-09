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
	Quit   bool
	Buffer *Buffer
	Window *Window
	message string
	messageFun func() string
	messageContinuationFun func(int) string
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

func (p *Prompt) Show() {
	p.Quit = false
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

	p.Buffer = NewBuffer(p.lexer, p.style)

	margin := &Margin{}
	p.Window = NewWindow(p.Buffer, margin)

	_redraw := func() {
		screen := NewScreen(t.Lines, t.Columns)
		p.Window.WriteToScreen(screen)
		renderer.Render(screen)
	}
	_redraw()

	// loop:
	for !p.Quit {
		// caution: the case handler must not block
		select {
		case dispatch := <-kbDispatch:
			hand := dispatch.Binding.Handler.(func(*Event))
			hand(&Event{Keys: dispatch.Binding.Keys, Data: dispatch.Data, Prompt: p})
			_redraw()
		case kp := <-keyPress:
			kProcessor.Feed(kp)
		case input := <-kbInput:
			kParser.Feed(input)
		}
	}
}
