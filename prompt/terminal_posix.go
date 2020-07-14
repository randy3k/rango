// +build !windows

package prompt

import (
	"sync"
	"time"
	// "math"
	// "fmt"

	"golang.org/x/sys/unix"
	"github.com/mattn/go-tty"
	"github.com/xo/terminfo"
)

type Terminal struct {
	tty *tty.TTY
	ti  *terminfo.Terminfo

	cleanup func() error

	input chan []rune
	quit chan struct{}

	h   int // window height
	w   int // window weight

	mu sync.Mutex
}

func NewTerminal() *Terminal {
	t := &Terminal{}
	t.init()
	return t
}

func (t *Terminal) init() error {
	var e error

	t.input = make(chan []rune)
	t.quit = make(chan struct{})

	if t.tty, e = tty.Open(); e != nil {
		return e
	}

	if t.ti, e = terminfo.LoadFromEnv(); e != nil {
		return e
	}

	if t.w, t.h, e = t.GetWinSize(); e != nil {
		return e
	}

	return nil
}

func (t *Terminal) Stop() error {
	close(t.quit)
	t.tty.Close()
	t.cleanup()
	return nil
}

func (t *Terminal) Start() (<-chan []rune) {
	go t.inputLoop()
	t.cleanup, _ = t.tty.Raw()
	return t.input
}


func (t *Terminal) inputLoop() {
loop:
	for {
		select {
		case <-t.quit:
			break loop
		default:
			if t.WaitForInput(100) {
				r, e := t.tty.ReadRune()
				rs := []rune{r}
				for t.tty.Buffered() {
					r, e = t.tty.ReadRune()
					rs = append(rs, r)
				}
				if e != nil {
					close(t.input)
					break loop
				}
				select {
				case t.input <- rs:
				case <-t.quit:
					break loop
				}
			}
		}
	}
}


func (t *Terminal) GetWinSize() (int, int, error) {
	return t.tty.Size()
}

func (t *Terminal) Resize() {

}

func (t *Terminal) WriteString(s string) {
	t.tty.Output().WriteString(s)
}

func (t *Terminal) Goto(row, col int) {
	t.WriteString(t.ti.Goto(row, col))
}

func (t* Terminal) WaitForInput(timeout int) bool {
	rFdSet := &unix.FdSet{}
	fd := t.tty.Input().Fd()
	rFdSet.Set(int(fd))
	timeval := unix.NsecToTimeval(int64(time.Duration(timeout) * time.Millisecond))
	n, err := unix.Select(int(fd + 1), rFdSet, nil, nil, &timeval)
	if err == unix.EINTR {
	    return false
	}
	return n >= 0 && rFdSet.IsSet(int(fd))
}
