package layout

import (
	. "github.com/randy3k/rango/prompt/char"
)

type Prefix struct {
	messageFunc func() string
	messageContinuationFunc func(int) string
}


func NewPrefix(messageFunc func() string, messageContinuationFunc func(int) string) *Prefix {
	return &Prefix{
		messageFunc: messageFunc,
		messageContinuationFunc: messageContinuationFunc,
	}
}

func (m *Prefix) Width() int {
	return len(ANSI(m.messageFunc()))
}

func (m *Prefix) GetPrefix(lineno int) Chars {
	plen := m.Width()
	if lineno == 0 {
		return ANSI(m.messageFunc())
	} else {
		if m.messageContinuationFunc != nil {
			return ANSI(m.messageContinuationFunc(lineno))[:plen]
		} else {
			chars := make(Chars, plen)
			for i := 0; i < len(chars); i++ {
				chars[i] = NewChar(' ', DefaultAttributes)
			}
			return chars
		}
	}
}
