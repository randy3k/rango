package prompt

import (
	"time"
)

type KeyPressEvent struct {
	Keys []Key
	Data []rune
	Handler BindingHandler
}

type KeyProcessor struct {
	bindings *KeyBindings
	kp   chan *KeyPress
	event chan *KeyPressEvent
	flush chan struct{}
	quit chan struct{}

	timeoutlen time.Duration
}

func NewKeyProcessor(bindings *KeyBindings) *KeyProcessor {
	p := &KeyProcessor{
		bindings: bindings.Normalize(),
		kp: make(chan *KeyPress),
		event: make(chan *KeyPressEvent),
		flush: make(chan struct{}),
		quit: make(chan struct{}),
		timeoutlen: time.Second,
	}
	return p
}

func (p *KeyProcessor) Stop() {
	close(p.quit)
}

func (p *KeyProcessor) Start() <-chan *KeyPressEvent{
	go p.processLoop()
	return p.event
}

func (p *KeyProcessor) Feed(kp *KeyPress) {
	p.kp <- kp
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
			case kp := <-p.kp:
				prefix = append(prefix, kp.Key)
			}
		}

		// eager keybindings
		if binding, ok := p.bindings.Get(prefix, true); ok {
			p.event <- &KeyPressEvent{Keys: prefix, Handler: binding.Handler}
			prefix = nil
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
				p.event <- &KeyPressEvent{Keys: prefix, Handler: binding.Handler}
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
			// discard the first keypress
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
