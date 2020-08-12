package prompt

import (
	"github.com/alecthomas/chroma"
)


type Option func(p *Prompt)


func WithMessage(s string) Option {
	return func(p *Prompt) {
		p.messageFunc = func() string {
			return s
		}
	}
}

func WithMessageFunc(f func() string) Option {
	return func(p *Prompt) {
		p.messageFunc = f
	}
}


func WithMessageContinuationFunc(f func(i int) string) Option {
	return func(p *Prompt) {
		p.messageContinuationFunc = f
	}
}


func WithLexerAndStyle(lexer chroma.Lexer, style *chroma.Style) Option {
	return func(p *Prompt) {
		p.lexer = lexer
		p.style = style
	}
}
