package infchan

import (
	"container/list"
)

// it is a channel which allows non blocking input and output
// the channel must be closed after done to avoid leaking goroutine

type InfChan struct {
	In chan interface{}
	Out chan interface{}
	quit chan struct{}
	l *list.List
}

func NewInfChan() *InfChan {
	ch := &InfChan{
		In: make(chan interface{}),
		Out: make(chan interface{}),
		quit: make(chan struct{}),
		l: list.New(),
	}
	go ch.loop()
	return ch
}

func (infch *InfChan) Close() {
	close(infch.quit)
}

func (infch *InfChan) loop() {
	n := 0
loop:
	for {
		if n > 0 {
			next := infch.l.Front()
			select {
			case in, ok := <-infch.In:
				if ok {
					infch.l.PushBack(in)
					n = n + 1
				} else {
					break loop
				}
			case infch.Out <- next.Value:
				infch.l.Remove(next)
				n = n - 1
			case <-infch.quit:
				break loop
			}
		} else {
			select {
			case in, ok := <-infch.In:
				if ok {
					infch.l.PushBack(in)
					n = n + 1
				} else {
					break loop
				}
			case <-infch.quit:
				break loop
			}
		}
	}
}
