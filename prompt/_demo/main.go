package main

import (
	// "fmt"
	"github.com/randy3k/rango/prompt"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func main() {
	p := prompt.NewPrompt(
		prompt.WithMessage("> "),
		prompt.WithLexerAndStyle(lexers.Get("python"), styles.Get("native")),
	)
	p.Show()
}
