package prompt

import (
	. "github.com/randy3k/rango/prompt/key"
)

type Event struct {
	Keys []Key
	Data []rune
	Prompt *Prompt
}
