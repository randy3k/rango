package prompt

import (
	"time"

	"github.com/randy3k/rango/infchan"
)

type KeyBindDispatch struct {
	Binding *KeyBinding
	Data []rune
}

type KeyProcessor struct {
	bindings *KeyBindings
	kp   *infchan.InfChan
	event chan *KeyBindDispatch
	flush chan struct{}
	quit chan struct{}

	timeoutlen time.Duration
}

func NewKeyProcessor(bindings *KeyBindings) *KeyProcessor {
	p := &KeyProcessor{
		bindings: bindings.Normalize(),
		kp: infchan.NewInfChan(),
		event: make(chan *KeyBindDispatch),
		flush: make(chan struct{}),
		quit: make(chan struct{}),
		timeoutlen: time.Second,
	}
	return p
}

func (p *KeyProcessor) Stop() {
	close(p.quit)
}

func (p *KeyProcessor) Start() <-chan *KeyBindDispatch{
	go p.processLoop()
	return p.event
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
			p.event <- &KeyBindDispatch{Binding: binding}
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
				p.event <- &KeyBindDispatch{Binding: binding}
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
