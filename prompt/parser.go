package prompt

import (
	"sync"
	"time"
)

type Parser struct {
	ch   chan KeyPress
	paste_mode bool

	buffer        []KeyPress
	deliver    chan struct{}

	rune       chan rune
	quit       chan struct{}
	flush      chan struct{}
	ttimeoutlen time.Duration

	mu sync.Mutex
}

func NewParser() *Parser {
	p := &Parser{}
	p.Init()
	return p
}

func (p *Parser) Init() {
	p.ch = make(chan KeyPress)
	p.deliver = make(chan struct{})
	p.rune = make(chan rune)
	p.quit = make(chan struct{})
	p.flush = make(chan struct{})
	p.ttimeoutlen = time.Duration(50) * time.Millisecond
}

func (p *Parser) Fini() {
	close(p.quit)
}

func (p *Parser) Start() <-chan KeyPress {
	go p.parseLoop()
	go p.deliverLoop()
	return p.ch
}

func (p *Parser) Feed(input []rune) {
	for _, x := range input {
		p.rune <- x
	}
}

func (p *Parser) Flush() {
	p.flush <- struct{}{}
}

func (p *Parser) SetFlushTimeout(d int) {
	p.ttimeoutlen = time.Duration(d) * time.Millisecond
}


func (p *Parser) appendBuffer(kp KeyPress) {
	p.mu.Lock()
	p.buffer = append(p.buffer, kp)
	p.mu.Unlock()
	select {
	case p.deliver <- struct{}{}:
	default:
	}
}


func (p *Parser) parseLoop() {
	prefix := ""
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
				if prefix == "" {
					// nothing to flush
					continue
				}
				flushing = true
			case r := <-p.rune:
				flushTimer.Stop()
				prefix += string(r)
			}
		}

		if !flushing && is_ansi_prefix(prefix) {
			flushTimer.Reset(p.ttimeoutlen)
			continue
		}

		found := -1

		// longest match
		for i := len(prefix); i > 0; i-- {
			if key, ok := ANSIKeyMap[prefix[:i]]; ok {
				for _, v := range key {
					p.appendBuffer(KeyPress{Key: v, Data: nil})
				}
				found = i
				break
			}
		}

		if found >= 0 {
			prefix = prefix[found:]
			if prefix != "" {
				retry = true
			}
		} else if prefix != "" {
			// consume the first byte finally
			p.appendBuffer(KeyPress{Key: Key(prefix[0]), Data: nil})
			prefix = prefix[1:]
			if prefix != "" {
				retry = true
			}
		}

		if flushing && prefix == "" {
			flushing = false
		}
	}
}

// it is needed because writing to output channel may be blocking
func (p *Parser) deliverLoop() {
loop:
	for {
		select {
		case <-p.deliver:
			for {
				p.mu.Lock()
				buffer := make([]KeyPress, len(p.buffer))
				copy(buffer, p.buffer)
				p.buffer = nil
				p.mu.Unlock()
				for _, v := range buffer {
					select {
					case p.ch <- v:
						continue
					case <-p.quit:
						break loop
					}
				}
				p.mu.Lock()
				done := len(p.buffer) == 0
				p.mu.Unlock()
				if done {
					break
				}
			}
		case <-p.quit:
			break loop
		}
	}
}
