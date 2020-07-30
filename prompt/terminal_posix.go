// +build !windows

package prompt

import (
	"os"
	"sync"
	"time"
	// "math"
	"fmt"

	"github.com/mattn/go-tty"
	"golang.org/x/sys/unix"
)

type Terminal struct {
	tty        *tty.TTY
	colorDepth ColorDepth

	cleanup func() error

	input chan []rune
	quit  chan struct{}

	Lines   int // window height
	Columns int // window weight

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

	t.colorDepth = t.ColorDepthFromTerm(os.Getenv("TERM"))

	if t.Columns, t.Lines, e = t.GetWinSize(); e != nil {
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

func (t *Terminal) Start() <-chan []rune {
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

func (t *Terminal) ColorDepthFromTerm(term string) ColorDepth {
	if term == "" || term == "dumb" || term == "unknown" {
		return ColorDepth1Bit
	} else if term == "xterm" || term == "xterm-color" || term == "xterm-16color" ||
			term == "linux" || term == "eterm-color" {
		return ColorDepth4Bits
	} else {
		return ColorDepth8Bits
	}
}

func (t *Terminal) GetWinSize() (int, int, error) {
	return t.tty.Size()
}

func (t *Terminal) Resize() {

}

func (t *Terminal) WriteString(s string) {
	// TODO: buffer stdout
	t.tty.Output().WriteString(s)
}

func (t *Terminal) WaitForInput(timeout int) bool {
	rFdSet := &unix.FdSet{}
	fd := t.tty.Input().Fd()
	rFdSet.Set(int(fd))
	timeval := unix.NsecToTimeval(int64(time.Duration(timeout) * time.Millisecond))
	n, err := unix.Select(int(fd+1), rFdSet, nil, nil, &timeval)
	if err == unix.EINTR {
		return false
	}
	return n >= 0 && rFdSet.IsSet(int(fd))
}

func (t *Terminal) HideCursor() { // DECTCEM
	t.WriteString("\x1b[?25l")
}

func (t *Terminal) ShowCursor() { // DECTCEM
	t.WriteString("\x1b[?25h")
}

func (t *Terminal) ColorSequence(c Char) string {
	s := ""
	if t.colorDepth == ColorDepth8Bits {
		if fg := c.Foreground.Code8Bits(); fg >= 0 {
			s += fmt.Sprintf("\x1b[38;5;%dm", fg)
		}
		if bg := c.Background.Code8Bits(); bg >= 0 {
			s += fmt.Sprintf("\x1b[48;5;%dm", bg)
		}
	} else if t.colorDepth == ColorDepth4Bits {
		if fg := c.Foreground.Code4Bits(); fg >= 0 {
			if fg <= 7 {
				s += fmt.Sprintf("\x1b[%dm", 30 + fg)
			} else {
				// high intensity
				s += fmt.Sprintf("\x1b[%dm", 90 + fg - 8)
			}
		}
		if bg := c.Background.Code4Bits(); bg >= 0 {
			if bg <= 7 {
				s += fmt.Sprintf("\x1b[%dm", 40 + bg)
			} else {
				// high intensity
				s += fmt.Sprintf("\x1b[%dm", 100 + bg - 8)
			}
		}
	}
	return s
}

func (t *Terminal) MoveCursorUp(x int) { // CUU
	if x > 0 {
		t.WriteString(fmt.Sprintf("\x1b[%dA", x))
	}
}

func (t *Terminal) MoveCursorDown(x int) { // CUD
	if x > 0 {
		t.WriteString(fmt.Sprintf("\x1b[%dB", x))
	}
}


func (t *Terminal) MoveCursorLeft(x int) { // CUF
	if x > 0 {
		t.WriteString(fmt.Sprintf("\x1b[%dC", x))
	}
}

func (t *Terminal) MoveCursorRight(x int) { // CUB
	if x > 0 {
		t.WriteString(fmt.Sprintf("\x1b[%dD", x))
	}
}
