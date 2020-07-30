// +build !windows

package prompt

import (
	"os"
	"sync"
	"time"
	// "math"
	// "fmt"

	"golang.org/x/sys/unix"
	"github.com/mattn/go-tty"
	_ "github.com/gdamore/tcell/terminfo/extended"
	"github.com/gdamore/tcell/terminfo"
)

type Terminal struct {
	tty *tty.TTY
	ti *terminfo.Terminfo
	colorDepth ColorDepth

	cleanup func() error

	input chan []rune
	quit chan struct{}

	Lines   int // window height
	Columns   int // window weight

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

	if t.ti, e = terminfo.LookupTerminfo(os.Getenv("TERM")); e != nil {
		return e
	}

	// TODO: detect it from env
	if t.ti.Colors >= 256 {
		t.colorDepth = ColorDepth8Bits
	} else if t.ti.Colors >= 8 {
		t.colorDepth = ColorDepth4Bits
	} else {
		t.colorDepth = ColorDepth1Bit
	}

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
	// TODO: buffer stdout
	t.tty.Output().WriteString(s)
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

func (t *Terminal) Goto(row, col int) string {
	return t.ti.TGoto(col, row)
}

func (t *Terminal) TColor(c Char) string {
	if t.colorDepth == ColorDepth8Bits {
		return t.ti.TParm(t.ti.TColor(c.Foreground.Code8Bits(), c.Background.Code8Bits()))
	} else if t.colorDepth == ColorDepth4Bits {
		return t.ti.TParm(t.ti.TColor(c.Foreground.Code4Bits(), c.Background.Code4Bits()))
	}
	return ""
}
