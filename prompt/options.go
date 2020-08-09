package prompt

import (
	"github.com/alecthomas/chroma"
)


type Option func(p *Prompt)


func WithMessage(s string) Option {
	return func(p *Prompt) {
		p.message = s
	}
}


func WithLexerAndStyle(lexer chroma.Lexer, style *chroma.Style) Option {
	return func(p *Prompt) {
		p.lexer = lexer
		p.style = style
	}
}
