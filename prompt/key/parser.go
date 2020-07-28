package key

import (
	"sync"
	"time"

	"github.com/randy3k/rango/infchan"
)

type KeyPress struct {
	Key Key
	Data []rune
}

type KeyParser struct {
	paste_mode bool

	input     *infchan.InfChan
	output   chan *KeyPress
	flush      chan struct{}
	quit       chan struct{}
	ttimeoutlen time.Duration

	mu sync.Mutex
}

func NewKeyParser() *KeyParser {
	p := &KeyParser{
		input: infchan.NewInfChan(),
		output: make(chan *KeyPress),
		flush: make(chan struct{}),
		quit: make(chan struct{}),
		ttimeoutlen: 50 * time.Millisecond,
	}
	return p
}

func (p *KeyParser) Stop() {
	close(p.quit)
	p.input.Close()
}

func (p *KeyParser) Start() <-chan *KeyPress {
	go p.parseLoop()
	return p.output
}

func (p *KeyParser) Feed(input []rune) {
	for _, x := range input {
		p.input.In <- x
	}
}

func (p *KeyParser) Flush() {
	p.flush <- struct{}{}
}

func (p *KeyParser) SetFlushTimeout(d int) {
	p.ttimeoutlen = time.Duration(d) * time.Millisecond
}


func (p *KeyParser) parseLoop() {
	prefix := []rune{}
	retry := false
	flushing := false
	flushTimer := time.AfterFunc(p.ttimeoutlen, p.Flush)
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
			case r := <-p.input.Out:
				flushTimer.Stop()
				prefix = append(prefix, r.(rune))
			}
		}

		if !flushing && isPrefixOfANSI(string(prefix)) {
			flushTimer.Reset(p.ttimeoutlen)
			continue
		}

		found := -1

		// longest match
		for i := len(prefix); i > 0; i-- {
			if key, ok := aNSIKeyMap[string(prefix[:i])]; ok {
				for _, v := range key {
					p.output <- &KeyPress{Key: v, Data: nil}
				}
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
			// consume the first rune finally
			p.output <- &KeyPress{Key: Key(prefix[0]), Data: nil}
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
