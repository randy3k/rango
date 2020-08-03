package prompt

// ultimately, we will move this package out of rango

import (
	// "fmt"
	// "runtime"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	. "github.com/randy3k/rango/prompt/key"
)

type Prompt struct {
	Quit   bool
	Buffer *Buffer
	Window *Window
}

func NewPrompt() *Prompt {
	p := &Prompt{}
	return p
}

func (p *Prompt) Show(message string) {
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

	if p.Buffer == nil {
		lexer := lexers.Get("python")
		style := styles.Get("native")
		p.Buffer = NewBuffer(lexer, style)
	}
	buffer := p.Buffer

	if p.Window == nil {
		p.Window = NewWindow(buffer)
	}
	window := p.Window

	_redraw := func() {
		screen := NewScreen(t.Lines, t.Columns)
		window.WriteToScreen(screen)
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
