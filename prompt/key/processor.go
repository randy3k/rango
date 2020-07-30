package key

import (
	"time"

	"github.com/randy3k/rango/infchan"
)

type KeyBindingDispatch struct {
	Binding *KeyBinding
	Data []rune
}

type KeyProcessor struct {
	bindings *KeyBindings
	kp   *infchan.InfChan
	dispatch chan *KeyBindingDispatch
	flush chan struct{}
	quit chan struct{}

	timeoutlen time.Duration
}

func NewKeyProcessor(bindings *KeyBindings) *KeyProcessor {
	p := &KeyProcessor{
		bindings: bindings.Normalize(),
		kp: infchan.NewInfChan(),
		dispatch: make(chan *KeyBindingDispatch),
		flush: make(chan struct{}),
		quit: make(chan struct{}),
		timeoutlen: time.Second,
	}
	return p
}

func (p *KeyProcessor) Stop() {
	close(p.quit)
}

func (p *KeyProcessor) Start() <-chan *KeyBindingDispatch{
	go p.processLoop()
	return p.dispatch
}

func (p *KeyProcessor) Feed(kp *KeyPress) {
	p.kp.In <- kp
}

func (p *KeyProcessor) Flush() {
	p.flush <- struct{}{}
}

func (p *KeyProcessor) processLoop() {
	prefix := []Key{}
	retry := false
	flushing := false
	flushTimer := time.AfterFunc(p.timeoutlen, p.Flush)
	flushTimer.Stop()
loop:
	for {
		if retry {
			retry = false
		} else {
			select {
			case <-p.quit:
				break loop
			case <-p.flush:
				// stop flushTimer in case of mannual flush
				flushTimer.Stop()
				if len(prefix) == 0 {
					// nothing to flush
					continue
				}
				flushing = true
			case kp := <-p.kp.Out:
				prefix = append(prefix, kp.(*KeyPress).Key)
			}
		}

		// eager keybindings
		if binding, ok := p.bindings.Get(prefix, true); ok {
			data := []rune{}
			if len(binding.Keys) == 1 && binding.Keys[0] == Any {
				data = []rune(prefix[0])
			}
			p.dispatch <- &KeyBindingDispatch{Binding: binding, Data: data}
			prefix = nil
			flushing = false
			continue
		}

		if !flushing && p.bindings.HasPrefix(prefix) {
			flushTimer.Reset(p.timeoutlen)
			continue
		}

		found := -1

		// longest match
		for i := len(prefix); i > 0; i-- {
			if binding, ok := p.bindings.Get(prefix[:i], false); ok {
				data := []rune{}
				if len(binding.Keys) == 1 && binding.Keys[0] == Any {
					data = []rune(prefix[0])
				}
				p.dispatch <- &KeyBindingDispatch{Binding: binding, Data: data}
				found = i
				break
			}
		}

		if found >= 0 {
			prefix = prefix[found:]
			if len(prefix) > 0 {
				retry = true
			}
		} else if len(prefix) > 0 {
			// discard the first key
			prefix = prefix[1:]
			if len(prefix) > 0 {
				retry = true
			}
		}

		if flushing && len(prefix) == 0 {
			flushing = false
		}
	}
}
