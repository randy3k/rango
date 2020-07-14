package prompt

import (
	"sync"
	"time"

	"github.com/randy3k/rango/prompt/infchan"
)

type KeyPress struct {
	Key Key
	Data []rune
}

type Parser struct {
	paste_mode bool

	input     *infchan.InfChan
	output   chan *KeyPress
	flush      chan struct{}
	quit       chan struct{}
	ttimeoutlen time.Duration

	mu sync.Mutex
}

func NewParser() *Parser {
	p := &Parser{
		input: infchan.NewInfChan(),
		output: make(chan *KeyPress),
		flush: make(chan struct{}),
		quit: make(chan struct{}),
		ttimeoutlen: 50 * time.Millisecond,
	}
	return p
}

func (p *Parser) Stop() {
	close(p.quit)
	p.input.Close()
}

func (p *Parser) Start() <-chan *KeyPress {
	go p.parseLoop()
	return p.output
}

func (p *Parser) Feed(input []rune) {
	for _, x := range input {
		p.input.In <- x
	}
}

func (p *Parser) Flush() {
	p.flush <- struct{}{}
}

func (p *Parser) SetFlushTimeout(d int) {
	p.ttimeoutlen = time.Duration(d) * time.Millisecond
}


func (p *Parser) parseLoop() {
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
			if key, ok := ANSIKeyMap[string(prefix[:i])]; ok {
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
